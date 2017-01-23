package test

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable{
	{Arg: `task1`, Out: `parent says value_parent_1`},
	{Arg: `task2`, Out: `(?s)parent says value_child_2.*?child says value_child_2`},
	{Arg: `task3`, Out: `child says value_child_3`},
	{Arg: `path`, Out: `PATH is [^:]+mixin/path_child:[^:]+mixin/bin:[^:]+mixin/parent/path_parent:[^:]+mixin/parent/bin`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/mixin", tests)
}
