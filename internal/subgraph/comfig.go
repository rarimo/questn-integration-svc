package subgraph

import (
	"github.com/hasura/go-graphql-client"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Subgrapher interface {
	Subgraph() Subgraph
}

func NewSubgrapher(getter kv.Getter) Subgrapher {
	return &subgrapher{
		getter: getter,
	}
}

type subgrapher struct {
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
		return New(graphql.NewClient(config.URL, nil))
	}).(Subgraph)
}
