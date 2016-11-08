package main

import (
	"myke/cli"

	"os"
)

func main() {
	cwd, _ := os.Getwd()
	w := cli.ParseWorkspace(cwd)
	w.List()
}
