package cli

import (
	"github.com/olekukonko/tablewriter"

	"os"
	"strings"
	"path/filepath"
)

type Workspace struct {
	Cwd      string
	Projects []Project
}

func (w *Workspace) List() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"project", "tags", "tasks"})
	for _, p := range w.Projects {
		tasks := []string{}
		for t, _ := range p.Tasks {
			tasks = append(tasks, t)
		}
		table.Append([]string{p.Name, strings.Join(p.Tags, ","), strings.Join(tasks, ",")})
	}
	table.Render()
}

func ParseWorkspace(cwd string) (Workspace) {
	in := make(chan Project)
	go func() {
		parseWorkspaceNested(cwd, "", in)
		close(in)
	}()

	w := Workspace{Cwd: cwd}
	for p := range in {
		w.Projects = append(w.Projects, p)
	}

	return w
}

func parseWorkspaceNested(cwd string, path string, in chan Project) {
	p, _ := ParseProject(filepath.Join(cwd, path))
	in <- p
	for _, includePath := range p.Includes {
		parseWorkspaceNested(p.Cwd, includePath, in)
	}
}
