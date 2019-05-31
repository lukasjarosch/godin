{{ define "logging_method" }}
// {{ .Name }} logs the request and response of the service.{{ .Name }} endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l logMiddleware) {{ .Name }}({{ .ParamList }}) ({{ .ReturnList }}) {
	l.logger.Log(
	    "endpoint", "{{ .Name }}",
	    "request", endpoint.{{ .RequestName }}{
	    {{- range .Params }}
	    {{- .Name }},
	    {{- end }}
	    },
	    )

	defer func(begin time.Time) {
	    resp := endpoint.{{ .ResponseName }}{
            {{- range .Returns }}
                {{- if ne .Type "error" }}
                    {{- .Name }},
                {{- end }}
            {{- end }}
	    }
		if err != nil {
		    resp.Err = err
		}

        i.logger.Log(
            "endpoint", "{{ .Name }}",
            "response", resp,
            "took", time.Since(begin),
        )
	}(time.Now())

	return i.next.{{ .Name }}({{ range .Params }}{{ .Name }}, {{ end }})
}
{{ end }}