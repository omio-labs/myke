package cli

type Workspace struct {
	Cwd      string
	Projects []Project
}

type Project struct {
	Src      string
	Cwd      string
	Name     string
	Desc     string
	Includes []string
	Extends  []string
	Env      map[string]string
	EnvFiles []string
	Tags     []string
	Tasks    map[string]Task
}

type Task struct {
	Name   string
	Desc   string
	Cmd    string
	Before []string
	After  []string
}
