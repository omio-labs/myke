package main

import (
	"gopkg.in/urfave/cli.v1"
	"myke/cmd"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "myke"
	app.Version = "0.2.0"
	app.Usage = "make with yml"
	app.Action = cmd.RunOrList
	app.Commands = []cli.Command{
		{
			Name: "list",
			Usage: "list available tasks",
			Action: cmd.List,
		},
		{
			Name: "run",
			Usage: "query to execute of format tag1/tag2/project/task[arg1=val1,arg2=val2,...]",
			Action: cmd.Run,
		},
		{
			Name: "template",
			Usage: "render a template file with environment variables",
			Action: cmd.Template,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
