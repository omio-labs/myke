package core

import (
	"path/filepath"
)

type Workspace struct {
	Cwd      string
	Projects []Project
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
