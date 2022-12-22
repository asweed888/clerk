package template

type rust struct {
    Lv0 rs_level0
    Lv1 rs_level1
}

type rs_level0 struct {}
type rs_level1 struct {}

var Rust = &rust{}

func (s *rs_level0) RootExportFile() string {
    return `{{ range .Spec -}}
pub mod {{ .Location }};{{- printf "\n"}}
{{- end -}}`
}


func (s *rs_level0) ExportFile() string {
    return `{{ range .Level0.Upstream -}}
pub mod {{ .Name }};{{- printf "\n"}}
{{- end -}}`
}


func (s *rs_level1) CommentTemplate() string {
    return `{{ $location := .Level0.Location -}}
/* <location: {{ $location }}.{{ .Level1.Name }} />

{{.Level1.Comment}}

*/`
}


func (s *rs_level1) MethodTemplate() string {
    return `

pub fn %s(){
    println!("this is clerk's default return value")
}`
}
