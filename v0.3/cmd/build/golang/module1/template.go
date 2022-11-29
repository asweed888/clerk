package module1

type TemplateSpec struct {
    Comment T_CommentSpec
    Method T_MethodSpec
}

var Template = &TemplateSpec{}
func (s *TemplateSpec) Get() string {
    return `{{ $location := .Mod0.Location -}}
/* <location: {{ $location }}.{{ .Mod1.Name }} />

{{.Mod1.Comment}}

*/
package {{ $location }}

type {{ .Mod1.Name }} struct {}{{- printf "\n" -}}`
}


type T_CommentSpec struct {}
func (s *T_CommentSpec) Get() string {
    return `{{ $location := .Mod0.Location -}}
/* <location: {{ $location }}.{{ .Mod1.Name }} />

{{.Mod1.Comment}}

*/`
}


type T_MethodSpec struct {}
func (s *T_MethodSpec) Get() string {
    return `

func (s *%s) %s() {
    println("this is clerk's default return value")
}`
}
