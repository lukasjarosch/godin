// Code generated by Godin v{{ .Godin.Version }}; DO NOT EDIT.

package grpc

import (
    "context"
    "errors"

    "{{ .Service.Module }}/internal/service/endpoint"
    pb "{{ .Protobuf.Package }}"
)

{{ range .Service.Methods }}
{{ template "grpc_encode_request" . }}
{{ template "grpc_decode_response" . }}
{{ template "grpc_encode_response" . }}
{{ template "grpc_decode_request" . }}
{{ end }}