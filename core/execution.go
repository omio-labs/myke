package core

import (
	"fmt"
	"time"
	"os"
	"os/exec"
)

type Execution struct {
	Workspace *Workspace
	Project   *Project
	Task      *Task
}

func ExecuteQuery(w *Workspace, q Query) error {
	matches := q.Search(w)
	for _, match := range matches {
		e := Execution{Workspace:w, Project:&match.Project, Task:&match.Task}
		err := e.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Execution) Execute() error {
	start := time.Now()
	fmt.Printf("%v/%v: Running\n", e.Project.Name, e.Task.Name)

	err := e.ExecuteDependent(e.Task.Before)
	if err == nil {
		err = e.ExecuteSelf()
		if err == nil {
			err = e.ExecuteDependent(e.Task.After)
		}
	}

	elapsed := time.Since(start)
	if err != nil {
		fmt.Printf("%v/%v: Failed, Took: %v\n", e.Project.Name, e.Task.Name, elapsed)
	} else {
		fmt.Printf("%v/%v: Completed, Took: %v\n", e.Project.Name, e.Task.Name, elapsed)
	}

	return err
}

func (e *Execution) ExecuteSelf() error {
	return e.Cmd().Run()
}

func (e *Execution) ExecuteDependent(qs []string) error {
	for _, q := range qs {
		query, err := ParseQuery(q)
		if err != nil {
			return err
		}
		if len(query.Tags) == 0 {
			query.Tags = []string{e.Project.Name}
		}
		if !query.Match(e.Project, e.Task) {
			err = ExecuteQuery(e.Workspace, query)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *Execution) Cmd() *exec.Cmd {
	cmd := exec.Command("sh", "-exc", e.Task.Cmd)
	cmd.Dir = e.Project.Cwd
	cmd.Env = envList(e.Project.Env)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func envList(env map[string]string) []string {
	env["PATH"] = env["PATH"] + string(os.PathListSeparator) + os.Getenv("PATH")
	envList := make([]string, len(env))
	for k, v := range env {
		envList = append(envList, fmt.Sprintf("%v=%v", k, v))
	}
	return envList
}
