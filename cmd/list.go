package cmd

import (
	"github.com/rdsubhas/go-elastictable"
	"gopkg.in/urfave/cli.v1"
	"sort"
	"strings"
)

var headers = []string{"PROJECT", "TAGS", "TASKS"}

// List lists tasks
func List(c *cli.Context) error {
	w := loadWorkspace(c.String("file"))
	t := elastictable.NewElasticTable(headers)
	for _, p := range w.Projects {
		tasks := []string{}
		for t := range p.Tasks {
			if !strings.HasPrefix(t, "_") {
				tasks = append(tasks, t)
			}
		}
		if len(tasks) > 0 {
			sort.Strings(tasks)
			t.AddRow([]string{p.Name, strings.Join(p.Tags, ", "), strings.Join(tasks, ", ")})
		}
	}

	t.Render(c.App.Writer)
	return nil
}
