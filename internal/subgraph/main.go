package subgraph

type Subgraph interface {
	UserStatus(address string) (bool, error)
}
