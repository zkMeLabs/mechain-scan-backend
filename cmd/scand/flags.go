package main

import (
	"github.com/urfave/cli/v2"
)

var (
	chainIDFlag = &cli.StringFlag{
		Name:     "chain-id",
		Usage:    "The chain ID of the mechain node",
		Required: true,
		EnvVars:  []string{"CHAIN_ID"},
	}
	endpointFlag = &cli.StringFlag{
		Name:     "endpoint",
		Usage:    "The RPC URL of the mechain node",
		Required: true,
		EnvVars:  []string{"ENDPOINT"},
	}
	privateKeyFlag = &cli.StringFlag{
		Name:     "private-key",
		Usage:    "The private key of the account",
		Required: true,
		EnvVars:  []string{"PRIVATE_KEY"},
	}
)
