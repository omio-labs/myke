package test

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable{
	{Arg: `before`, Out: `running before`},
	{Arg: `after`, Out: `running after`},
	{Arg: `error`, Out: `(?s)foobar.*there was an error`, Err: true},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/hooks", tests)
}
