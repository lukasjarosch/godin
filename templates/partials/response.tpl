{{ block "response" . }}
{{ .Name }}Response struct {
    {{ range .Returns }}
    {{- if eq .Type "error" }}
        {{ title .Name }} {{ .Type }} `json:"-"`
    {{- else }}
        {{ title .Name }} {{ .Type }} `json:"{{ .Name }}"`
    {{- end }}
    {{- end }}
}
{{ end }}