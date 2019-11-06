package test

import (
	. "github.com/omio-labs/myke/examples/util"
	"testing"
)

var tests = []TestTable{
	{Arg: `-v=5 retry`, Out: `(?s)false.*retry/retry: Failed, Retrying 1/5 in 10ms.*false.*Retrying 2/5.*false.*Retrying 3/5.*false.*Retrying 4/5.*false.*retry/retry: Failed`, Err: true},
	{Arg: `retry`, Out: `(Retrying \d+){0}`, Err: true},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/retry", tests)
}
