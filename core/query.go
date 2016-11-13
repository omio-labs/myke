package core

import (
	"strings"
	"errors"
	"fmt"
	"path/filepath"
)

type Query struct {
	Task   string
	Tags   []string
	Params map[string]string
}

func ParseQuery(q string) (Query, error) {
	tokens := strings.SplitN(strings.Trim(q, " ],/"), "[", 2)
	if len(tokens) == 0 || len(tokens) > 2 {
		return Query{}, errors.New(fmt.Sprintf("Bad query: %s", q))
	}

	tasks := strings.Split(strings.Trim(tokens[0], " /"), "/")
	task, tags := tasks[len(tasks) - 1], tasks[:len(tasks) - 1]

	params := make(map[string]string)
	if len(tokens) > 1 {
		for _, pair := range strings.Split(tokens[1], ",") {
			kv := strings.SplitN(pair, "=", 2)
			if len(kv) == 2 {
				params[kv[0]] = kv[1]
			}
		}
	}

	return Query{Task:task, Tags:tags, Params:params}, nil
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

func (q *Query) Matches(p *Project, t *Task) bool {
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
