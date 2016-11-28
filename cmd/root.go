package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/core"
	"os"
)

func RunOrList(c *cli.Context) error {
	if c.NArg() > 0 {
		return Run(c)
	} else {
		return List(c)
	}
}

func loadWorkspace() core.Workspace {
	cwd, _ := os.Getwd()
	w := core.ParseWorkspace(cwd)
	return w
}
