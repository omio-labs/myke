package core

import (
	"github.com/tidwall/gjson"
	"github.com/ghodss/yaml"

	"os"
	"path/filepath"
	"io/ioutil"
)

type Project struct {
	Src      string
	Cwd      string
	Name     string
	Desc     string
	Includes []string
	Extends  []string
	Env      map[string]string
	EnvFiles []string
	Tags     []string
	Tasks    map[string]Task
}

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

	p, err := loadProjectYaml(src)
	if err != nil {
		return Project{}, err
	}

	p.Src = src
	p.Cwd = filepath.Dir(src)
	p.EnvFiles = append(p.EnvFiles, filepath.Join(p.Cwd, "myke.env"), filepath.Join(p.Cwd, "myke.env.local"))
	for _, epath := range p.EnvFiles {
		p.Env = mergeEnv(p.Env, loadEnvFile(epath))
	}
	// TODO: Merge OsEnv()
	p.Env["PATH"] = normalizePaths(p.Cwd, p.Env["PATH"])

	for _, epath := range p.Extends {
		if p, err = extendProject(p, epath); err != nil {
			return Project{}, err
		}
	}

	return p, nil
}

func loadProjectYaml(path string) (Project, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return Project{}, err
	}

	jsonbytes, err := yaml.YAMLToJSON(bytes)
	if err != nil {
		return Project{}, err
	}

	json := gjson.Parse(string(jsonbytes))
	return loadProjectJson(json), nil
}

func loadProjectJson(json gjson.Result) Project {
	p := Project{}
	if j := json.Get("project"); j.Exists() {
		p.Name = j.String()
	}
	if j := json.Get("desc"); j.Exists() {
		p.Desc = j.String()
	}
	if j := json.Get("includes"); j.Exists() {
		for _, s := range j.Array() {
			p.Includes = append(p.Includes, s.String())
		}
	}
	if j := json.Get("extends"); j.Exists() {
		for _, s := range j.Array() {
			p.Extends = append(p.Extends, s.String())
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
	}
	p.Tasks = make(map[string]Task)
	if j := json.Get("tasks"); j.Exists() {
		for k, v := range j.Map() {
			p.Tasks[k] = loadTaskJson(k, v)
		}
	}
	return p
}

func extendProject(child Project, path string) (Project, error) {
	parent, err := ParseProject(filepath.Join(child.Cwd, path))
	if err != nil {
		return Project{}, err
	}

	child.Tags = mergeTags(parent.Tags, child.Tags)
	child.Includes = mergeTags(parent.Includes, child.Includes)
	child.EnvFiles = mergeTags(parent.EnvFiles, child.EnvFiles)
	child.Env = mergeEnv(parent.Env, child.Env)

	for taskName, parentTask := range parent.Tasks {
		child.Tasks[taskName] = extendTask(taskName, child.Tasks[taskName], parentTask)
	}

	return child, nil
}
