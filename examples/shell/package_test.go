package test

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable{
	{Arg: `subshell`, Out: `subshell works`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/shell", tests)
}
