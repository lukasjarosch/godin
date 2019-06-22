package middleware

import (
    "context"
    "time"

    "github.com/lukasjarosch/godin/pkg/log"

    "{{ .Service.Module }}/internal/service"
    "{{ .Service.Module }}/internal/service/endpoint"
)

type loggingMiddleware struct {
    next service.{{ title .Service.Name }}
    logger log.Logger
}

func LoggingMiddleware(logger log.Logger) Middleware {
    return func(next service.{{ title .Service.Name }}) service.{{ title .Service.Name }} {
        return &loggingMiddleware{next, logger}
    }
}

{{ range .Service.Methods }}
{{ template "logging_method" . }}
{{ end }}
