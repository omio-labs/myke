package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/core"
	"log"
	"fmt"
	"io/ioutil"
)

func Template(c *cli.Context) error {
	bytes, err := ioutil.ReadFile(c.String("template"))
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
