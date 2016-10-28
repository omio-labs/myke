package cli

import (
	"os"
	"path/filepath"
)

func ParseFile(path string) (Project, error) {
	src, err := filepath.Abs(path)
	if err != nil {
		return Project{}, err
	}

	if info, err := os.Stat(src); err != nil {
		return Project{}, err
	} else if info.IsDir() {
		src = filepath.Join(src, "myke.yml")
	}

	p, err := LoadFile(src)
	if err != nil {
		return Project{}, err
	}

	p.Src = src
	p.Cwd = filepath.Dir(src)

	for _, epath := range p.Extends {
		if p, err = ExtendProject(p, epath); err != nil {
			return Project{}, err
		}
	}

	return p, nil
}

func ExtendProject(p Project, path string) (Project, error) {
	o, err := LoadFile(filepath.Join(p.Cwd, path))
	if err != nil {
		return Project{}, err
	}

	p.Tags = mergeTags(p.Tags, o.Tags)
	p.Includes = mergeTags(p.Includes, o.Includes)
	p.EnvFiles = mergeTags(p.EnvFiles, o.EnvFiles)

	return p, nil
}

func mergeTags(first []string, next []string) ([]string) {
	for _, v := range next {
		if !containsTag(first, v) {
			first = append(first, v)
		}
	}
	return first
}

func mergeEnv(first map[string]string, next map[string]string) (map[string]string) {
	for k, v := range next {
		if k == "PATH" {
			first[k] = first[k] + string(os.PathListSeparator) + v
		} else if len(first[k]) == 0 {
			first[k] = v
		}
	}
	return first
}

func containsTag(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
