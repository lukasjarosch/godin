{{ define "response" }}
{{ .Response }} struct {
    {{- range .Returns }}
    {{- if eq .Type "error" }}
        {{ title .Name }} {{ .Type }} `json:"-"`
    {{- else }}
        {{ title .Name }} {{ .ResolveType }} `json:"{{ .Name }}"`
    {{- end }}
    {{- end }}
}
{{ end }}