package main

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/cmd"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "myke"
	app.Version = "0.3.2"
	app.Usage = "make with yml"
	app.Action = cmd.RunOrList
	app.Flags = []cli.Flag {
		cli.StringFlag{
			 Name: "f, file",
			 Value: "myke.yml",
			 Usage: "`yml-file` to load",
		},
	}
	app.Commands = []cli.Command{
		cmd.ListCmd,
		cmd.RunCmd,
		cmd.TemplateCmd,
		cmd.LicenseCmd,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
