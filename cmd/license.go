package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"log"
	"fmt"
	"myke/core"
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
