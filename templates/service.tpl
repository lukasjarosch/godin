package {{ .ServiceName }}

import (
    "context"
	"errors"

	"{{ .ModuleName }}/internal/config"
	"github.com/sirupsen/logrus"
)

{{ $receiver := .ServiceName -}}

// {{ .ServiceName }}API is the actual business-logic which you want to provide
type {{ .ServiceName }}API struct {
    {{- range .Spec.Service.Dependencies }}
    {{ .Name }} {{ .Type }}
    {{- end }}
}

var (
    {{ range .Spec.Service.Errors }}
    {{ .Name }} = errors.New("{{ .Message }}")
    {{- end }}
)

// NewExampleAPI returns our business-implementation of the ExampleAPI
func New{{ .GrpcServiceName }}({{ deps_param_list }}) *{{ .ServiceName }}API{

	service := &{{ .ServiceName }}API{
		logger: logger,
		config: config,
	}

	return service
}

{{ range .Spec.Service.Methods -}}
// Greeting implements the business-logic for this RPC
func (svc *{{ $receiver }}API) {{ .Name }}({{ arg_list .Name }}) ({{ ret_list .Name }}) {
    return {{ default_value_list .Returns }}
}
{{- end }}

