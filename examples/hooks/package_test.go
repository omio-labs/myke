package test

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable {
	{`before`, `before`, `running before`},
	{`after`, `after`, `running after`},
	// {`before_after`, `before_after`, `running before.+running cmd.+running after`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/hooks", tests)
}
