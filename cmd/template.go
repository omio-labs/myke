package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"log"
	"fmt"
	"io/ioutil"
	"myke/core"
)

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
