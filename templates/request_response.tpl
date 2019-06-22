// Code generated by Godin v{{ .Godin.Version }}; DO NOT EDIT.

package endpoint

import (
    "context"
    "github.com/go-kit/kit/endpoint"

    service "{{ .Service.Module }}"
)

type (
{{ range .Service.Methods }}
{{ template "request" . }}
{{ template "response" . }}
{{ end }}
)

// Implement the Failer interface for all responses
{{- range .Service.Methods }}
func (resp {{ .Response }}) Failed() error { return r.Err }
{{- end }}
