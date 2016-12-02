package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/core"
	"path/filepath"
	"os"
)

func RunOrList(c *cli.Context) error {
	if c.NArg() > 0 {
		return Run(c)
	} else {
		return List(c)
	}
}

func loadWorkspace(path string) core.Workspace {
	if !filepath.IsAbs(path) {
		cwd, _ := os.Getwd()
		path = filepath.Join(cwd, path)
	}
	return core.ParseWorkspace(path)
}
