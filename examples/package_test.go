package test

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable{
	{`heading`, ``, `PROJECT\s*\|\s*TAGS\s*\|\s*TASKS`},
	// TODO: tasks are not sorted
	// {`env`, ``, `env\s*\|\s*\|\s*env`},
	// {`tags1`, ``, `tags1\s*\|\s*tagA, tagB\s*\|\s*tag`},
	// {`tags2`, ``, `tags2\s*\|\s*tagB, tagC\s*\|\s*tag`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples", tests)
}
