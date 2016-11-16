package core

import (
	"github.com/flosch/pongo2"
	"strings"
)

func RenderTemplate(tmpl string, vars map[string]string) (string, error) {
	tpl, err := pongo2.FromString(tmpl)
	if err != nil {
		return "", err
	}

	out, err := tpl.Execute(pongo2.Context{"vars": vars})
	if err != nil {
		return "", err
	} else {
		return out, nil
	}
}

func filterRequired(in *(pongo2.Value), param *(pongo2.Value)) (out *(pongo2.Value), err *(pongo2.Error)) {
	if len(strings.TrimSpace(in.String())) == 0 {
		return in, &(pongo2.Error{Sender: "filter:required", ErrorMsg: "required parameter missing"})
	} else {
		return in, nil
	}
}

func init() {
	pongo2.RegisterFilter("required", filterRequired)
}
