package template

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Data struct {
	ProtoFileName string
	Package string
	ImportPath string
	Services []Service
}

type Service struct {
	Name    string
	Methods []*descriptor.MethodDescriptorProto
}
