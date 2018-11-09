package cmd

import (
	"fmt"

	"github.com/goeuro/myke/core"
)

// License prints all open source licenses
func License(opts *mykeOpts) error {
	data, _ := core.FS.String("/tmp/LICENSES")
	fmt.Fprintln(opts.Writer, string(data))
	return nil
}
