package cli

import (
  "github.com/cloudfoundry-incubator/candiedyaml"
  "os"
)

func Parse(path string) (map[interface{}]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
  }
  defer file.Close()

	var doc map[interface{}]interface{}
	decoder := candiedyaml.NewDecoder(file)
	err = decoder.Decode(&doc)
	if err != nil {
		return nil, err
  }

  return doc, nil
}
