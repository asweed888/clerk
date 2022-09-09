package get

var clerkRootTemplate = `{{ $location := .Clerk.Location -}}
{{ range .Clerk.Modules -}}
import {{ $location }}.{{ .Name }}
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
def Clerk(location):
    match location:
{{ range .Clerk.Modules -}}
{{"        "}}case {{ printf "%q" .Name }}: return {{ $location }}.{{ .Name }}.Clerk{{- printf "\n" -}}
{{- end -}}
`

var mod1Template = `{{ $location := .Location -}}
{{ $mod1_name := .Mod1.Name -}}
{{ range .Mod1.Upstreams -}}
import {{ $location }}.{{ $mod1_name }}.{{ .Name }}
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
def Clerk(location):
    match location:
{{ range .Mod1.Upstreams -}}
{{"        "}}case {{ printf "%q" .Name }}: return {{ $location }}.{{ $mod1_name }}.{{ .Name }}.Clerk{{- printf "\n" -}}
{{- end -}}`

var mod2Template = `{{ $location := .Location -}}
{{ $mod1_name := .Mod1.Name -}}
{{ $mod2_name := .Mod2.Name -}}
{{ range .Mod2.Upstreams -}}
import {{ $location }}.{{ $mod1_name }}.{{ $mod2_name }}.{{ .Name }}
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
def Clerk(location):
    match location:
{{ range .Mod2.Upstreams -}}
{{"        "}}case {{ printf "%q" .Name }}: return {{ $location }}.{{ $mod1_name }}.{{ $mod2_name }}.{{ .Name }}.Clerk{{- printf "\n" -}}
{{- end -}}`

var mod3Template = `{{ $location := .Location -}}
""" <location: {{ $location }}.{{ .Mod1.Name }}.{{ .Mod2.Name }}.{{ .Mod3.Name }} />

{{.Mod3.Comment}}

"""


def Clerk():
    return{{- printf "\n" -}}`

func ModuleTemplate(modLevel string) string {
    tmpl := map[string]string{
        "clerkRoot": clerkRootTemplate,
        "mod1": mod1Template,
        "mod2": mod2Template,
        "mod3": mod3Template,
    }
    return tmpl[modLevel]
}
