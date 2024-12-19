package main

import (
	"os"
	"runtime/debug"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/crypto-org-chain/chain-main/v4/app"
	cmd "github.com/crypto-org-chain/chain-main/v4/cmd/chain-maind/app"
)

func main() {
	// Set the maximum stack size
	debug.SetMaxStack(20 * 1024 * 1024) // 20MB

	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, cmd.EnvPrefix, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}
