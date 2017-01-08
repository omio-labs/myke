package cmd

import (
	"fmt"
	"github.com/goeuro/myke/core"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
)

// Template renders the given template file
func Template(c *cli.Context) error {
	bytes, err := ioutil.ReadFile(c.String("template"))
	if err != nil {
		return errors.Wrap(err, "error rendering template")
	}

	rendered, err := core.RenderTemplate(string(bytes), core.OsEnv(), map[string]string{})
	if err != nil {
		return errors.Wrap(err, "error rendering template")
	}

	fmt.Fprint(c.App.Writer, rendered)
	return nil
}
