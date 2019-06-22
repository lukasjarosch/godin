package middleware

import (
    "context"
    "time"

    "github.com/go-kit/kit/log"

    "{{ .Service.Module }}/internal/service"
    "{{ .Service.Module }}/internal/service/endpoint"
)

type loggingMiddleware struct {
    logger log.Logger
    next service.{{ title .Service.Name }}
}

{{ range .Service.Methods }}
{{ template "logging_method" . }}
{{ end }}
