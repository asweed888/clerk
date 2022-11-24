package module0

type TemplateSpec struct {}

var Template = &TemplateSpec{}

func (s *TemplateSpec) Get() string {
    return `{{ range .Mod0.Upstream -}}
import {{ .Name | ToTitle }} from "./{{ .Name }}.js"
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
export {{ .JsConfig.ExportTo }}{
    Clerk: {
{{ range .Mod0.Upstream -}}
{{"        "}}{{ .Name | ToTitle }}{{- printf ",\n"}}
{{- end -}}
{{"    "}}}
}`
}
