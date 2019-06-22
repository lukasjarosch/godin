package generate

import (
	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/vetcher/go-astra/types"
)

type GrpcServer struct {
	BaseGenerator
}

func NewGrpcServer(box packr.Box, serviceInterface *types.Interface, ctx template.Context, options ...Option) *GrpcServer {
	defaults := &Options{
		Context:    ctx,
		Overwrite:  true,
		IsGoSource: true,
		Template:   "grpc_server",
		TargetFile: "internal/grpc/server.go",
	}

	for _, opt := range options {
		opt(defaults)
	}

	return &GrpcServer{
		BaseGenerator{
			box:   box,
			iface: serviceInterface,
			opts:  defaults,
		},
	}
}

// Update will call GenerateFull. The grpc/server cannot be updated.
func (s *GrpcServer) Update() error {
	return s.GenerateFull()
}

