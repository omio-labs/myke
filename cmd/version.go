package cmd

import (
	"fmt"
	"github.com/goeuro/myke/core"
	"github.com/pkg/errors"
	"strings"
)

// Version prints myke version
func Version(opts *mykeOpts) error {
	data, err := core.Asset("tmp/version")
	if err != nil {
		return errors.Wrap(err, "error showing version")
	}

	version := strings.TrimSpace(string(data))
	fmt.Fprintf(opts.Writer, "myke version %s\n", version)
	return nil
}
