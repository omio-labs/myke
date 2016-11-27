package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"log"
	"fmt"
	"io/ioutil"
	"myke/core"
)

var TemplateCmd = cli.Command{
	Name: "template",
	Usage: "render a template file with environment variables",
	Action: Template,
}

func Template(c *cli.Context) error {
	bytes, err := ioutil.ReadFile(c.Args().First())
	if err != nil {
		log.Fatal(err)
	}

	rendered, err := core.RenderTemplate(string(bytes), core.OsEnv(), map[string]string{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(rendered)
	return nil
}
