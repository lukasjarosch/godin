package grpc

import (
    "context"
    "errors"

    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    "{{ .Service.Module }}/endpoint"
    service "{{ .Service.Module }}"
)

// EncodeError encodes domain-level errors into gRPC transport-level errors
func EncodeError(err error) error {
    switch err {
    default:
        return status.Error(codes.Unknown, err.Error())
    }
    return err
}
