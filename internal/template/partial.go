package template

import (
	"bytes"
	"text/template"

	"fmt"

	"github.com/Masterminds/sprig"
	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
)

var PartialTemplates = map[string]string{
	"service_method":       "partials/service_method.tpl",
	"logging_method":       "partials/logging_method.tpl",
	"request":              "partials/request.tpl",
	"response":             "partials/response.tpl",
	"grpc_encode_request":  "partials/grpc/request_response/encode_request.tpl",
	"grpc_encode_response": "partials/grpc/request_response/encode_response.tpl",
	"grpc_decode_request":  "partials/grpc/request_response/decode_request.tpl",
	"grpc_decode_response": "partials/grpc/request_response/decode_response.tpl",
	"grpc_request_decoder": "partials/grpc/encode_decode/request_decoder.tpl",
	"grpc_request_encoder": "partials/grpc/encode_decode/request_encoder.tpl",
	"grpc_response_decoder": "partials/grpc/encode_decode/response_decoder.tpl",
	"grpc_response_encoder": "partials/grpc/encode_decode/response_encoder.tpl",
}

type Partial struct {
	BaseTemplate
	tpl       *template.Template
	templates map[string]string
}

func NewPartial(name string, isGoSource bool) *Partial {

	return &Partial{
		BaseTemplate: BaseTemplate{
			name:       name,
			isGoSource: isGoSource,
		},
	}
}

func (p *Partial) Render(fs packr.Box, templateContext interface{}) (rendered []byte, err error) {

	templatePath, ok := p.templates[p.name]
	if !ok {
		return nil, fmt.Errorf("unknown partial template: %s", p.name)
	}

	templateData, err := fs.FindString(templatePath)
	if err != nil {
		return nil, errors.Wrap(err, "FindString")
	}

	// in order to use the partial template definitions, we need to include them somewhere
	// thus we just add the template to use itself
	templateData += fmt.Sprintf("{{ template \"%s\" . }}", p.name)

	p.tpl, err = template.New(p.Filename()).Funcs(sprig.TxtFuncMap()).Parse(templateData)
	if err != nil {
		return nil, errors.Wrap(err, "Parse")
	}

	buf := bytes.Buffer{}
	if err := p.tpl.Execute(&buf, templateContext); err != nil {
		return nil, err
	}

	if p.isGoSource {
		return p.FormatCode(buf.Bytes())
	}

	return buf.Bytes(), nil
}
