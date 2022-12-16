package template

type golang struct {
    Lv0 go_level0
    Lv1 go_level1
}

type go_level0 struct {}
type go_level1 struct {}

var Golang = &golang{}


func (s *go_level0) ExportGo() string {
    return `{{ $appName := .GoConfig.AppName -}}
{{ $goModName := .GoConfig.GoModName -}}
package {{ $appName }}
{{ printf "\n" -}}
import (
{{ range .Spec -}}
{{"    "}}"{{ $goModName }}/{{ $appName }}/{{ .Location }}"{{- printf "\n"}}
{{- end -}}
){{- printf "\n" -}}
{{ printf "\n" -}}
{{ range .Spec -}}
var {{ .Location | ToTitle }} = &{{ .Location }}Mod{}{{- printf "\n"}}
{{- end -}}
{{ printf "\n" -}}
{{ range .Spec -}}
{{ $location := .Location -}}
type {{ $location }}Mod struct {
{{ range .Upstream -}}
{{"    "}}{{ .Name | ToTitle }} {{ $location }}.{{ .Name | ToTitle }}Mod{{- printf "\n"}}
{{- end -}}
}{{- printf "\n" -}}
{{- end -}}`
}


func (s *go_level1) FullTemplate() string {
    return `{{ $location := .Level0.Location -}}
{{ $upstream := .Level1.Name -}}
{{ $isExport := .IsExport -}}
/* <location: {{ $location }}.{{ .Level1.Name }} />

{{.Level1.Comment}}

*/
package {{ $location }}

type {{ ToExportable $isExport $upstream }}Mod struct {}
var {{ .Level1.Name | ToTitle }} = &{{ ToExportable $isExport $upstream }}Mod{}{{- printf "\n" -}}`
}


func (s *go_level1) CommentTemplate() string {
    return `{{ $location := .Level0.Location -}}
/* <location: {{ $location }}.{{ .Level1.Name }} />

{{.Level1.Comment}}

*/`
}


func (s *go_level1) MethodTemplate() string {
    return `

func (s *%sMod) %s() {
    println("this is clerk's default return value")
}`
}
