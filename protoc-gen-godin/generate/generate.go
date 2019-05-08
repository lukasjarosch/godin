package generate

import (
	"io"
	"io/ioutil"
	"os"

	"fmt"

	"github.com/gobuffalo/packr"
	"github.com/golang/protobuf/proto"
	protodescriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	gen "github.com/golang/protobuf/protoc-gen-go/generator"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	tpl "github.com/lukasjarosch/godin/protoc-gen-godin/template"
	"github.com/pkg/errors"
)

type Maker interface {
	Make(*protodescriptor.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error)
}

type Generator struct {
	*gen.Generator
	box    packr.Box
	reader io.Reader
	writer io.Writer
}

func NewGenerator(box packr.Box) *Generator {
	return &Generator{
		Generator: gen.New(),
		reader:    os.Stdin,
		writer:    os.Stdout,
		box:       box,
	}
}

func (g *Generator) Generate() {
	input, err := ioutil.ReadAll(g.reader)
	if err != nil {
		g.Error(err, "reading input")
	}

	request := g.Request
	if err := proto.Unmarshal(input, request); err != nil {
		g.Error(err, "unmarshalling input data")
	}

	if len(request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}

	g.CommandLineParameters(g.Request.GetParameter())
	g.WrapTypes()
	g.SetPackageNames()
	g.BuildTypeNameMap()
	g.GenerateAllFiles()
	g.Reset()

	response, err := g.generate(request)
	if err != nil {
		g.Error(err, "failed to generate response")
	}

	output, err := proto.Marshal(response)
	if err != nil {
		g.Error(err, "failed to marshal output proto")
	}
	_, err = g.writer.Write(output)
	if err != nil {
		g.Error(err, "failed to write output proto")
	}
}

func (g *Generator) generate(request *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	response := new(plugin.CodeGeneratorResponse)
	for _, protoFile := range request.ProtoFile {
		if len(protoFile.GetService()) <= 0 {
			continue
		}

		data := tpl.Data{
			Package:  protoFile.GetPackage(),
			Services: make([]tpl.Service, len(protoFile.Service)),
		}

		for _, service := range protoFile.Service {
			data.Services = append(data.Services, tpl.Service{
				Name:    gen.CamelCase(service.GetName()),
				Methods: service.Method,
			})
		}
		s, err := g.box.Find("./godin.pb.go.tmpl")
		if err != nil {
			return nil, errors.Wrap(err, "could not load asset")
		}

		template, err := tpl.NewTemplate(s)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse template")
		}

		output, err := template.Render(data)
		if err != nil {
			return nil, errors.Wrap(err, "failed to render template")
		}

		g.P(output.String())

		/*
			formatted, err := format.Source(g.Bytes())
			if err != nil {
				return nil, errors.Wrap(err, "failed to format source")
			}
		*/

		filename := protoFile.GetName()[0 : len(".proto")+1]
		file := &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(fmt.Sprintf("%s.godin.go", filename)),
			Content: proto.String(string(output.Bytes())),
		}
		response.File = append(response.File, file)
	}
	return response, nil
}
