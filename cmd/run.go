package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/core"
	"github.com/apex/log"
)

func Run(c *cli.Context) error {
	qs := make([]string, len(c.Args()))
	for i, v := range c.Args() {
		qs[i] = v
	}

	queries, err := core.ParseQueries(qs)
	if err != nil {
		log.WithError(err).Fatal("error parsing command")
	}

	w := loadWorkspace(c.String("file"))
	for _, q := range queries {
		err := core.ExecuteQuery(&w, q)
		if err != nil {
			log.WithError(err).Fatal("error executing command")
		}
	}

	return nil
}
