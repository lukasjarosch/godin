{{- define "endpoint" }}
// {{ .Name }}Endpoint provides service.{{ .Name }}() as general endpoint
// which can be used with arbitrary transport layers.
func {{ .Name }}Endpoint(service service.{{ title .ServiceName }}) endpoint.Endpoint {
    return func (ctx context.Context, request interface{}) (interface{}, error) {
        req := request.({{ .RequestName }})
        _ = req // bypass "declared and not used" compiler error if the request is empty and not used
        {{ .ReturnVariableList }} := service.{{ .Name }}(
            {{- range .Params }}
                {{- if eq .Name "ctx" }}
                    {{- .Name }},
                {{- else -}}
                    req.{{ title .Name }},
                {{- end }}
            {{- end }})

        return {{ .ResponseName }}{
            {{- range .Returns }}
            {{ title .Name }}: {{ .Name }},
            {{- end }}
        }, err
    }
}
{{- end }}
