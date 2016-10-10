package main

import (
  "github.com/cloudfoundry-incubator/candiedyaml"
  "os"
)

type Holder interface{}

func Parse(path string) (*Project, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
  }
  defer file.Close()

	doc := new(Holder)
	decoder := candiedyaml.NewDecoder(file)
	err = decoder.Decode(doc)
	if err != nil {
		return nil, err
  }

  return nil, nil
}
