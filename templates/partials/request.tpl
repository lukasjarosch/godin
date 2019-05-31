{{ define "request" }}
{{ .Name }}Request struct {
    {{ range .Params }}
        {{- if ne .Name "ctx" }}
        {{ title .Name }} {{ .Type }} `json:"{{ .Name }}"`
        {{- end }}
    {{- end }}
}
{{ end }}