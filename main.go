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
	app.Version = "0.3.0"
	app.Usage = "make with yml"
	app.Action = cmd.RunOrList
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
