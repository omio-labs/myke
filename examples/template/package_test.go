package test

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable {
	{`args_from_to`, `args[from=a,to=b]`, `from=a to=b`},
	{`args_from`, `args[from=a]`, `from=a to=something_to`},
	{`args`, `args`, `(?s)template/args: Failed`},
	// {`file`, `file`, `(?s)I am a template.*PARAM1=value1.*PARAM2=value2`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/template", tests)
}
