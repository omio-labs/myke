package cli

import (
	"os"
	"strings"
	"path/filepath"
)

const PathSep = string(os.PathListSeparator)

func ParseProject(path string) (Project, error) {
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
	p.Env["PATH"] = normalizePaths(p.Cwd, p.Env["PATH"])

	for _, epath := range p.Extends {
		if p, err = ExtendProject(p, epath); err != nil {
			return Project{}, err
		}
	}

	return p, nil
}

func ExtendProject(p Project, path string) (Project, error) {
	o, err := ParseProject(filepath.Join(p.Cwd, path))
	if err != nil {
		return Project{}, err
	}

	p.Tags = mergeTags(p.Tags, o.Tags)
	p.Includes = mergeTags(p.Includes, o.Includes)
	p.EnvFiles = mergeTags(p.EnvFiles, o.EnvFiles)
	p.Env = mergeEnv(p.Env, o.Env)

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

func containsTag(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func mergeEnv(first map[string]string, next map[string]string) (map[string]string) {
	for k, v := range next {
		if k == "PATH" {
			first[k] = string(first[k]) + PathSep + v
		} else if len(first[k]) == 0 {
			first[k] = v
		}
	}
	return first
}

func normalizePaths(cwd string, paths string) string {
	newPaths := []string{}
	for _, path := range strings.Split(strings.TrimSpace(paths), PathSep) {
		if len(path) > 0 {
			if !filepath.IsAbs(path) {
				path = filepath.Clean(filepath.Join(cwd, path))
			}
			newPaths = append(newPaths, path)
		}
	}

	newPaths = append(newPaths, filepath.Join(cwd, "bin"))
	return strings.Trim(strings.Join(newPaths, PathSep), PathSep)
}
