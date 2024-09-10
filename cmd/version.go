package cmd

import (
	"go.szostok.io/version/extension"
)

const (
	repoOwner = "ondrovic"
	repoName  = "folder-structure-cli"
)

func init() {

	RootCmd.AddCommand(
		extension.NewVersionCobraCmd(
			extension.WithUpgradeNotice(repoOwner, repoName),
		),
	)
}
