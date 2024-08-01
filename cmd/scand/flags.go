package main

import (
	"github.com/urfave/cli/v2"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:     "rpc-url",
		Usage:    "The RPC URL of the mechain node",
		Required: true,
		EnvVars:  []string{"RPC_URL"},
	},
	&cli.StringFlag{
		Name:     "chain-id",
		Usage:    "The chain ID of the mechain node",
		Required: true,
		EnvVars:  []string{"CHAIN_ID"},
	},
}
