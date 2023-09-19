package subgraph

import "context"

type Subgraph interface {
	UserStatus(ctx context.Context, address string) (bool, error)
}
