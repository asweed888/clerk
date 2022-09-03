package get

var mod1Template = `{{ $mod1_name := .Name -}}
{{ range .Upstreams -}}
import clerk.{{ $mod1_name }}.{{ .Name }}
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
def Clerk(location):
    match location:
{{ range .Upstreams -}}
{{"        "}}case {{ printf "%q" .Name }}: return clerk.{{ $mod1_name }}.{{ .Name }}.Clerk{{- printf "\n" -}}
{{- end -}}`

var mod2Template = `{{ $mod1_name := .Mod1.Name -}}
{{ $mod2_name := .Mod2.Name -}}
{{ range .Mod2.Upstreams -}}
import clerk.{{ $mod1_name }}.{{ $mod2_name }}.{{ .Name }}
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
def Clerk(location):
    match location:
{{ range .Mod2.Upstreams -}}
{{"        "}}case {{ printf "%q" .Name }}: return clerk.{{ $mod1_name }}.{{ $mod2_name }}.{{ .Name }}.Clerk{{- printf "\n" -}}
{{- end -}}`

var mod3Template = `""" <location: {{ .Mod1.Name }}.{{ .Mod2.Name }}.{{ .Mod3.Name }} />

{{.Mod3.Comment}}

"""


def Clerk():
    return{{- printf "\n" -}}`

func ModuleTemplate(modLevel string) string {
    tmpl := map[string]string{
        "mod1": mod1Template,
        "mod2": mod2Template,
        "mod3": mod3Template,
    }
    return tmpl[modLevel]
}
