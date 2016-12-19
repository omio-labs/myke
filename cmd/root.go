package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/core"
	"path/filepath"
	"strings"
	"os"
)

func Action(c *cli.Context) error {
	if len(c.String("template")) > 0 {
		return Template(c)
	} else if c.Bool("license") {
		return License(c)
	} else if c.NArg() > 0 {
		return Run(c)
	} else {
		return List(c)
	}
}

func Version() string {
	version, _ := core.Asset("tmp/version")
	return strings.TrimSpace(string(version))
}

func loadWorkspace(path string) core.Workspace {
	if !filepath.IsAbs(path) {
		cwd, _ := os.Getwd()
		path = filepath.Join(cwd, path)
	}
	return core.ParseWorkspace(path)
}
