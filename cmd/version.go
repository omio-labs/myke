package cmd

import (
	"fmt"
	"strings"

	"github.com/goeuro/myke/core"
)

// Version prints myke version
func Version(opts *mykeOpts) error {
	data, _ := core.FS.String("/tmp/version")
	version := strings.TrimSpace(string(data))
	fmt.Fprintf(opts.Writer, "myke version %s\n", version)
	return nil
}
