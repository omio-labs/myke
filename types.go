package main

type Task struct {
	Name string
	Desc string
	Cmd string
	Before []string
	After []string
	Env map[string]string
	Invoked bool
}

type Project struct {
	Src string
	Cwd string
	Name string
	Desc string
	Env map[string]string
	Tags []string
	Tasks []Task
}

type Query struct {
	Task string
	Tags []string
	Params map[string]string
}
