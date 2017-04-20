package test

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable{
	{Arg: `args`, Out: `(?s)template/args: Failed`, Err: true},
	{Arg: `args --from=a`, Out: `from=a to=something_to`},
	{Arg: `args --from=a --to=b`, Out: `from=a to=b`},
	{Arg: `args --from=a args --from=b`, Out: `(?s).*from=a to=something_to.*from=b to=something_to`},
	{Arg: `envs`, Out: `(?s).*PARAM1=value2 PARAM2=value2`},
	// Cannot invoke myke subcommand in a test
	// {Arg:`file`, Out:`(?s)I am a template.*PARAM1=value1.*PARAM2=value2`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/template", tests)
}
