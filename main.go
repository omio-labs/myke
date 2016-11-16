package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"myke/cmd"
)

var (
	runCommand = kingpin.Command("run", "Run tasks").Default()
	runQueries = runCommand.Arg("query", "Query to execute of the format tag1/tag2/project/task[arg1, arg2, ...]").Strings()

	listCommand = kingpin.Command("list", "List available tasks")

	templateCommand = kingpin.Command("template", "Render a template file with environment variables")
	templateFile = templateCommand.Arg("file", "Template file").Required().String()
)

func main() {
	switch kingpin.Parse() {
	case "list":
		cmd.List()
	case "run":
		if len(*runQueries) == 0 {
			cmd.List()
		} else {
			cmd.Run(*runQueries)
		}
	case "template":
		cmd.Template(*templateFile)
	}
}
