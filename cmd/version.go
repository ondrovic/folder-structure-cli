package cmd

import (
	"go.szostok.io/version/extension"
)

func init() {
	RootCmd.AddCommand(
		extension.NewVersionCobraCmd(
			extension.WithUpgradeNotice("ondrovic", "folder-structure-cli"),
		),
	)
}
