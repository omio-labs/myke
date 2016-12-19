package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/core"
	"github.com/apex/log"
	logcli "github.com/apex/log/handlers/cli"
	"path/filepath"
	"strings"
	"os"
)

func Action(c *cli.Context) error {
	log.SetHandler(&logcli.Handler{Writer: os.Stderr, Padding: 0})
	if level, err := log.ParseLevel(c.String("loglevel")); err == nil {
		log.SetLevel(level)
	}

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
