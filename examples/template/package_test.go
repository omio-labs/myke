package test

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable{
	{`args`, `args`, `(?s)template/args: Failed`},
	{`args_from`, `args --from=a`, `from=a to=something_to`},
	{`args_from_to`, `args --from=a --to=b`, `from=a to=b`},
	{`args_multiple_tasks`, `args --from=a args --from=b`, `(?s).*from=a to=something_to.*from=b to=something_to`},
	// Cannot invoke myke subcommand in a test
	// {`file`, `file`, `(?s)I am a template.*PARAM1=value1.*PARAM2=value2`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/template", tests)
}
