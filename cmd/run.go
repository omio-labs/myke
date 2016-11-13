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
		ms := q.Search(&w)
		for _, m := range ms {
			core.Execute(&w, &m.Project, &m.Task)
		}
	}
}
