package subgraph

import (
	"context"
	"github.com/hasura/go-graphql-client"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
)

type Subgrapher interface {
	Subgraph() Subgraph
}

func NewSubgrapher(log *logan.Entry, getter kv.Getter) Subgrapher {
	return &subgrapher{
		log:    log.WithFields(logan.F{"service": "subgraph"}),
		getter: getter,
	}
}

type subgrapher struct {
	ctx    context.Context
	log    *logan.Entry
	getter kv.Getter
	comfig.Once
}

func (b *subgrapher) Subgraph() Subgraph {
	return b.Do(func() interface{} {
		var config struct {
			URL string `fig:"url,required"`
		}
		err := figure.
			Out(&config).
			From(kv.MustGetStringMap(b.getter, "subgraph")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out subgraph"))
		}
		return New(b.log, graphql.NewClient(config.URL, nil))
	}).(Subgraph)
}
