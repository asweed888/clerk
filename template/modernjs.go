package template

type modernjs struct {
    Lv0 js_level0
    Lv1 js_level1
}

type js_level0 struct {}
type js_level1 struct {}

var Modernjs = &modernjs{}


func (s *js_level0) ExportFile() string {
    return `{{ $appName := .JsConfig.AppName -}}
{{ $codeFileName := .JsConfig.CodeFileName -}}
{{ $codeFileExt := .JsConfig.CodeFileExt -}}
{{ $rootExportTo := .JsConfig.RootExportTo -}}
{{ range .Spec -}}
import {{ .Location }} from "./{{ .Location }}/{{ $codeFileName }}.{{ $codeFileExt }}"{{- printf "\n"}}
{{- end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
export {{ $rootExportTo }}{
{{ range .Spec -}}
{{"    "}}{{ .Location }}{{- printf ",\n"}}
{{- end -}}
}`
}


func (s *js_level0) FullTemplate() string {
    return `{{ $ext := .JsConfig.CodeFileExt -}}
{{ range .Level0.Upstream -}}
import {{ .Name }} from "./{{ .Name }}.{{ $ext }}"
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
export {{ .JsConfig.ExportTo }}{
{{ range .Level0.Upstream -}}
{{"    "}}{{ .Name }}{{- printf ",\n"}}
{{- end -}}
}`
}


func (s *js_level1) FullTemplate() string {
    return `{{ $location := .Level0.Location -}}
/* <location: {{ $location }}.{{ .Level1.Name }} />

{{.Level1.Comment}}

*/

export default {
{{ range .Level1.Methods -}}
{{"    "}}{{ . }}{{- printf ",\n" -}}
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
{{"    "}}{{ . }}{{- printf ",\n" -}}
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
