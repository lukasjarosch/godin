package usecase

import (
    "context"

    "github.com/go-godin/log"
    _ "{{ .Service.Module }}/internal/service"
	"{{ .Service.Module }}/internal/service/domain"
)

type serviceImplementation struct {
    logger log.Logger
}

func NewServiceImplementation(logger log.Logger) *serviceImplementation {
    return &serviceImplementation{
        logger: logger,
    }
}

{{- range .Service.Methods }}
{{- template "service_method" . }}
{{- end }}


