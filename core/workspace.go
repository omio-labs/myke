package core

import (
	"path/filepath"
)

// Workspace represents the current myke workspace
type Workspace struct {
	Cwd      string
	Projects []Project
}

type parseResult struct {
	p   Project
	err error
}

// ParseWorkspace parses the current workspace
func ParseWorkspace(cwd string) (Workspace, error) {
	in := make(chan parseResult)
	go func() {
		parseWorkspaceNested(cwd, "", in)
		close(in)
	}()

	w := Workspace{Cwd: cwd}
	for p := range in {
		if p.err != nil {
			return w, p.err
		}
		w.Projects = append(w.Projects, p.p)
	}

	return w, nil
}

func parseWorkspaceNested(cwd string, path string, in chan parseResult) {
	file := filepath.Join(cwd, path)
	p, err := ParseProject(file)
	if err != nil {
		in <- parseResult{p, err}
	} else {
		in <- parseResult{p, nil}
		for _, includePath := range p.Discover {
			parseWorkspaceNested(p.Cwd, includePath, in)
		}
	}
}
