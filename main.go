package main

import (
	"myke/core"

	"os"
)

func main() {
	cwd, _ := os.Getwd()
	w := core.ParseWorkspace(cwd)
	w.List()
}
