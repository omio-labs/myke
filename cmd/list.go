package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"strings"
)

func List(c *cli.Context) error {
	w := loadWorkspace(c.String("file"))
	t := new(ElasticTable)
	t.Init([]string{"PROJECT", "TAGS", "TASKS"})
	for _, p := range w.Projects {
		tasks := []string{}
		for t := range p.Tasks {
			if !strings.HasPrefix(t, "_") {
				tasks = append(tasks, t)
			}
		}
		if len(tasks) > 0 {
			t.AddRow([]string{p.Name, strings.Join(p.Tags, ", "), strings.Join(tasks, ", ")})
		}
	}

	t.Render(c.App.Writer)
	return nil
}
