package cli

import (
	"io/ioutil"
	"github.com/ghodss/yaml"
	"github.com/tidwall/gjson"
)

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
