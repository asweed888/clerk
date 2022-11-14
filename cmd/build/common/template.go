package common

import (
	"bytes"
	"text/template"
)

type TemplateSpec struct {}

var Template = &TemplateSpec{}

func (s *TemplateSpec) Fill(tmplTxt string, inputData interface{}) string {
    var re bytes.Buffer

    tmpl, err := template.New("clerk").Parse(tmplTxt)
    if err != nil {
        return ""
    }

    err = tmpl.Execute(&re, inputData) // 置換して標準出力へ
	if err != nil {
		return ""
	}


	return re.String()
}
