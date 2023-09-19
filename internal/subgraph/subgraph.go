package subgraph

import (
	"context"
	"github.com/hasura/go-graphql-client"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/rarimo/questn-integration-svc/internal/subgraph/queries"
)

type subgraph struct {
	graph *graphql.Client
}

func New(graph *graphql.Client) Subgraph {
	return &subgraph{graph}
}

func (s *subgraph) UserStatus(address string) (bool, error) {
	var query queries.UsersByAddressQuery

	variables := map[string]interface{}{
		"address": address,
	}

	err := s.graph.Query(context.TODO(), &query, variables)
	if err != nil {
		return false, errors.Wrap(err, "failed to query subgraph for user status")
	}

	return len(query.Users) > 0, nil
}
