package cli

import (
	"path/filepath"
)

func LoadWorkspace(cwd string, path string) (Workspace) {
	in := make(chan Project)
	go loadWorkspaceChannel(cwd, path, in)

	w := Workspace{Cwd: cwd}
	for p := range in {
		w.Projects = append(w.Projects, p)
	}

	return w
}



func loadWorkspaceChannel(cwd string, path string, in chan Project) {
	loadWorkspaceNested(cwd, path, in)
	close(in)
}

func loadWorkspaceNested(cwd string, path string, in chan Project) {
	p, _ := ParseProject(filepath.Join(cwd, path))
	in <- p
	for _, includePath := range p.Includes {
		loadWorkspaceNested(p.Cwd, includePath, in)
	}
}
