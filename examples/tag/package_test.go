package test

import (
	. "github.com/omio-labs/myke/examples/util"
	"testing"
)

var tests = []TestTable{
	{Arg: `tag`, Out: `tags1/tag`},
	{Arg: `tag`, Out: `tags2/tag`},
	{Arg: `--dry-run tag`, Out: `(?s)tags1/tag: Will run.*tags2/tag: Will run`},
	{Arg: `tagA/tag`, Out: `tags1 tag`},
	{Arg: `tagA/tag`, Out: `(tags2){0}`},
	{Arg: `tagB/tag`, Out: `tags1/tag`},
	{Arg: `tagB/tag`, Out: `tags2/tag`},
	{Arg: `tagC/tag`, Out: `(tags1){0}`},
	{Arg: `tagC/tag`, Out: `tags2 tag`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/tag", tests)
}
