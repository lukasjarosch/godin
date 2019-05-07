{{- range .Comments }}
// {{ . }}
{{- end }}
func ({{ .Receiver }}) {{ .Name }}({{ .ArgList }}) ({{ .ReturnList }}) {
    // TODO: Build something awesome...
    return {{ .DefaultReturn }}
}