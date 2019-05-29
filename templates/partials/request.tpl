{{ block "request" . }}
{{ .Name }}Request struct {
    {{ range .Params }}
    {{ title .Name }} {{ .Type }} `json:"{{ .Name }}"`
    {{- end }}
}
{{ end }}