package core

import (
	"github.com/apex/log"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"time"
)

// Task reprents a parsed task
type Task struct {
	Name       string
	Desc       string
	Cmd        string
	Before     string
	After      string
	Shell      string
	Retry      int
	RetryDelay time.Duration
}

const defaultDelay = time.Duration(1) * time.Second

func loadTaskJSON(name string, json gjson.Result) Task {
	t := Task{Retry: 0, RetryDelay: defaultDelay}
	t.Name = name

	if j := json.Get("desc"); j.Exists() {
		t.Desc = strings.TrimSpace(j.String())
	}
	if j := json.Get("cmd"); j.Exists() {
		t.Cmd = strings.TrimSpace(j.String())
	}
	if j := json.Get("before"); j.Exists() {
		t.Before = strings.TrimSpace(j.String())
	}
	if j := json.Get("after"); j.Exists() {
		t.After = strings.TrimSpace(j.String())
	}
	if j := json.Get("shell"); j.Exists() {
		t.Shell = strings.TrimSpace(j.String())
	}
	if j := json.Get("retry"); j.Exists() {
		if retry, err := strconv.Atoi(strings.TrimSpace(j.String())); err != nil {
			log.WithError(err).Warnf("invalid retry in task %s", name)
		} else {
			t.Retry = retry
		}
	}
	if j := json.Get("retry_delay"); j.Exists() {
		if delay, err := time.ParseDuration(strings.TrimSpace(j.String())); err != nil {
			log.WithError(err).Warnf("invalid retry_delay in task %s", name)
		} else {
			t.RetryDelay = delay
		}
	}
	return t
}

func mixinTask(taskName string, child Task, parent Task) Task {
	child.Name = taskName
	if len(child.Cmd) == 0 {
		child.Cmd = parent.Cmd
	}
	if len(child.Desc) == 0 {
		child.Desc = parent.Desc
	}
	child.Before = strings.TrimSpace(child.Before + "\n" + parent.Before)
	child.After = strings.TrimSpace(child.After + "\n" + parent.After)
	return child
}
