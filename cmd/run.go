package cmd

import (
	"myke/core"
	"log"
)

func Run(qs []string) {
	queries, err := core.ParseQueries(qs)
	if err != nil {
		log.Fatal(err)
	}

	w := loadWorkspace()
	for _, q := range queries {
		err := core.ExecuteQuery(&w, q)
		if err != nil {
			log.Fatal(err)
		}
	}
}
