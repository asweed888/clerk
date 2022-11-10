package module1

type TemplateSpec struct {
    Comment T_CommentSpec
    Clerk T_ClerkSpec
    Method T_MethodSpec
}

var Template = &TemplateSpec{}
func (s *TemplateSpec) Get() string {
    return `{{ $location := .Mod0.Location -}}
""" <location: {{ $location }}.{{ .Mod1.Name }} />

{{.Mod1.Comment}}

"""


def clerk(location):
    match location:
{{ range .Mod1.Methods -}}
{{"        "}}case {{ printf "%q" . }}: return _{{ . }}{{- printf "\n" -}}
{{- end -}}
# end clerk{{- printf "\n" -}}`
}


type T_CommentSpec struct {}
func (s *T_CommentSpec) Get() string {
    return `{{ $location := .Mod0.Location -}}
""" <location: {{ $location }}.{{ .Mod1.Name }} />

{{.Mod1.Comment}}

"""`
}


type T_ClerkSpec struct {}
func (s *T_ClerkSpec) Get() string {
    return `def clerk(location):
    match location:
{{ range .Mod1.Methods -}}
{{"        "}}case {{ printf "%q" . }}: return _{{ . }}{{- printf "\n" -}}
{{- end -}}
# end clerk`
}


type T_MethodSpec struct {}
func (s *T_MethodSpec) Get() string {
    return `

def _%s():
    print("this is clerk's default return value")`
}
