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

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "myke"
	app.Version = Version()
	app.Usage = "make with yml"
	app.Action = Action
	app.Flags = []cli.Flag {
		cli.StringFlag{
			 Name: "f, file",
			 Value: "myke.yml",
			 Usage: "`yml-file` to load",
		},
		cli.StringFlag{
			 Name: "template",
			 Usage: "render template `tpl-file` (will not run any command)",
		},
		cli.BoolFlag{
			 Name: "license",
			 Usage: "show license",
		},
		cli.StringFlag{
			Name: "loglevel",
			Value: "info",
			Usage: "log level, one of debug|`info`|warn|error|fatal",
		},
	}
	return app
}

func Action(c *cli.Context) error {
	log.SetHandler(&logcli.Handler{Writer: os.Stdout, Padding: 0})
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
