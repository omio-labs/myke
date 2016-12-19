package core

import (
	"github.com/Masterminds/sprig"
	"text/template"
	"errors"
	"bytes"
	"fmt"
)

func RenderTemplate(tmpl string, env map[string]string, args map[string]string) (string, error) {
	w := new(bytes.Buffer)
	params := union(env, args)
	funcs := template.FuncMap{
		"required": templateRequired,
	}

	tpl, err := template.New("test").
		Funcs(sprig.TxtFuncMap()).
		Funcs(funcs).
		Option("missingkey=zero").
		Parse(tmpl)
	if err != nil {
		return "", err
	}

	err = tpl.Execute(w, params)
	return w.String(), err
}

func templateRequired(s string) (interface{}, error) {
	if len(s) > 0 {
		return s, nil
	} else {
		return s, errors.New(fmt.Sprintf("variable not provided to template"))
	}
}
