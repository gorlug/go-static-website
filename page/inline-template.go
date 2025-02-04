package page

import (
	"bytes"
	"fmt"
	"html/template"
)

func InlineTemplate(templateString string, data any) string {
	tmpl, err := template.New("inline").Delims("$", "$").Parse(templateString)
	if err != nil {
		return fmt.Sprint("Failed to parse template ", err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return fmt.Sprint("Failed to execute template ", err)
	}
	return buf.String()
}
