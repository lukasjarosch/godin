package {{ .ServiceName }}

import (
	"errors"

	"{{ .ModuleName }}/internal/config"
	"github.com/sirupsen/logrus"
)

{{- $receiver := .ServiceName -}}

// {{ .ServiceName }}API is the actual business-logic which you want to provide
type {{ .ServiceName }}API struct {
	config *config.Config
	logger *logrus.Logger
}

var (
	ErrEmptyName = errors.New("the given name is empty")
)

// NewExampleAPI returns our business-implementation of the ExampleAPI
func New{{ .GrpcServiceName }}(config *config.Config, logger *logrus.Logger) *{{ .ServiceName }}API{

	service := &{{ .ServiceName }}API{
		logger: logger,
		config: config,
	}

	return service
}

{{ range .Spec.Service.Methods -}}
// Greeting implements the business-logic for this RPC
func (svc *{{ $receiver }}API) {{ .Name }}({{ arg_list .Name }}) ({{ ret_list .Name }}) {
}
{{- end }}

