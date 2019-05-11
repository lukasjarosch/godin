package generate

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"fmt"

	"github.com/gobuffalo/packr"
	"github.com/golang/protobuf/proto"
	protodescriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	gen "github.com/golang/protobuf/protoc-gen-go/generator"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	tpl "github.com/lukasjarosch/godin/protoc-gen-godin/template"
	"github.com/pkg/errors"
	"path"
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
	for _, file := range request.ProtoFile {
		if len(file.GetService()) <= 0 {
			continue
		}

		if file.Options == nil || file.Options.GoPackage == nil {
			log.Fatalf("No go_package option defined in %s", *file.Name)
		}

		i, p, _ := goPackageOption(*file)

		data := tpl.Data{
			Package:  p,
			ProtoFileName: *file.Name,
			ImportPath: i,
			Services: make([]tpl.Service, len(file.Service)),
		}

		for _, service := range file.Service {
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

		filename := file.GetName()[0 : len(".proto")+1]
		file := &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(path.Join(data.ImportPath, fmt.Sprintf("%s.godin.go", filename))),
			Content: proto.String(string(output.Bytes())),
		}
		response.File = append(response.File, file)
	}
	return response, nil
}

// SNIPPETS BELOW ARE SHAMELESSLY STOLEN FROM: https://github.com/micro/protoc-gen-micro

func cleanPackageName(name string) string {
	name = strings.Map(badToUnderscore, name)
	// Identifier must not be keyword or predeclared identifier: insert _.
	if isGoKeyword[name] {
		name = "_" + name
	}
	// Identifier must not begin with digit: insert _.
	if r, _ := utf8.DecodeRuneInString(name); unicode.IsDigit(r) {
		name = "_" + name
	}
	return string(name)
}

// badToUnderscore is the mapping function used to generate Go names from package names,
// which can be dotted in the input .proto file.  It replaces non-identifier characters such as
// dot or dash with underscore.
func badToUnderscore(r rune) rune {
	if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
		return r
	}
	return '_'
}

func goPackageOption(file protodescriptor.FileDescriptorProto) (impPath string, pkg string, ok bool) {
	opt := file.GetOptions().GetGoPackage()
	if opt == "" {
		return "", "", false
	}
	// A semicolon-delimited suffix delimits the import path and package name.
	sc := strings.Index(opt, ";")
	if sc >= 0 {
		return string(opt[:sc]), cleanPackageName(opt[sc+1:]), true
	}
	// The presence of a slash implies there's an import path.
	slash := strings.LastIndex(opt, "/")
	if slash >= 0 {
		return string(opt), cleanPackageName(opt[slash+1:]), true
	}
	return "", cleanPackageName(opt), true
}

var isGoKeyword = map[string]bool{
	"break":       true,
	"case":        true,
	"chan":        true,
	"const":       true,
	"continue":    true,
	"default":     true,
	"else":        true,
	"defer":       true,
	"fallthrough": true,
	"for":         true,
	"func":        true,
	"go":          true,
	"goto":        true,
	"if":          true,
	"import":      true,
	"interface":   true,
	"map":         true,
	"package":     true,
	"range":       true,
	"return":      true,
	"select":      true,
	"struct":      true,
	"switch":      true,
	"type":        true,
	"var":         true,
}
