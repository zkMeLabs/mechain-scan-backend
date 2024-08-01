package api

import "github.com/bnb-chain/greenfield-go-sdk/client"

const (
	iso8601DateFormat = "2006-01-02 15:04:05"
	maxListMemberNum  = 50
	defaultMaxKey     = 500
)

type API struct {
	cli client.IClient
}

func NewAPI(c client.IClient) *API {
	return &API{cli: c}
}
