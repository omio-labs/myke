package main

import (
	"github.com/apex/log"
	"github.com/goeuro/myke/cmd"
	"os"
)

func main() {
	if err := cmd.NewApp().Run(os.Args); err != nil {
		log.WithError(err).Fatal("error")
	}
}
