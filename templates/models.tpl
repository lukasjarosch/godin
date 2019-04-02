package {{ .ServiceName }}

{{ range .Spec.Models.Structs }}
{{- range .Comment }}
// {{ . }}
{{- end }}
type {{ .Name }} struct {
    {{- range .Fields }}
    {{ .Name }} {{ .Type }}
    {{- end }}
}
{{- end }}


{{ range .Spec.Models.Enums }}
{{- range .Comment }}
// {{ . }}
{{- end }}
const (
    {{ enum_body . }}
)
{{- end }}