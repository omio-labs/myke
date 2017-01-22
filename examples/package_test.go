package test

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable{
	{`heading`, ``, `(?m)^\s*PROJECT\s*\|\s*TAGS\s*\|\s*TASKS\s*$`},
	{`env`, ``, `(?m)^\s*env\s*\|\s*\|\s*file_custom, file_custom_local, file_default, file_default_local, path, yml\s*$`},
	{`hooks`, ``, `(?m)^\s*hooks\s*\|\s*\|\s*after, before, before_after\s*$`},
	{`mixin`, ``, `(?m)^\s*mixin\s*\|\s*\|\s*path, task1, task2, task3\s*$`},
	{`retry`, ``, `(?m)^\s*retry\s*\|\s*\|\s*retry\s*$`},
	{`tags1`, ``, `(?m)^\s*tags1\s*\|\s*tagA, tagB\s*\|\s*tag\s*$`},
	{`tags2`, ``, `(?m)^\s*tags2\s*\|\s*tagB, tagC\s*\|\s*tag\s*$`},
	{`template`, ``, `(?m)^\s*template\s*\|\s*\|\s*args, file\s*$`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples", tests)
}
