package core

import (
	"github.com/ghodss/yaml"
	"github.com/tidwall/gjson"

	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

// Project represents a parsed project
type Project struct {
	Src      string
	Cwd      string
	Name     string
	Desc     string
	Tags     []string
	Discover []string
	Mixin    []string
	Env      map[string]string
	EnvFiles []string
	Tasks    map[string]Task
}

// ParseProject parses the given project
func ParseProject(path string) (Project, error) {
	src, err := filepath.Abs(path)
	if err != nil {
		return Project{}, err
	}

	if filepath.Ext(src) != ".yml" {
		return ParseProject(filepath.Join(src, "myke.yml"))
	}

	p, err := loadProjectYAML(src)
	if err != nil {
		return Project{}, err
	}

	p.Src = src
	p.Cwd = filepath.Dir(src)
	p.EnvFiles = append(p.EnvFiles, strings.TrimSuffix(src, ".yml")+".env")
	for _, epath := range p.EnvFiles {
		p.Env = mergeEnv(p.Env, loadEnvFile(normalizeFilePath(p.Cwd, epath)))
		p.Env = mergeEnv(p.Env, loadEnvFile(normalizeFilePath(p.Cwd, epath+".local")))
	}

	p.Env = mergeEnv(p.Env, OsEnv())
	p.Env["PATH"] = normalizeEnvPaths(p.Cwd, p.Env["PATH"])

	for _, epath := range p.Mixin {
		if p, err = mixinProject(p, epath); err != nil {
			return Project{}, err
		}
	}

	return p, nil
}

func loadProjectYAML(path string) (Project, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return Project{}, err
	}

	jsonbytes, err := yaml.YAMLToJSON(bytes)
	if err != nil {
		return Project{}, err
	}

	json := gjson.Parse(string(jsonbytes))
	return loadProjectJSON(json), nil
}

func loadProjectJSON(json gjson.Result) Project {
	p := Project{}
	if j := json.Get("project"); j.Exists() {
		p.Name = j.String()
	}
	if j := json.Get("desc"); j.Exists() {
		p.Desc = j.String()
	}
	if j := json.Get("discover"); j.Exists() {
		for _, s := range j.Array() {
			p.Discover = append(p.Discover, s.String())
		}
	}
	if j := json.Get("mixin"); j.Exists() {
		for _, s := range j.Array() {
			p.Mixin = append(p.Mixin, s.String())
		}
	}
	p.Env = make(map[string]string)
	if j := json.Get("env"); j.Exists() {
		for k, v := range j.Map() {
			p.Env[k] = v.String()
		}
	}
	if j := json.Get("env_files"); j.Exists() {
		for _, s := range j.Array() {
			p.EnvFiles = append(p.EnvFiles, s.String())
		}
	}
	if j := json.Get("tags"); j.Exists() {
		for _, s := range j.Array() {
			p.Tags = append(p.Tags, s.String())
		}
		sort.Strings(p.Tags)
	}
	p.Tasks = make(map[string]Task)
	if j := json.Get("tasks"); j.Exists() {
		for k, v := range j.Map() {
			p.Tasks[k] = loadTaskJSON(k, v)
		}
	}
	return p
}

func mixinProject(child Project, path string) (Project, error) {
	parent, err := ParseProject(normalizeFilePath(child.Cwd, path))
	if err != nil {
		return Project{}, err
	}

	child.Tags = mergeTags(parent.Tags, child.Tags)
	child.Discover = mergeTags(parent.Discover, child.Discover)
	child.EnvFiles = mergeTags(parent.EnvFiles, child.EnvFiles)
	child.Env = mergeEnv(parent.Env, child.Env)
	sort.Strings(child.Tags)

	for taskName, parentTask := range parent.Tasks {
		child.Tasks[taskName] = mixinTask(taskName, child.Tasks[taskName], parentTask)
	}

	return child, nil
}
