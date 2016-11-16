package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"log"
	"myke/core"
)

func RunOrList(c *cli.Context) error {
	if c.NArg() > 0 {
		return Run(c)
	} else {
		return List(c)
	}
}

func Run(c *cli.Context) error {
	qs := make([]string, len(c.Args()))
	for i, v := range c.Args() {
		qs[i] = v
	}

	queries, err := core.ParseQueries(qs)
	if err != nil {
		log.Fatal(err)
	}

	w := loadWorkspace()
	for _, q := range queries {
		err := core.ExecuteQuery(&w, q)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
