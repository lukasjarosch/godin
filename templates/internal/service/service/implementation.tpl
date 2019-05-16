package {{ .Service.Name }}

import (
	"github.com/go-kit/kit/log"
)

type implementation struct {
    logger log.Logger
}

// NewImplementation returns the actual implementation of the {{ title .Service.Name }} service
func NewImplementation(logger log.Logger) *implementation {
    return &implementation{logger:logger}
}

