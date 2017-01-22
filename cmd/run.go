package cmd

import (
	"github.com/goeuro/myke/core"
	"github.com/pkg/errors"
)

// Run runs the given tasks
func Run(opts *mykeOpts, tasks []string) error {
	queries, err := core.ParseQueries(tasks)
	if err != nil {
		return errors.Wrap(err, "error parsing command")
	}

	w := loadWorkspace(opts.File)
	for _, q := range queries {
		err := core.ExecuteQuery(&w, q)
		if err != nil {
			return errors.Wrap(err, "error running command")
		}
	}

	return nil
}
