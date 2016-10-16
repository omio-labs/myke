package cli

import (
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Project struct {
	Src string
	Name string `yaml:"project"`
	Desc string
	Includes []string
	Extends []string
	Env map[string]string
	EnvFiles []string
	Tags []string
	Tasks map[string]*Task
}

type Task struct {
	Name string
	Desc string
	Cmd string
	Before []string
	After []string
}

func ParseFile(src string, p *Project) (error) {
	abssrc, err := filepath.Abs(src)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	p.Src = abssrc
	return Parse(bytes, p)
}

func Parse(in []byte, p *Project) (error) {
	err := yaml.Unmarshal(in, &p)
	if err != nil {
		return err
	}

  for key, task := range p.Tasks {
  	task.Name = key
  }
  return nil
}
