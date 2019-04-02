package template

import (
	"github.com/lukasjarosch/godin/internal/specification"
)

type Data struct {
	ProjectRootPath string
	ServiceName string
	ModuleName string
	GrpcServiceName string
	Spec *specification.Specification
}

