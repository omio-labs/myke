package cmd

import (
	"log"
	"fmt"
	"io/ioutil"
	"myke/core"
)

func Template(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	rendered, err := core.RenderTemplate(string(bytes), core.OsEnv())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rendered)
}
