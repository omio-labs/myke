package core

import (
	"text/template"
	"bytes"
)

func commandTemplate(tmpl string, params map[string]string) (string, error) {
	t, err := template.New("cmd").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var doc bytes.Buffer
	err = t.Execute(&doc, params)
	if err != nil {
		return "", err
	} else {
		return doc.String(), nil
	}
}
