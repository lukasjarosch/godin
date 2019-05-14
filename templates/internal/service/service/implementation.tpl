package {{ .Service.Name }}

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
)

// Implementation of the service's business-logic of the endpoints
type implementation struct {
    logger log.Logger
}

func NewImplementation(logger log.Logger) *implementation {
    return &implementation{logger:logger}
}

