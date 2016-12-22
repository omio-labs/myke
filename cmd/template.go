package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/core"
	"github.com/pkg/errors"
	"fmt"
	"io/ioutil"
)

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
