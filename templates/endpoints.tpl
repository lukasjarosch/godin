package endpoint

import (
    "context"
    "github.com/go-kit/kit/endpoint"

    "{{ .Service.Module }}/internal/service"
)

{{- range .Service.Methods -}}
{{- template "endpoint" . -}}
{{- end -}}