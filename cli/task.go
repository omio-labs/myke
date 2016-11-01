package cli

type Task struct {
	Name   string
	Desc   string
	Cmd    string
	Before []string
	After  []string
}
