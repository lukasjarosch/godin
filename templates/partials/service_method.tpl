{{ define "service_method" }}
{{- range .Comments }}
{{ . }}
{{- end }}
func (s *serviceImplementation) {{ .Name }}( {{ .ParamList }}) ({{ .ReturnList }}) {
    return {{ .ReturnImplementationMissing }}
}
{{ end }}