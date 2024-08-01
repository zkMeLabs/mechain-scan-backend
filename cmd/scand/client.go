package main

// NewClient returns a new greenfield client
import (
	"github.com/bnb-chain/greenfield-go-sdk/client"
)

func NewClient(chainID string, endpoint string) client.IClient {
	cli, err := client.New(chainID, endpoint, client.Option{})
	if err != nil {
		panic(err)
	}
	return cli
}
