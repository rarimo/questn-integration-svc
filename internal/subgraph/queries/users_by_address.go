package queries

type UsersByAddressQuery struct {
	Users []struct {
		ID string
	} `graphql:"users(where:{senderAddr: $address})"`
}
