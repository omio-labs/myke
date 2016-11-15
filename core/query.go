package core

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

type Query struct {
	Raw    string
	Task   string
	Tags   []string
	Params map[string]string
}

type QueryMatch struct {
	Project Project
	Task    Task
}

func ParseQuery(q string) (Query, error) {
	tokens := strings.SplitN(strings.Trim(q, " ],/"), "[", 2)
	if len(tokens) == 0 || len(tokens) > 2 {
		return Query{}, errors.New(fmt.Sprintf("Bad query: %s", q))
	}

	tasks := strings.Split(strings.Trim(tokens[0], " /"), "/")
	task, tags := tasks[len(tasks)-1], tasks[:len(tasks)-1]

	params := make(map[string]string)
	if len(tokens) > 1 {
		for _, pair := range strings.Split(tokens[1], ",") {
			kv := strings.SplitN(pair, "=", 2)
			if len(kv) == 2 {
				params[kv[0]] = kv[1]
			}
		}
	}

	return Query{Raw: q, Task: task, Tags: tags, Params: params}, nil
}

func ParseQueries(qs []string) ([]Query, error) {
	queries := make([]Query, len(qs))
	for i, q := range qs {
		if query, err := ParseQuery(q); err != nil {
			return nil, err
		} else {
			queries[i] = query
		}
	}
	return queries, nil
}

func (q *Query) Search(w *Workspace) []QueryMatch {
	matches := []QueryMatch{}
	for _, p := range w.Projects {
		for _, t := range p.Tasks {
			if q.Match(&p, &t) {
				match := QueryMatch{Project: p, Task: t}
				matches = append(matches, match)
			}
		}
	}
	return matches
}

func (q *Query) Match(p *Project, t *Task) bool {
	for _, tag := range q.Tags {
		projectMatch, _ := filepath.Match(tag, p.Name)
		for _, projectTag := range p.Tags {
			tagMatch, _ := filepath.Match(tag, projectTag)
			projectMatch = projectMatch || tagMatch
		}
		if !projectMatch {
			return false
		}
	}
	taskMatch, _ := filepath.Match(q.Task, t.Name)
	return taskMatch
}
