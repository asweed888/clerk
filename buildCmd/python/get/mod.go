package get

import (
	"os"
	"path/filepath"
)

func ApiRoot(isExport bool) string {
    wd, _ := os.Getwd()
    if isExport {
        return filepath.Base(wd)
    } else {
        return ""
    }
}

var mod0Template = `{{ $apiRoot := .ApiRoot -}}
{{ $location := .Mod0.Location -}}
{{ if $apiRoot -}}
{{ range .Mod0.Services -}}
import {{ $apiRoot }}.clerk.{{ $location }}.{{ .Name }}
{{ end -}}
{{ else -}}
{{ range .Mod0.Services -}}
import clerk.{{ $location }}.{{ .Name }}
{{ end -}}
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
def Clerk(location):
    match location:
{{ if $apiRoot -}}
{{ range .Mod0.Services -}}
{{"        "}}case {{ printf "%q" .Name }}: return {{ $apiRoot }}.clerk.{{ $location }}.{{ .Name }}.Clerk{{- printf "\n" -}}
{{- end -}}
{{- else -}}
{{ range .Mod0.Services -}}
{{"        "}}case {{ printf "%q" .Name }}: return clerk.{{ $location }}.{{ .Name }}.Clerk{{- printf "\n" -}}
{{- end -}}
{{- end -}}`


var mod1Template = `{{ $apiRoot := .ApiRoot -}}
{{ $location := .Mod0.Location -}}
{{ $mod1_name := .Mod1.Name -}}
{{ if $apiRoot -}}
{{ range .Mod1.Upstreams -}}
import {{ $apiRoot }}.clerk.{{ $location }}.{{ $mod1_name }}.{{ .Name }}
{{ end -}}
{{ else -}}
{{ range .Mod1.Upstreams -}}
import clerk.{{ $location }}.{{ $mod1_name }}.{{ .Name }}
{{ end -}}
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
def Clerk(location):
    match location:
{{ if $apiRoot -}}
{{ range .Mod1.Upstreams -}}
{{"        "}}case {{ printf "%q" .Name }}: return {{ $apiRoot }}.clerk.{{ $location }}.{{ $mod1_name }}.{{ .Name }}.Clerk{{- printf "\n" -}}
{{- end -}}
{{- else -}}
{{ range .Mod1.Upstreams -}}
{{"        "}}case {{ printf "%q" .Name }}: return clerk.{{ $location }}.{{ $mod1_name }}.{{ .Name }}.Clerk{{- printf "\n" -}}
{{- end -}}
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
