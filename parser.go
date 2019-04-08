package gls

import (
	"html/template"
	"io"
	"time"
)

// TODO ......
const (
	templateOut = `
+------+------+----------+
| Name | Size | Modified |
+------+------+----------+
{{if . -}}
{{- range .}}
| {{.Name}} | {{.SizeInKilobByte}} kb | {{formatDate .ModifiedDate}} |
{{- end }}
{{- end }}
	`
)

func parseOutput(in Assets, out io.Writer) error {
	funcMap := template.FuncMap{
		"formatDate": func(t time.Time) string {
			return t.Format("02 January 2006 15:04:05")
		},
	}

	t := template.New("tmpl").Funcs(funcMap)

	t, err := t.Parse(templateOut)
	if err != nil {
		return err
	}

	err = t.Execute(out, in)
	if err != nil {
		return err
	}

	return nil

}
