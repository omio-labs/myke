package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/core"
	"log"
)

var RunCmd = cli.Command{
	Name: "run",
	Usage: "query to execute of format tag1/tag2/project/task[arg1=val1,arg2=val2,...]",
	Action: Run,
}

func Run(c *cli.Context) error {
	qs := make([]string, len(c.Args()))
	for i, v := range c.Args() {
		qs[i] = v
	}

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

	return nil
}
