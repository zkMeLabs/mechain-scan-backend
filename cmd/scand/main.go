package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	cli "github.com/urfave/cli/v2"
	"github.com/zkMeLabs/mechain-scan/api"
)

func main() {
	app := &cli.App{
		Name:  "scanner",
		Usage: "Query mechain information via API",
		Flags: flags,
		Action: func(c *cli.Context) error {
			r := gin.Default()
			r.GET("/objects", api.GetObjects)
			r.GET("/buckets", api.GetBuckets)
			r.GET("/groups", api.GetGroups)

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
