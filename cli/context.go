package cli

import (
	"path/filepath"
)

func Discover(cwd string, path string, in chan Project) {
	discover(cwd, path, in)
	close(in)
}

func discover(cwd string, path string, in chan Project) {
	p, _ := ParseProject(filepath.Join(cwd, path))
	in <- p
	for _, includePath := range p.Includes {
		discover(p.Cwd, includePath, in)
	}
}
