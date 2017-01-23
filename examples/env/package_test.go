package test

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable{
	{Arg: `yml`, Out: `envvar from yml is value_from_yml`},
	{Arg: `file_default`, Out: `envvar from myke.env is value_from_myke.env`},
	{Arg: `file_default_local`, Out: `envvar from myke.env.local is value_from_myke.env.local`},
	{Arg: `file_custom`, Out: `envvar from test.env is value_from_test.env`},
	{Arg: `file_custom_local`, Out: `envvar from test.env.local is value_from_test.env.local`},
	{Arg: `path`, Out: `PATH is [^:]+env/path_from_myke.env.local:[^:]+env/path_from_myke.env:[^:]+env/path_from_test.env.local:[^:]+env/path_from_test.env:[^:]+env/path_from_yml:[^:]+env/bin`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/env", tests)
}
