package test

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable {
	{`retry_debug`, `--loglevel debug retry`, `(?s)false.*template/retry: Failed, Retrying 1/5 in 10ms.*false.*Retrying 2/5.*false.*Retrying 3/5.*false.*Retrying 4/5.*false.*template/retry: Failed`},
	{`retry`, `retry`, `(Retrying \d+){0}`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/retry", tests)
}
