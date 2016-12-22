package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/core"
	"github.com/pkg/errors"
	"fmt"
)

func License(c *cli.Context) error {
	data, err := core.Asset("tmp/LICENSES")
	if err != nil {
		return errors.Wrap(err, "error showing licenses")
	}

	fmt.Fprintln(c.App.Writer, string(data))
	return nil
}
