package main

import (
	"github.com/apex/log"
	"github.com/omio-labs/myke/cmd"
	"os"
)

func main() {
	if err := cmd.Exec(os.Args[1:]); err != nil {
		log.WithError(err).Error("error")
		os.Exit(1)
	}
}
