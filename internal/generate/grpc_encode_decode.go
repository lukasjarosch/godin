package generate

import (
	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/sirupsen/logrus"
	"github.com/vetcher/go-astra/types"
)

type GrpcEncodeDecode struct {
	BaseGenerator
}

func NewGrpcEncodeDecode(box packr.Box, serviceInterface *types.Interface, ctx template.Context, options ...Option) *RequestResponse {
	defaults := &Options{
		Context:    ctx,
		Overwrite:  true,
		IsGoSource: true,
		Template:   "grpc_encode_decode",
		TargetFile: "internal/grpc/encode_decode.go",
	}

	for _, opt := range options {
		opt(defaults)
	}

	return &RequestResponse{
		BaseGenerator{
			box:   box,
			iface: serviceInterface,
			opts:  defaults,
		},
	}
}

// Update is disabled for this file, it will only proxy the call to GenerateFull()
func (r *GrpcEncodeDecode) Update() error {
	if r.TargetExists() {
		return r.GenerateMissing()
	}
	return r.GenerateFull()
}

func (r *GrpcEncodeDecode) GenerateMissing() error {
	logrus.Warn("not yet implemented")
	return nil
}

