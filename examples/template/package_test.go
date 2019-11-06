package test

import (
	. "github.com/omio-labs/myke/examples/util"
	"testing"
)

var tests = []TestTable{
	{Arg: `args`, Out: `(?s)template/args: Failed`, Err: true},
	{Arg: `args --from=a`, Out: `from=a to=something_to`},
	{Arg: `args --from=a --to=b`, Out: `from=a to=b`},
	{Arg: `args --from=a args --from=b`, Out: `(?s).*from=a to=something_to.*from=b to=something_to`},
	{Arg: `envs`, Err: true, Out: `variable not provided to template`},
	{Arg: `envs PARAM1=value1 PARAM2=value2`, Out: `(?s).*PARAM1=value1 PARAM2=value2`},
	{Arg: `--template template.tpl`, Out: `(?s)^I am a template.*TEST=TEST.*`},
	{Arg: `--template foobar.tpl`, Err: true, Out: `error rendering template`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/template", tests)
}
