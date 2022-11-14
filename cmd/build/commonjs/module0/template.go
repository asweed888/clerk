package module0

type TemplateSpec struct {}

var Template = &TemplateSpec{}

func (s *TemplateSpec) Get() string {
    return `{{ range .Mod0.Upstream -}}
const _{{.Name}} = require("./{{ .Name }}")
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
module.exports.clerk = (location) => {
    switch (location) {
{{ range .Mod0.Upstream -}}
{{"        "}}case {{ printf "%q" .Name }}: return _{{.Name}}.clerk{{- printf "\n"}}
{{- end -}}
{{"    "}}}
}`
}
