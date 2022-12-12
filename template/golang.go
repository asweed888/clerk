package template

type golang struct {
    Lv0 go_level0
    Lv1 go_level1
}

type go_level0 struct {}
type go_level1 struct {}

var Golang = &golang{}


func (s *go_level1) FullTemplate() string {
    return `{{ $location := .Level0.Location -}}
/* <location: {{ $location }}.{{ .Level1.Name }} />

{{.Level1.Comment}}

*/
package {{ $location }}

type {{ .Level1.Name }}Mod struct {}
var {{ .Level1.Name | ToTitle }} = &{{ .Level1.Name }}Mod{}{{- printf "\n" -}}`
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
