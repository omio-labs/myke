package cmd

import (
	"myke/core"
	"os"
)

func loadWorkspace() core.Workspace {
	cwd, _ := os.Getwd()
	w := core.ParseWorkspace(cwd)
	return w
}
