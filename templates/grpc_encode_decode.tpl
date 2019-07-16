package grpc

import (
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    "{{ .Service.Module }}/internal/service/endpoint"
	"{{ .Service.Module }}/internal/service/domain"
    pb "{{ .Protobuf.Package }}"
)

// ----------------[ ERRORS ]----------------

// EncodeError encodes domain-level errors into gRPC transport-level errors
func EncodeError(err error) error {
    switch err {
	case domain.ErrNotImplemented:
		return status.Error(codes.Unimplemented, err.Error())
    default:
        return status.Error(codes.Unknown, err.Error())
    }
    return err
}

// ----------------[ MAPPING FUNCS ]----------------

// TODO: this is a nice spot for convenience mapping functions :)

// ----------------[ ENCODER / DECODER ]----------------
{{- range .Service.Methods }}
{{- template "grpc_request_decoder" . }}
{{- template "grpc_response_encoder" . }}
{{- template "grpc_request_encoder" . }}
{{- template "grpc_response_decoder" . }}
{{- end }}
