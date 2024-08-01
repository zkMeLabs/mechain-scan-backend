package main

import (
	"log"
	"os"

	"github.com/bnb-chain/greenfield-go-sdk/client"
	"github.com/bnb-chain/greenfield-go-sdk/types"
	"github.com/gin-gonic/gin"
	cli "github.com/urfave/cli/v2"
	"github.com/zkMeLabs/mechain-scan/api"
)

func main() {
	app := &cli.App{
		Name:  "scanner",
		Usage: "Query mechain information via API",
		Flags: []cli.Flag{
			chainIDFlag,
			endpointFlag,
			privateKeyFlag,
		},
		Action: func(c *cli.Context) error {
			account, err := types.NewAccountFromPrivateKey("gnfd-account", c.String(privateKeyFlag.Name))
			if err != nil {
				return err
			}
			cli, err := client.New(c.String(chainIDFlag.Name), c.String(endpointFlag.Name), client.Option{DefaultAccount: account})
			if err != nil {
				return err
			}
			a := api.NewAPI(cli)
			r := gin.Default()
			r.GET("/buckets", a.GetBuckets)
			r.GET("/objects", a.GetObjects)
			r.GET("/groups", a.GetGroups)

			port := os.Getenv("PORT")
			if port == "" {
				port = "8080"
			}
			return r.Run(":" + port)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
