package handlers

import (
	"context"
	"gitlab.com/rarimo/questn-integration-svc/internal/subgraph"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/rarimo/questn-integration-svc/internal/config"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	configCtxKey
	subgraphCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

// CtxConfig adds config provider instance to ctx.
func CtxConfig(cfg config.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, configCtxKey, cfg)
	}
}

// Config returns the config provider instance stored in ctx.
func Config(r *http.Request) config.Config {
	return r.Context().Value(configCtxKey).(config.Config)
}

// CtxSubgraph adds subgraph provider instance to ctx.
func CtxSubgraph(sub subgraph.Subgraph) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, subgraphCtxKey, sub)
	}
}

// Subgraph returns the subgraph provider instance stored in ctx.
func Subgraph(r *http.Request) subgraph.Subgraph {
	return r.Context().Value(subgraphCtxKey).(subgraph.Subgraph)
}
