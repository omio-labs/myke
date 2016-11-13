package cmd

import (
	"myke/core"
	"log"
)

func Run(qs []string) {
	queries, err := core.ParseQueries(qs)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(queries)
	}
}
