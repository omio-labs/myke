package core

import (
	"github.com/tidwall/gjson"

	"strings"
)

type Task struct {
	Name   string
	Desc   string
	Cmd    string
	Before string
	After  string
}

func loadTaskJson(name string, json gjson.Result) Task {
	t := Task{}
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
