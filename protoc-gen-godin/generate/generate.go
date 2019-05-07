package generate

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"go/format"

	"github.com/golang/protobuf/proto"
	protodescriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	gen "github.com/golang/protobuf/protoc-gen-go/generator"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	tpl "github.com/lukasjarosch/godin/protoc-gen-godin/template"
	"path"
)

type Maker interface {
	Make(*protodescriptor.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error)
}

type Generator struct {
	*gen.Generator
	reader io.Reader
	writer io.Writer
}

func NewGenerator() *Generator {
	return &Generator{
		Generator: gen.New(),
		reader:    os.Stdin,
		writer:    os.Stdout,
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
		wd, _ := os.Getwd()
		// TODO: asset needs to be zipped into the binary
		buf, err := ioutil.ReadFile(path.Join(wd, "template", "godin.pb.go.tmpl"))
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		template, err := tpl.NewTemplate(buf)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		output, err := template.Render(data)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		g.P(output.String())

		formatted, err := format.Source(g.Bytes())
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		file := &plugin.CodeGeneratorResponse_File{
			Name:    proto.String("godin.pb.go"),
			Content: proto.String(string(formatted)),
		}
		response.File = append(response.File, file)
	}
	return response, nil
}
