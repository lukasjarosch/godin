package {{ .Service.Name }}

import (
    "context"
    "github.com/go-kit/kit/log"
    "{{ .Service.Module }}/internal/service"
)

type serviceImplementation struct {
    logger log.Logger
}

func NewServiceImplementation(logger log.Logger) *serviceImplementation {
    return &serviceImplementation{
        logger: logger,
    }
}

{{ range .Service.Methods }}
 {{ template "service_method" . }}
{{ end }}


