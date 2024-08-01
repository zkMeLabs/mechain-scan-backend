package api

// NewClient returns a new greenfield client
import (
	"github.com/bnb-chain/greenfield-go-sdk/client"
)

func NewClient(chainID string, endpoint string) client.IClient {
	cli, err := client.NewClient(chainID, endpoint)
	if err != nil {
		panic(err)
	}
	return cli
}
