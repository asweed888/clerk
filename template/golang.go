package template

type golang struct {
    Lv0 go_level0
    Lv1 go_level1
}

type go_level0 struct {}
type go_level1 struct {}

var Golang = &golang{}



// func (s *go_level0) FullTemplate() string {
//     return `{{ $location := .Level0.Location -}}
// package {{ $location }}
// {{ printf "\n" -}}
// type clerk struct {
// {{ range .Level0.Upstream -}}
// {{"    "}}{{ .Name | ToTitle }} {{ .Name }}_mod
// {{ end -}}
// }
// {{ printf "\n" -}}
// var Clerk = &clerk{}`
// }


func (s *go_level1) FullTemplate() string {
    return `{{ $location := .Level0.Location -}}
/* <location: {{ $location }}.{{ .Level1.Name }} />

{{.Level1.Comment}}

*/
package {{ $location }}

type {{ .Level1.Name }}_mod struct {}
var {{ .Level1.Name | ToTitle }} = &{{ .Level1.Name }}_mod{}{{- printf "\n" -}}`
}


func (s *go_level1) CommentTemplate() string {
    return `{{ $location := .Level0.Location -}}
/* <location: {{ $location }}.{{ .Level1.Name }} />

{{.Level1.Comment}}

*/`
}


func (s *go_level1) MethodTemplate() string {
    return `

func (s *%s_mod) %s() {
    println("this is clerk's default return value")
}`
}
