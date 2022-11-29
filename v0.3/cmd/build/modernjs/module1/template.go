package module1

type TemplateSpec struct {
    Comment T_CommentSpec
    Clerk T_ClerkSpec
    Method T_MethodSpec
}

var Template = &TemplateSpec{}
func (s *TemplateSpec) Get() string {
    return `{{ $location := .Mod0.Location -}}
/* <location: {{ $location }}.{{ .Mod1.Name }} />

{{.Mod1.Comment}}

*/

export default {
{{ range .Mod1.Methods -}}
{{"    "}}{{ . | ToTitle }}{{- printf ",\n" -}}
{{- end -}}
}
// end export{{- printf "\n" -}}`
}


type T_CommentSpec struct {}
func (s *T_CommentSpec) Get() string {
    return `{{ $location := .Mod0.Location -}}
/* <location: {{ $location }}.{{ .Mod1.Name }} />

{{.Mod1.Comment}}

*/`
}


type T_ClerkSpec struct {}
func (s *T_ClerkSpec) Get() string {
    return `export default {
{{ range .Mod1.Methods -}}
{{"    "}}{{ . | ToTitle }}{{- printf ",\n" -}}
{{- end -}}
}
// end export`
}


type T_MethodSpec struct {}
func (s *T_MethodSpec) Get() string {
    return `

function %s(){
    console.log("this is clerk's default return value")
}`
}
