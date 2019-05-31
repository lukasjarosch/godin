package middleware

import (
    "context"
    "time"

    "github.com/go-kit/kit/log"

    "{{ .Service.Module }}/internal/service"
)

type loggingMiddleware struct {
    logger log.Logger
    next service.{{ title .Service.Name }}
}

{{ range .Service.Methods }}
{{ template "logging_method" . }}
{{ end }}
