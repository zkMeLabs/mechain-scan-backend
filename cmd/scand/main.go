package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "mechain-scan",
		Usage: "mechain scanner",
		Flags: []cli.Flag{
			configFileFlag,
			rpcFlag,
			dbHostFlag,
			dbPortFlag,
			dbUserFlag,
			dbPassFlag,
			apiHostFlag,
			apiPortFlag,
		},
		Commands: []*cli.Command{},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
