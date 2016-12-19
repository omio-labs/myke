package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/core"
	"log"
	"fmt"
)

func License(c *cli.Context) error {
	data, err := core.Asset("tmp/LICENSES")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
	return nil
}
