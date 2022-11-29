package module0

type TemplateSpec struct {}

var Template = &TemplateSpec{}

func (s *TemplateSpec) Get() string {
    return `{{ $location := .Mod0.Location -}}
package {{ $location }}
{{ printf "\n" -}}
type clerk struct {
{{ range .Mod0.Upstream -}}
{{"    "}}{{ .Name | ToTitle }} {{ .Name }}
{{ end -}}
}
{{ printf "\n" -}}
var Clerk = &clerk{}`
}
