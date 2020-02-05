package main

import (
	"github.com/tendermint/tendermint/libs/cli"

	app "github.com/shikhardb/tienchain"
	"github.com/shikhardb/tienchain/starter"
)

func main() {

	params := starter.NewServerCommandParams(
		"tcd",
		"tienchain AppDaemon",
		starter.NewAppCreator(app.NewTienChainApp),
		starter.NewAppExporter(app.NewTienChainApp),
	)

	serverCmd := starter.NewServerCommand(params)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(serverCmd, "TC", starter.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
