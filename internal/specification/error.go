package specification

import (
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

var (
	ErrorNameEmpty    = errors.New("missing error name")
	ErrorMessageEmpty = errors.New("missing error message")
	ErrorCodeEmpty    = errors.New("missing error code")
	ErrorInvalidCode  = errors.New("invalid error code")
)

type ErrorSpec struct {
	Name    string
	Message string
	Code    string
}

func (e ErrorSpec) Validate() error {
	if e.Name == "" {
		return ErrorNameEmpty
	}
	if e.Message == "" {
		return ErrorMessageEmpty
	}
	if e.Code == "" {
		return ErrorCodeEmpty
	}

	// If grpc code is valid, set the Code param
	_, err := e.resolveCode()
	if err != nil {
		return err
	}

	return nil
}

func (e ErrorSpec) CodeString() string {
	c, _ := e.resolveCode()
	return fmt.Sprintf("codes.%s", c.String())
}

func (e ErrorSpec) resolveCode() (*codes.Code, error) {
	if c, ok := strToCode[e.Code]; ok {
		return &c, nil
	}
	return nil, errors.New(fmt.Sprintf("invalid error code: %s", e.Code))
}

// from: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
var strToCode = map[string]codes.Code{
	"OK": codes.OK,
	"Cancelled":/* [sic] */ codes.Canceled,
	"Unknown":            codes.Unknown,
	"InvalidArgument":    codes.InvalidArgument,
	"DeadlineExceeded":   codes.DeadlineExceeded,
	"NotFound":           codes.NotFound,
	"AlreadyExists":      codes.AlreadyExists,
	"PermissionDenied":   codes.PermissionDenied,
	"ResourceExhausted":  codes.ResourceExhausted,
	"FailedPrecondition": codes.FailedPrecondition,
	"Aborted":            codes.Aborted,
	"OutOfRange":         codes.OutOfRange,
	"Unimplemented":      codes.Unimplemented,
	"Internal":           codes.Internal,
	"Unavailable":        codes.Unavailable,
	"DataLoss":           codes.DataLoss,
	"Unauthenticated":    codes.Unauthenticated,
}
