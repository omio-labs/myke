package main

import (
	"github.com/goeuro/myke/cmd"
	"github.com/apex/log"
	"os"
)

func main() {
	if err := cmd.NewApp().Run(os.Args); err != nil {
		log.WithError(err).Fatal("error")
	}
}
