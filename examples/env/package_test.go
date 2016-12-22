package test

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable {
	{`yml`, `yml`, `envvar from yml is value_from_yml`},
	{`file_default`, `file_default`, `envvar from myke.env is value_from_myke.env`},
	{`file_default_local`, `file_default_local`, `envvar from myke.env.local is value_from_myke.env.local`},
	{`file_custom`, `file_custom`, `envvar from test.env is value_from_test.env`},
	{`file_custom_local`, `file_custom_local`, `envvar from test.env.local is value_from_test.env.local`},
	{`path`, `path`, `PATH is [^:]+env/path_from_myke.env.local:[^:]+env/path_from_myke.env:[^:]+env/path_from_test.env.local:[^:]+env/path_from_test.env:[^:]+env/path_from_yml:[^:]+env/bin`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/env", tests)
}
