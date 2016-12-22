package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/olekukonko/tablewriter"
	"strings"
)

func List(c *cli.Context) error {
	w := loadWorkspace(c.String("file"))
	table := tablewriter.NewWriter(c.App.Writer)
	table.SetBorder(false)
	table.SetHeader([]string{"project", "tags", "tasks"})

	for _, p := range w.Projects {
		tasks := []string{}
		for t := range p.Tasks {
			if !strings.HasPrefix(t, "_") {
				tasks = append(tasks, t)
			}
		}
		if len(tasks) > 0 {
			table.Append([]string{p.Name, strings.Join(p.Tags, ", "), strings.Join(tasks, ", ")})
		}
	}

	table.Render()
	return nil
}
