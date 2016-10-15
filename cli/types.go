package cli

type Task struct {
	Name string
	Desc string
	Cmd string
	Before []string
	After []string
	Env map[string]string
	Invoked bool
}

type Query struct {
	Task string
	Tags []string
	Params map[string]string
}
