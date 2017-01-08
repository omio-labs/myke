package core

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"text/template"
)

// RenderTemplate renders the given template with env/args map
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
	}
	return s, fmt.Errorf("variable not provided to template")
}
