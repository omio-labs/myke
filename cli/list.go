package cli

import (
	"github.com/olekukonko/tablewriter"

	"myke/core"
	"os"
	"strings"
)

func List() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetHeader([]string{"project", "tags", "tasks"})

	w := loadWorkspace()
	for _, p := range w.Projects {
		tasks := []string{}
		for t, _ := range p.Tasks {
			tasks = append(tasks, t)
		}
		table.Append([]string{p.Name, strings.Join(p.Tags, ", "), strings.Join(tasks, ", ")})
	}

	table.Render()
}

func loadWorkspace() core.Workspace {
	cwd, _ := os.Getwd()
	w := core.ParseWorkspace(cwd)
	return w
}
