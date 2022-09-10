package get

var mod0Template = `{{ $location := .Mod0.Location -}}
{{ range .Mod0.Modules -}}
import clerk.{{ $location }}.{{ .Name }}
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
def Clerk(location):
    match location:
{{ range .Mod0.Modules -}}
{{"        "}}case {{ printf "%q" .Name }}: return clerk.{{ $location }}.{{ .Name }}.Clerk{{- printf "\n" -}}
{{- end -}}`


var mod1Template = `{{ $location := .Mod0.Location -}}
{{ $mod1_name := .Mod1.Name -}}
{{ range .Mod1.Upstreams -}}
import clerk.{{ $location }}.{{ $mod1_name }}.{{ .Name }}
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
def Clerk(location):
    match location:
{{ range .Mod1.Upstreams -}}
{{"        "}}case {{ printf "%q" .Name }}: return clerk.{{ $location }}.{{ $mod1_name }}.{{ .Name }}.Clerk{{- printf "\n" -}}
{{- end -}}`


var mod2Template = `{{ $location := .Mod0.Location -}}
""" <location: {{ $location }}.{{ .Mod1.Name }}.{{ .Mod2.Name }} />

{{.Mod2.Comment}}

"""


def Clerk():
    return{{- printf "\n" -}}`

func ModuleTemplate(modLevel string) string {
    tmpl := map[string]string{
        "mod0": mod0Template,
        "mod1": mod1Template,
        "mod2": mod2Template,
    }
    return tmpl[modLevel]
}
