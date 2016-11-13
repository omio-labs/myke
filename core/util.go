package core

import (
	"github.com/joho/godotenv"

	"os"
	"strings"
	"path/filepath"
)

const PathSep = string(os.PathListSeparator)

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

func loadEnvFile(path string) (map[string]string) {
	if env, err := godotenv.Read(path); err != nil {
		return make(map[string]string)
	} else {
		return env
	}
}

func loadOsEnv() (map[string]string) {
	env := make(map[string]string)
  for _, e := range os.Environ() {
    pair := strings.SplitN(e, "=", 2)
  	env[pair[0]] = pair[1]
  }
  return env
}

func mergeEnv(first map[string]string, next map[string]string) (map[string]string) {
	for k, v := range next {
		if k == "PATH" {
			first[k] = v + PathSep + string(first[k])
		} else {
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
