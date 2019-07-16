// Code generated by Godin v{{ .Godin.Version }}; DO NOT EDIT.

package middleware

import (
    "context"
    "time"

    "github.com/go-kit/kit/log"
	grpc_metadata "github.com/go-godin/grpc-metadata"

    service "{{ .Service.ImportPath }}"
)

type loggingMiddleware struct {
    logger log.Logger
    next service.{{ .Service.Name }}
}

func NewLoggingMiddleware(logger log.Logger) Middleware {
    return func(next service.{{ .Service.Name }}) service.{{ .Service.Name }} {
        return &loggingMiddleware{logger, next}
    }
}

{{ range .Service.Methods }}
func (l loggingMiddleware) {{ .Name }}({{ .ParamList }}) ({{ .ReturnList }}) {
}
{{- end }}