package core

import (
	"github.com/tidwall/gjson"

	"strings"
)

type Task struct {
	Name   string
	Desc   string
	Cmd    string
	Before []string
	After  []string
}

func loadTaskJson(name string, json gjson.Result) Task {
	t := Task{}
	t.Name = name

	if j := json.Get("desc"); j.Exists() {
		t.Desc = j.String()
	} else {
		t.Desc = ""
	}
	if j := json.Get("cmd"); j.Exists() {
		t.Cmd = j.String()
	} else {
		t.Cmd = ""
	}
	if j := json.Get("before"); j.Exists() {
		for _, s := range j.Array() {
			t.Before = append(t.Before, s.String())
		}
	}
	if j := json.Get("after"); j.Exists() {
		for _, s := range j.Array() {
			t.After = append(t.After, s.String())
		}
	}
	return t
}

func mixinTask(taskName string, child Task, parent Task) Task {
	child.Name = taskName
	if len(strings.TrimSpace(child.Cmd)) == 0 {
		child.Cmd = parent.Cmd
	}
	if len(strings.TrimSpace(child.Desc)) == 0 {
		child.Desc = parent.Desc
	}
	child.Before = mergeTags(child.Before, parent.Before)
	child.After = mergeTags(child.After, parent.After)
	return child
}
