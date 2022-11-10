package module0

type TemplateSpec struct {}

var Template = &TemplateSpec{}

func (s *TemplateSpec) Get() string {
    return `from . import (
{{ range .Mod0.Upstream -}}
{{"    "}}{{ .Name }},
{{ end -}}
){{ printf "\n" -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
def clerk(location):
    match location:
{{ range .Mod0.Upstream -}}
{{"        "}}case {{ printf "%q" .Name }}: return {{ .Name }}.clerk{{- printf "\n" -}}
{{- end -}}`
}
