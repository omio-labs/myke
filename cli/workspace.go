package cli

import (
	"path/filepath"
)

type Workspace struct {
	Cwd      string
	Projects []Project
}

func ParseWorkspace(cwd string, path string) (Workspace) {
	in := make(chan Project)
	go parseWorkspace(cwd, path, in)

	w := Workspace{Cwd: cwd}
	for p := range in {
		w.Projects = append(w.Projects, p)
	}

	return w
}

func parseWorkspace(cwd string, path string, in chan Project) {
	parseWorkspaceNested(cwd, path, in)
	close(in)
}

func parseWorkspaceNested(cwd string, path string, in chan Project) {
	p, _ := ParseProject(filepath.Join(cwd, path))
	in <- p
	for _, includePath := range p.Includes {
		parseWorkspaceNested(p.Cwd, includePath, in)
	}
}
