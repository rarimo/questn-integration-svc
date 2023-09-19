package subgraph

import (
	"context"
	"github.com/hasura/go-graphql-client"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/rarimo/questn-integration-svc/internal/subgraph/queries"
)

type subgraph struct {
	log   *logan.Entry
	graph *graphql.Client
}

func New(log *logan.Entry, graph *graphql.Client) Subgraph {
	return &subgraph{log, graph}
}

func (s *subgraph) UserStatus(ctx context.Context, address string) (bool, error) {
	log := s.log.WithFields(logan.F{"address": address})
	var query queries.UsersByAddressQuery

	variables := map[string]interface{}{
		"address": address,
	}

	log.Debug("querying subgraph for user status")

	err := s.graph.Query(ctx, &query, variables)
	if err != nil {
		return false, errors.Wrap(err, "failed to query subgraph for user status")
	}

	result := len(query.Users) > 0

	log.WithFields(logan.F{"result": result}).Debug("subgraph query successful for user status")

	return result, nil
}
