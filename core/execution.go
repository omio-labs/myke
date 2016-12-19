package core

import (
	"github.com/apex/log"
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
	log.Infof("%v/%v: Running\n", e.Project.Name, e.Task.Name)

	err := e.ExecuteCmd(e.Task.Before)
	if err == nil {
		err = e.ExecuteCmd(e.Task.Cmd)
		if err == nil {
			err = e.ExecuteCmd(e.Task.After)
		}
	}

	elapsed := time.Since(start)
	if err != nil {
		log.WithError(err).Fatalf("%v/%v: Failed, Took: %v\n", e.Project.Name, e.Task.Name, elapsed)
	} else {
		log.Infof("%v/%v: Completed, Took: %v\n", e.Project.Name, e.Task.Name, elapsed)
	}

	return err
}

func (e *Execution) ExecuteCmd(cmd string) error {
	if len(strings.TrimSpace(cmd)) == 0 {
		return nil
	}

	cmd, err := RenderTemplate(cmd, e.Project.Env, e.Query.Params)
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
		envList = append(envList, k+"="+v)
	}
	return envList
}
