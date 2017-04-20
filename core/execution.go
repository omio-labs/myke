package core

import (
	"errors"
	"github.com/apex/log"
	"github.com/kardianos/osext"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Execution represents a task and context being invoked
type Execution struct {
	Workspace *Workspace
	Query     *Query
	Project   *Project
	Task      *Task
	DryRun    bool
}

// ExecuteQuery executes the given query in the workspace
func ExecuteQuery(w *Workspace, q Query, dryRun bool) error {
	matches := q.search(w)
	if len(matches) == 0 {
		return errors.New("no task matched: " + q.Raw)
	}
	for _, match := range matches {
		e := Execution{
			Workspace: w,
			Query:     &q,
			Project:   &match.Project,
			Task:      &match.Task,
			DryRun:    dryRun,
		}
		err := e.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

// Execute executes the current task
func (e *Execution) Execute() error {
	start := time.Now()
	displayName := e.Project.Name + "/" + e.Task.Name
	if e.DryRun {
		log.Infof("%v: Will run", displayName)
		return nil
	}
	log.Infof("%v: Running", displayName)

	err := retry(func(attempt int) (bool, error) {
		err := e.executeTask()
		if err != nil && attempt < e.Task.Retry {
			log.Debugf("%v: Failed, Retrying %v/%v in %v", displayName, attempt, e.Task.Retry, e.Task.RetryDelay)
			time.Sleep(e.Task.RetryDelay)
		}
		return attempt < e.Task.Retry, err
	})

	elapsed := time.Since(start)
	if err != nil {
		log.Errorf("%v/%v: Failed, Took: %v", e.Project.Name, e.Task.Name, elapsed)
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

	if err != nil && len(e.Task.Error) > 0 {
		e.executeCmd(e.Task.Error)
	}
	return err
}

func (e *Execution) executeCmd(cmd string) error {
	if len(strings.TrimSpace(cmd)) == 0 {
		return nil
	}

	env := e.env()
	cmd, err := RenderTemplate(cmd, e.Project.Env, e.Query.Params)
	if err != nil {
		return err
	}
	cmd = os.Expand(cmd, func(key string) string {
		return env[key]
	})

	e.beforeExecuteCmd(cmd, env)

	shell := executionShell()
	if len(e.Task.Shell) > 0 {
		shell = strings.Split(strings.TrimSpace(e.Task.Shell), " ")
	}
	shell = append(shell, cmd)

	proc := exec.Command(shell[0], shell[1:]...)
	proc.Dir = e.Project.Cwd
	proc.Env = mapToSlice(env)
	proc.Stdin = os.Stdin
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr
	return proc.Run()
}

func (e *Execution) env() map[string]string {
	myke, _ := osext.Executable()
	extraEnv := map[string]string{
		"myke":         myke,
		"MYKE_PROJECT": e.Project.Name,
		"MYKE_TASK":    e.Task.Name,
		"MYKE_CWD":     e.Project.Cwd,
	}
	env := mergeEnv(mergeEnv(e.Project.Env, extraEnv), OsEnv())
	env["PATH"] = strings.Join([]string{env["PATH"], os.Getenv("PATH")}, pathSep)
	return env
}
