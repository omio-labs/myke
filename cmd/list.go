package cmd

import (
	"github.com/rdsubhas/go-elastictable"
	"sort"
	"strings"
)

var headers = []string{"PROJECT", "TAGS", "TASKS"}

// List tasks
func List(opts *mykeOpts) error {
	w := loadWorkspace(opts.File)
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

	t.Render(opts.Writer)
	return nil
}
