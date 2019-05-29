{{ block "endpoint" . }}
// {{ .Name }}Endpoint provides service.{{ .Name }}() as general endpoint
// which can be used with arbitrary transport layers.
func {{ .Name }}Endpoint(service service.{{ .ServiceName }}) endpoint.Endpoint {
    return func (ctx context.Context, request interface{}) (response interface{}, err error) {
        req := request.({{ .RequestName }})
        {{ .ReturnVariableList }} := service.{{ .Name }}(
            {{- range .Params }}
                {{- if eq .Type "context.Context" }}
                    {{- .Name }},
                {{- else -}}
                    req.{{ title .Name }},
                {{- end }}
            {{- end }})
        return {{ .ResponseName }}{
            {{- range .Returns -}}
            {{- .Name }},
            {{- end }}
            }
    }
}
{{ end }}
