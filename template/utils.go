package template

import (
	"bytes"
	"strings"
	"text/template"
)

type utils struct {}

var Utils = &utils{}

func (s *utils) FillTemplate(tmplString string, inputData interface{}) string {
    var re bytes.Buffer
    funcMap := template.FuncMap{
        "ToTitle": strings.Title,
    }

    tmpl, err := template.New("clerk").Funcs(funcMap).Parse(tmplString)
    if err != nil {
        return ""
    }

    err = tmpl.Execute(&re, inputData) // 置換して標準出力へ
	if err != nil {
		return ""
	}


	return re.String()
}
