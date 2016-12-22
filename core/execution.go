package core

import (
	"github.com/apex/log"
	"os"
	"os/exec"
	"time"
	"strings"
	"errors"
)

type Execution struct {
	Workspace *Workspace
	Query     *Query
	Project   *Project
	Task      *Task
}

func ExecuteQuery(w *Workspace, q Query) error {
	matches := q.Search(w)
	if len(matches) == 0 {
		return errors.New("no task matched: " + q.Raw)
	}
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
	displayName := e.Project.Name + "/" + e.Task.Name
	log.Infof("%v: Running", displayName)

	err := retry(func(attempt int) (bool, error) {
		err := e.executeTask()
		if err != nil && attempt < e.Task.Retries {
			retryMs := time.Duration(e.Task.RetryMs) * time.Millisecond
			log.Debugf("%v: Failed, Retrying %v/%v in %v", displayName, attempt, e.Task.Retries, retryMs)
			time.Sleep(retryMs)
		}
		return attempt < e.Task.Retries, err
	})

	elapsed := time.Since(start)
	if err != nil {
		log.WithError(err).Fatalf("%v: Failed, Took: %v", displayName, elapsed)
	} else {
		log.Infof("%v: Completed, Took: %v", displayName, elapsed)
	}
	return err
}

func (e *Execution) executeTask() error {
	err := e.executeCmd(e.Task.Before)
	if err == nil {
		err = e.executeCmd(e.Task.Cmd)
		if err == nil {
			err = e.executeCmd(e.Task.After)
		}
	}
	return err
}

func (e *Execution) executeCmd(cmd string) error {
	if len(strings.TrimSpace(cmd)) == 0 {
		return nil
	}

	cmd, err := RenderTemplate(cmd, e.Project.Env, e.Query.Params)
	if err != nil {
		return err
	}

	shell := []string{"sh", "-exc",}
	if len(e.Task.Shell) > 0 {
		shell = strings.Split(strings.TrimSpace(e.Task.Shell), " ")
	}
	shell = append(shell, cmd)

	proc := exec.Command(shell[0], shell[1:]...)
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
