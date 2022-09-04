package get

var mod1Template = `{{ range .Upstreams -}}
const _{{.Name}} = require("./{{ .Name }}")
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
module.exports.Clerk = (location) => {
    switch (location) {
{{ range .Upstreams -}}
{{"        "}}case {{ printf "%q" .Name }}: return _{{.Name}}.Clerk{{- printf "\n"}}
{{- end -}}
{{"    "}}}
}`

var mod2Template = `{{ range .Mod2.Upstreams -}}
const _{{.Name}} = require("./{{ .Name }}")
{{ end -}}
{{ printf "\n" -}}
{{ printf "\n" -}}
module.exports.Clerk = (location) => {
    switch (location) {
{{ range .Mod2.Upstreams -}}
{{"        "}}case {{ printf "%q" .Name }}: return _{{.Name}}.Clerk{{- printf "\n" -}}
{{- end -}}
{{"    "}}}
}`

var mod3Template = `/* <location: {{ .Mod1.Name }}.{{ .Mod2.Name }}.{{ .Mod3.Name }} />

{{.Mod3.Comment}}

*/


module.exports.Clerk = () => {}{{- printf "\n" -}}`

func ModuleTemplate(modLevel string) string {
    tmpl := map[string]string{
        "mod1": mod1Template,
        "mod2": mod2Template,
        "mod3": mod3Template,
    }
    return tmpl[modLevel]
}
