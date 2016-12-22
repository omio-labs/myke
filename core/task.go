package core

import (
	"github.com/tidwall/gjson"
	"github.com/apex/log"
	"strings"
	"strconv"
)

type Task struct {
	Name    string
	Desc    string
	Cmd     string
	Before  string
	After   string
	Shell   string
	Retries int
	RetryMs int
}

func loadTaskJson(name string, json gjson.Result) Task {
	t := Task{Retries:0,RetryMs:1000,}
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
	if j := json.Get("retries"); j.Exists() {
		if retries, err := strconv.Atoi(strings.TrimSpace(j.String())); err != nil {
			log.WithError(err).Warnf("invalid retries in task %s", name)
		} else {
			t.Retries = retries
		}
	}
	if j := json.Get("retry_ms"); j.Exists() {
		if retryMs, err := strconv.Atoi(strings.TrimSpace(j.String())); err != nil {
			log.WithError(err).Warnf("invalid retry_ms in task %s", name)
		} else {
			t.RetryMs = retryMs
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
