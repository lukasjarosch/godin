package usecase

import (
    "context"

    "github.com/go-godin/log"
    "{{ .Service.Module }}/internal/service"
	"{{ .Service.Module }}/internal/service/domain"
)

// UseCase implements all business use-cases
type UseCase struct {
    logger log.Logger
}

// NewServiceImplementation constructs the use-case (service) layer of the microservice.
// It provides the whole business use-cases and works on the domain entities.
func NewServiceImplementation(logger log.Logger) *UseCase {
    return &serviceImplementation{
        logger: logger,
    }
}

{{- range .Service.Methods }}
{{- template "service_method" . }}
{{- end }}


