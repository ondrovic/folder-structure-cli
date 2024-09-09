package main

import (
	"fmt"
	"github.com/ondrovic/folder-structure-cli/cmd"
	"runtime"

	sCli "github.com/ondrovic/common/utils/cli"
)

func main() {
	if err := sCli.ClearTerminalScreen(runtime.GOOS); err != nil {
		fmt.Println(err)
		return
	}
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
