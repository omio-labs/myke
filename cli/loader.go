package cli

import (
	"github.com/ghodss/yaml"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

func LoadFile(path string) (Project, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return Project{}, err
	} else {
		return LoadYaml(bytes)
	}
}

func LoadYaml(bytes []byte) (Project, error) {
	jsonbytes, err := yaml.YAMLToJSON(bytes)
	if err != nil {
		return Project{}, err
	}

	json := gjson.Parse(string(jsonbytes))
	return LoadProject(json), nil
}

func LoadProject(json gjson.Result) Project {
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
			p.Tasks[k] = LoadTask(k, v)
		}
	}
	return p
}

func LoadTask(name string, json gjson.Result) Task {
	t := Task{}
	t.Name = name

	if j := json.Get("desc"); j.Exists() {
		t.Desc = j.String()
	}
	if j := json.Get("cmd"); j.Exists() {
		t.Cmd = j.String()
	}
	if j := json.Get("before"); j.Exists() {
		for _, s := range j.Array() {
			t.Before = append(t.Before, s.String())
		}
	}
	if j := json.Get("after"); j.Exists() {
		for _, s := range j.Array() {
			t.After = append(t.After, s.String())
		}
	}
	return t
}
