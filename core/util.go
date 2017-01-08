package core

import (
	"github.com/joho/godotenv"

	"os"
	"path/filepath"
	"strings"
)

const pathSep = string(os.PathListSeparator)

func mergeTags(first []string, next []string) []string {
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

func loadEnvFile(path string) map[string]string {
	env, err := godotenv.Read(path)
	if err != nil {
		return make(map[string]string)
	}
	return env
}

// OsEnv returns current process environment as a map
func OsEnv() map[string]string {
	env := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if pair[0] != "PATH" {
			// PATH is handled as a special case, so lets skip it
			env[pair[0]] = pair[1]
		}
	}
	return env
}

func mergeEnv(first map[string]string, next map[string]string) map[string]string {
	for k, v := range next {
		if k == "PATH" {
			first[k] = v + pathSep + string(first[k])
		} else {
			first[k] = v
		}
	}
	return first
}

func union(first map[string]string, next map[string]string) map[string]string {
	res := make(map[string]string)
	for k, v := range first {
		res[k] = v
	}
	for k, v := range next {
		res[k] = v
	}
	return res
}

func normalizeEnvPaths(cwd string, paths string) string {
	newPaths := []string{}
	for _, path := range strings.Split(strings.TrimSpace(paths), pathSep) {
		if len(path) > 0 {
			if !filepath.IsAbs(path) {
				path = filepath.Clean(filepath.Join(cwd, path))
			}
			newPaths = append(newPaths, path)
		}
	}

	newPaths = append(newPaths, filepath.Join(cwd, "bin"))
	return strings.Trim(strings.Join(newPaths, pathSep), pathSep)
}

// Feature TODO: envvar interpolation
func normalizeFilePath(cwd string, path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(cwd, path)
}

func retry(f func(attempt int) (retry bool, err error)) error {
	attempt := 1
	for {
		cont, err := f(attempt)
		if err == nil || !cont {
			return err
		}
		attempt++
	}
}
