package cli

type Project struct {
	Src string
	Cwd string
	Name string
	Desc string
	Env map[string]string
	Tags []string
	Tasks []Task
}

