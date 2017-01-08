package cmd

import (
	"github.com/goeuro/myke/core"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v1"
)

// Run runs the given tasks
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
