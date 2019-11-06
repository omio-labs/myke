package test

import (
	. "github.com/omio-labs/myke/examples/util"
	"testing"
)

var tests = []TestTable{
	{Arg: `task1`, Out: `parent says value_parent_1`},
	{Arg: `task2`, Out: `(?s)parent says value_child_2.*?child says value_child_2`},
	{Arg: `task3`, Out: `child says value_child_3`},
	{Arg: `path`, Out: `PATH is [^$PLS$]+mixin$PS$path_child$PLS$[^$PLS$]+mixin$PS$bin$PLS$[^$PLS$]+mixin$PS$parent$PS$path_parent$PLS$[^$PLS$]+mixin$PS$parent$PS$bin`},
	{Arg: `-f myke-error.yml`, Err: true, Out: `.*open.*foobar.*`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/mixin", tests)
}
