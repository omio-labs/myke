package core

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"strings"
)

type Execution struct {
	Workspace *Workspace
	Query     *Query
	Project   *Project
	Task      *Task
}

func ExecuteQuery(w *Workspace, q Query) error {
	matches := q.Search(w)
	for _, match := range matches {
		e := Execution{
			Workspace: w,
			Query:     &q,
			Project:   &match.Project,
			Task:      &match.Task,
		}
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
	cmd, err := RenderTemplate(e.Task.Cmd, e.Project.Env, e.Query.Params)
	if err != nil {
		return err
	}

	proc := exec.Command("sh", "-exc", cmd)
	proc.Dir = e.Project.Cwd
	proc.Env = e.EnvList()
	proc.Stdin = os.Stdin
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr
	return proc.Run()
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

func (e *Execution) Env() map[string]string {
	extraEnv := map[string]string{
		"MYKE_PROJECT": e.Project.Name,
		"MYKE_TASK": e.Task.Name,
		"MYKE_CWD": e.Project.Cwd,
	}
	env := mergeEnv(mergeEnv(e.Project.Env, extraEnv), OsEnv())
	env["PATH"] = strings.Join([]string{ env["PATH"], os.Getenv("PATH") }, PathSep)
	return env
}

func (e *Execution) EnvList() []string {
	env := e.Env()
	envList := make([]string, len(env))
	for k, v := range env {
		envList = append(envList, fmt.Sprintf("%v=%v", k, v))
	}
	return envList
}
