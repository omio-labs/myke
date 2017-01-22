package cmd

import (
	"fmt"
	"github.com/goeuro/myke/core"
	"github.com/pkg/errors"
)

// License prints all open source licenses
func License(opts *mykeOpts) error {
	data, err := core.Asset("tmp/LICENSES")
	if err != nil {
		return errors.Wrap(err, "error showing licenses")
	}

	fmt.Fprintln(opts.Writer, string(data))
	return nil
}
