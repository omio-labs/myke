package cmd

import (
	"fmt"
	"github.com/goeuro/myke/core"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v1"
)

// License prints all open source licenses
func License(c *cli.Context) error {
	data, err := core.Asset("tmp/LICENSES")
	if err != nil {
		return errors.Wrap(err, "error showing licenses")
	}

	fmt.Fprintln(c.App.Writer, string(data))
	return nil
}
