package service

import (
    "errors"
)

// TODO: {{ .ServiceName }} documentation
type {{ .ServiceName }} interface {
}

// Application errors
var (
    ErrNotImplemented = errors.New("endpoint not implemented")
)
