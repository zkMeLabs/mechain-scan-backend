package main

import (
	"github.com/urfave/cli/v2"
)

var (
	configFlag = "config"
	rpcAddr    = "rpc-addr"
	dbHost     = "db-host"
	dbPort     = "db-port"
	dbUser     = "db-user"
	dbPass     = "db-pass"
	apiHost    = "api-host"
	apiPort    = "api-addr"
)

var (
	configFileFlag = &cli.StringFlag{
		Name:  configFlag,
		Usage: "Load configuration from `FILE`",
	}
	rpcFlag = &cli.StringFlag{
		Name:  rpcAddr,
		Usage: "rpc address",
	}
	dbUserFlag = &cli.StringFlag{
		Name:  dbUser,
		Usage: "database username",
	}
	dbPassFlag = &cli.StringFlag{
		Name:  dbPass,
		Usage: "database password",
	}
	dbHostFlag = &cli.StringFlag{
		Name:  dbHost,
		Usage: "database host",
	}
	dbPortFlag = &cli.StringFlag{
		Name:  dbPort,
		Usage: "database port",
	}
	apiHostFlag = &cli.StringFlag{
		Name:  apiHost,
		Usage: "api host",
	}
	apiPortFlag = &cli.StringFlag{
		Name:  apiPort,
		Usage: "api address",
	}
)
