package module0

type TemplateSpec struct {}

var Template = &TemplateSpec{}

func (s *TemplateSpec) Get() string {
    return `{{ range .Mod0.Upstream -}}
import _{{.Name}} from "./{{ .Name }}"
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
export default { clerk }
function clerk(location) {
    switch (location) {
{{ range .Mod0.Upstream -}}
{{"        "}}case {{ printf "%q" .Name }}: return _{{.Name}}.clerk{{- printf "\n"}}
{{- end -}}
{{"    "}}}
}`
}
