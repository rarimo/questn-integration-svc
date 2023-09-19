package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/rarimo/questn-integration-svc/internal/subgraph"
)

type Config interface {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	subgraph.Subgrapher
}

type config struct {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	subgraph.Subgrapher
	getter kv.Getter
}

func New(getter kv.Getter) Config {
	logger := comfig.NewLogger(getter, comfig.LoggerOpts{})
	return &config{
		getter:     getter,
		Logger:     logger,
		Copuser:    copus.NewCopuser(getter),
		Listenerer: comfig.NewListenerer(getter),
		Subgrapher: subgraph.NewSubgrapher(getter),
	}
}
