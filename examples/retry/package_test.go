package test

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable{
	{Arg: `-v=0 retry`, Out: `(Running){0}`, Err: true},
	{Arg: `-v=0 retry`, Out: `(Failed){0}`, Err: true},
	{Arg: `-v=0 retry`, Out: `(Running){0}`, Err: true},
	{Arg: `-v=1 retry`, Out: `Failed`, Err: true},
	{Arg: `retry`, Out: `(Retrying \d+){0}`, Err: true},
	{Arg: `-v=5 retry`, Out: `(?s)false.*retry/retry: Failed, Retrying 1/5 in 10ms.*false.*Retrying 2/5.*false.*Retrying 3/5.*false.*Retrying 4/5.*false.*retry/retry: Failed`, Err: true},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/retry", tests)
}
