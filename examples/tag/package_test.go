package test

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable{
	{`tag`, `tag`, `tags1/tag`},
	{`tag`, `tag`, `tags2/tag`},
	{`tag`, `--dry-run tag`, `(?s)tags1/tag: Will run.*tags2/tag: Will run`},
	{`tagA/tag`, `tagA/tag`, `tags1 tag`},
	{`tagA/tag`, `tagA/tag`, `(tags2){0}`},
	{`tagB/tag`, `tagB/tag`, `tags1/tag`},
	{`tagB/tag`, `tagB/tag`, `tags2/tag`},
	{`tagC/tag`, `tagC/tag`, `(tags1){0}`},
	{`tagC/tag`, `tagC/tag`, `tags2 tag`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/tag", tests)
}
