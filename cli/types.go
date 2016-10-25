package cli

type Project struct {
	Src string
	Name string
	Desc string
	Includes []string
	Extends []string
	Env map[string]string
	EnvFiles []string
	Tags []string
	Tasks []Task
}

type Task struct {
	Name string
	Desc string
	Cmd string
	Before []string
	After []string
}
