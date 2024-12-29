package main

import (
	"e-wallet-wallet/cmd"
	"e-wallet-wallet/helpers"
)

func main() {
	// load config
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	// load db
	helpers.SetupMySql()

	// run grpc
	//go cmd.ServeGRPC()

	// run http
	cmd.ServeHTTP()
}
