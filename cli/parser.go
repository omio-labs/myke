package cli

import (
	"io/ioutil"
	"github.com/ghodss/yaml"
	"github.com/tidwall/gjson"
)

func Parse(path string) (Project, error) {
	p := Project{}
	v, err := ParseYaml(path)
	if err != nil {
		return p, err
	}

	p.Name = v.Get("project").String()
	p.Desc = v.Get("desc").String()
	return p, nil
}

func ParseYaml(path string) (gjson.Result, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return gjson.Result{}, err
	}

	json, err := yaml.YAMLToJSON(bytes)
	if err != nil {
		return gjson.Result{}, err
	}

	return gjson.Parse(string(json)), nil
}
