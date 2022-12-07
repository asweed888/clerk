package template

type modernjs struct {
    Lv0 js_level0
    Lv1 js_level1
}

type js_level0 struct {}
type js_level1 struct {}

var Modernjs = &modernjs{}


func (s *js_level0) FullTemplate() string {
    return `{{ $ext := .JsConfig.CodeFileExt -}}
{{ range .Level0.Upstream -}}
import {{ .Name | ToTitle }} from "./{{ .Name }}.{{ $ext }}"
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
const Clerk = {
{{ range .Level0.Upstream -}}
{{"    "}}{{ .Name | ToTitle }}{{- printf ",\n"}}
{{- end -}}
}
{{ printf "\n" -}}
export {{ .JsConfig.ExportTo }}{ Clerk }`
}


func (s *js_level1) FullTemplate() string {
    return `{{ $location := .Level0.Location -}}
/* <location: {{ $location }}.{{ .Level1.Name }} />

{{.Level1.Comment}}

*/

export default {
{{ range .Level1.Methods -}}
{{"    "}}{{ . | ToTitle }}{{- printf ",\n" -}}
{{- end -}}
}
// end export{{- printf "\n" -}}`
}


func (s *js_level1) CommentTemplate() string {
    return `{{ $location := .Level0.Location -}}
/* <location: {{ $location }}.{{ .Level1.Name }} />

{{.Level1.Comment}}

*/`
}


func (s *js_level1) ExportTemplate() string {
    return `export default {
{{ range .Level1.Methods -}}
{{"    "}}{{ . | ToTitle }}{{- printf ",\n" -}}
{{- end -}}
}
// end export`
}


func (s *js_level1) MethodTemplate() string {
    return `

function %s(){
    console.log("this is clerk's default return value")
}`
}
