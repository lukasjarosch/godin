package template

import (
    "text/template"
	"bytes"
	"strings"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/Masterminds/sprig"
)

type Template struct {
    *template.Template
}

func getFuncs() template.FuncMap {
	return template.FuncMap{
		"getProtoName": func(proto descriptor.MethodDescriptorProto) string {
			return proto.GetName()
		},
		"getProtoNameToLower": func(proto descriptor.MethodDescriptorProto) string {
			return strings.ToLower(proto.GetName())
		},
		"getProtoInputType": func(proto descriptor.MethodDescriptorProto) string {
			return strings.Split(proto.GetInputType(), ".")[2]
		},
		"getProtoOutputType": func(proto descriptor.MethodDescriptorProto) string {
			return strings.Split(proto.GetOutputType(), ".")[2]
		},
	}
}

func NewTemplate(buf []byte) (*Template, error) {
	t := template.New("godin")
	t.Funcs(getFuncs())
	t.Funcs(sprig.TxtFuncMap())
	_, err := t.Parse(string(buf))
	if err != nil {
	    return &Template{}, nil
	}

    return &Template{t}, nil
}


func (t *Template) Render(data interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := t.Execute(buf, data)
	if err != nil {
	    return nil, err
	}
	return buf, nil
}
