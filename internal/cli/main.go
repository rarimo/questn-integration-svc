package cli

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rarimo/questn-integration-svc/internal/services/api"

	"github.com/rarimo/questn-integration-svc/internal/config"

	"github.com/alecthomas/kingpin"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func Run(args []string) {
	defer func() {
		if rvr := recover(); rvr != nil {
			logan.New().WithRecover(rvr).Fatal("app panicked")

		}
		os.Stdout.Sync()
		os.Stderr.Sync()
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := config.New(kv.MustFromEnv())
	log := cfg.Log()

	var wg sync.WaitGroup
	run := func(f func(config.Config, context.Context)) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() {
				os.Stdout.Sync()
				os.Stderr.Sync()

				if rvr := recover(); rvr != nil {
					err := errors.FromPanic(rvr)
					logan.New().WithError(err).Fatal("one of the services panicked")
				}
			}()
			f(cfg, ctx)
		}()
	}

	app := kingpin.New("questn-integration-svc", "")

	runCmd := app.Command("run", "run command")
	apiCmd := runCmd.Command("api", "run api")

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Fatal("failed to parse arguments")
	}

	switch cmd {
	case apiCmd.FullCommand():
		log.Info("starting API")
		run(api.Run)
	default:
		log.Fatal("unknown command %s", cmd)
	}

	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	// making WaitGroup usable in select
	wgch := make(chan struct{})
	go func() {
		wg.Wait()
		close(wgch)
	}()

	select {
	// listening for runners stop
	case <-wgch:
		cfg.Log().Warn("all services stopped")
	// listening for OS signals
	case <-gracefulStop:
		cfg.Log().Info("received signal to stop")
		cancel()
		<-wgch
	}
}
