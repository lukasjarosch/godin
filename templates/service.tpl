package {{ .ServiceName }}

import (
    "context"
	"errors"

    {{ range .Spec.ResolvedDependencies }}
    "{{ .Import }}"
    {{- end }}
)

{{ $receiver := .ServiceName | camelcase -}}

// {{ .Spec.Service.Description }}
type {{ .ServiceName | camelcase }} struct {
    {{- range .Spec.ResolvedDependencies }}
    {{ .Name }} {{ .Type }}
    {{- end }}
}

var (
    {{ range .Spec.Service.Errors }}
    {{ .Name }} = errors.New("{{ .Message }}")
    {{- end }}
)

// New{{ .GrpcServiceName }} returns the business implementation of {{ .Spec.Service.API.Package }}.{{ .Spec.Service.API.Service }}
func New{{ .GrpcServiceName }}({{ deps_param_list }}) *{{ .ServiceName | camelcase }}{

	service := &{{ .ServiceName | camelcase }}{
	    {{ deps_value_mapping }}
	}

	return service
}

{{ range .Spec.Service.Methods -}}
{{- range .Comments }}
// {{ . }}
{{- end }}
func (svc *{{ $receiver }}) {{ .Name }}({{ arg_list .Name }}) ({{ ret_list .Name }}) {
    return {{ default_value_list .Returns }}
}
{{- end }}

