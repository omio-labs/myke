package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/core"
	"github.com/pkg/errors"
)

func Run(c *cli.Context) error {
	qs := make([]string, len(c.Args()))
	for i, v := range c.Args() {
		qs[i] = v
	}

	queries, err := core.ParseQueries(qs)
	if err != nil {
		return errors.Wrap(err, "error parsing command")
	}

	w := loadWorkspace(c.String("file"))
	for _, q := range queries {
		err := core.ExecuteQuery(&w, q)
		if err != nil {
			return errors.Wrap(err, "error running command")
		}
	}

	return nil
}
