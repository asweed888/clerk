package template

type python struct {
    Lv0 py_level0
    Lv1 py_level1
    Lv2 py_level2
}

type py_level0 struct {}
type py_level1 struct {}
type py_level2 struct {}

var Python = &python{}


func (s *py_level0) RootExportFile() string {
    return `from . import (
{{ range .Spec -}}
{{"    "}}{{ .Location | ToTitle }},
{{ end -}}
){{ printf "\n" -}}`
}

func (s *py_level1) FullTemplate() string {
    return `from . import (
{{ range .Level0.Upstream -}}
{{"    "}}{{ .Name | ToTitle }},
{{ end -}}
){{ printf "\n" -}}`
}


func (s *py_level2) FullTemplate() string {
    return `{{ $location := .Level0.Location -}}
""" <location: {{ $location }}.{{ .Level2.Name }} />

{{.Level2.Comment}}

"""{{ printf "\n" -}}`
}


func (s *py_level2) CommentTemplate() string {
    return `{{ $location := .Level0.Location -}}
""" <location: {{ $location }}.{{ .Level2.Name }} />

{{.Level2.Comment}}

"""`
}


func (s *py_level2) MethodTemplate() string {
    return `

def %s():
    print("this is clerk's default return value")`
}
