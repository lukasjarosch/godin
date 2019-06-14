package grpc

import (
    "context"
    "errors"

    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    "{{ .Service.Module }}/endpoint"
    service "{{ .Service.Module }}"
)

// ----------------[ ERRORS ]----------------

// EncodeError encodes domain-level errors into gRPC transport-level errors
func EncodeError(err error) error {
    switch err {
    default:
        return status.Error(codes.Unknown, err.Error())
    }
    return err
}

// ----------------[ MAPPING FUNCS ]----------------

// TODO: this is a nice spot for convenience mapping functions :)

// ----------------[ SERVER ]----------------

{{ range .Service.Methods }}
{{ template "grpc_request_decoder" . }}
{{ template "grpc_response_encoder" . }}
{{ end }}

// ----------------[ CLIENT ]----------------

{{ range .Service.Methods }}
{{ template "grpc_request_encoder" . }}
{{ template "grpc_response_decoder" . }}
{{ end }}
