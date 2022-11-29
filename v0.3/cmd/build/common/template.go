package common

import (
	"bytes"
	"strings"
	"text/template"
)

type TemplateSpec struct {}

var Template = &TemplateSpec{}

func (s *TemplateSpec) Fill(tmplTxt string, inputData interface{}) string {
    var re bytes.Buffer
    funcMap := template.FuncMap{
        "ToTitle": strings.Title,
    }

    tmpl, err := template.New("clerk").Funcs(funcMap).Parse(tmplTxt)
    if err != nil {
        return ""
    }

    err = tmpl.Execute(&re, inputData) // 置換して標準出力へ
	if err != nil {
		return ""
	}


	return re.String()
}
