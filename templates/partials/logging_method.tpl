{{ define "logging_method" }}
// {{ .Name }} logs the request and response of the service.{{ .Name }} endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l loggingMiddleware) {{ .Name }}({{ .ParamList }}) ({{ .ReturnList }}) {
	l.logger.Log(
	    "endpoint", "{{ .Name }}",
	    "request", endpoint.{{ .RequestName }}{
	    {{- range .Params }}
	        {{- if ne .Name "ctx" }} {{ title .Name }}: {{ .Name }}, {{end}}
        {{ end }}
	    },
	    )

	defer func(begin time.Time) {
	    resp := endpoint.{{ .ResponseName }}{
            {{- range .Returns }}
                {{- if ne .Type "error" }}
                    {{- title .Name }}: {{ .Name }},
                {{- end }}
            {{- end }}
	    }

        l.logger.Log(
            "endpoint", "{{ .Name }}",
            "response", resp,
            "error", err,
            "success", err == nil,
            "took", time.Since(begin),
        )
	}(time.Now())

	return l.next.{{ .Name }}({{ range .Params }}{{ .Name }}, {{ end }})
}
{{ end }}