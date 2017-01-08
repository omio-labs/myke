package test

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable{
	{`task1`, `task1`, `parent says value_parent_1`},
	{`task2`, `task2`, `(?s)parent says value_child_2.*?child says value_child_2`},
	{`task3`, `task3`, `child says value_child_3`},
	{`path`, `path`, `PATH is [^:]+mixin/path_child:[^:]+mixin/bin:[^:]+mixin/parent/path_parent:[^:]+mixin/parent/bin`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/mixin", tests)
}
