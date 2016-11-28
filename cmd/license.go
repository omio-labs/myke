package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/core"
	"log"
	"fmt"
)

var LicenseCmd = cli.Command{
	Name: "license",
	Usage: "prints licenses",
	Action: License,
}

func License(c *cli.Context) error {
	data, err := core.Asset("tmp/LICENSES")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
	return nil
}
