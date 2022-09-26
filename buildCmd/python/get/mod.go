package get

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"
)

func ApiRoot(isExport bool) string {
    wd, _ := os.Getwd()
    if isExport {
        return filepath.Base(wd)
    } else {
        return ""
    }
}

var mod0Template = `{{ $apiRoot := .ApiRoot -}}
{{ $location := .Mod0.Location -}}
{{ if $apiRoot -}}
{{ range .Mod0.Upstreams -}}
import {{ $apiRoot }}.clerks.{{ $location }}.{{ .Name }}
{{ end -}}
{{ else -}}
{{ range .Mod0.Upstreams -}}
import clerks.{{ $location }}.{{ .Name }}
{{ end -}}
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
def clerk(location):
    match location:
{{ if $apiRoot -}}
{{ range .Mod0.Upstreams -}}
{{"        "}}case {{ printf "%q" .Name }}: return {{ $apiRoot }}.clerks.{{ $location }}.{{ .Name }}.clerk{{- printf "\n" -}}
{{- end -}}
{{- else -}}
{{ range .Mod0.Upstreams -}}
{{"        "}}case {{ printf "%q" .Name }}: return clerks.{{ $location }}.{{ .Name }}.clerk{{- printf "\n" -}}
{{- end -}}
{{- end -}}`


var mod1Template = `{{ $location := .Mod0.Location -}}
""" <location: {{ $location }}.{{ .Mod1.Name }} />

{{.Mod1.Comment}}

"""


def clerk(location):
    match location:
{{ range .Mod1.Methods -}}
{{"        "}}case {{ printf "%q" . }}: return _{{ . }}{{- printf "\n" -}}
{{- end -}}
# end clerk{{- printf "\n" -}}`


var mod1CommentTemplate = `{{ $location := .Mod0.Location -}}
""" <location: {{ $location }}.{{ .Mod1.Name }} />

{{.Mod1.Comment}}

"""`

var mod1ClerkTemplate = `def clerk(location):
    match location:
{{ range .Mod1.Methods -}}
{{"        "}}case {{ printf "%q" . }}: return _{{ . }}{{- printf "\n" -}}
{{- end -}}
# end clerk`

func ModuleTemplate(modLevel string) string {
    tmpl := map[string]string{
        "mod0": mod0Template,
        "mod1": mod1Template,
        "mod1_comment": mod1CommentTemplate,
        "mod1_clerk": mod1ClerkTemplate,
    }
    return tmpl[modLevel]
}

func CompletedTemplate(tmplTxt string, inputData interface{}) string {
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
