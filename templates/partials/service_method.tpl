{{ define "service_method" }}
{{- range .Comments }}
{{ . }}
{{- end }}
func (uc *UseCase) {{ .Name }}( {{ .ParamList }}) ({{ .ReturnList }}) {
    return {{ .ReturnImplementationMissing }}
}
{{ end }}