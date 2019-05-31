{{ define "logging_method" }}
func (l logMiddleware) {{ .Name }}({{ .ParamList }}) ({{ .ReturnList }}) {

    // log the request of {{ .Name }}
	l.logger.Log(
	    "endpoint", "{{ .Name }}",
	    "request", endpoint.{{ .RequestName }}{
	    {{- range .Params }}
	    {{- .Name }},
	    {{- end }}
	    },
	    )

	// log the response of {{ .Name }}
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