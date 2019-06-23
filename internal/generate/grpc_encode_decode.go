package generate

import (
	"fmt"

	"strings"

	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/parse"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/vetcher/go-astra/types"
)

type GrpcEncodeDecode struct {
	BaseGenerator
}

const (
	RequestEncoderPartial  = "grpc_request_encoder"
	RequestDecoderPartial  = "grpc_request_decoder"
	ResponseEncoderPartial = "grpc_response_encoder"
	ResponseDecoderPartial = "grpc_response_decoder"
)

func NewGrpcEncodeDecode(box packr.Box, serviceInterface *types.Interface, ctx template.Context, options ...Option) *GrpcEncodeDecode {
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

	return &GrpcEncodeDecode{
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
	implementation := parse.NewTransportEncodeDecodeParser(r.opts.TargetFile, r.iface)
	if err := implementation.Parse(); err != nil {
		return errors.Wrap(err, "Parse")
	}

	if len(implementation.MissingFunctions) > 0 {
		for _, missingFunction := range implementation.MissingFunctions {
			templateName, err := r.templateFromFunction(missingFunction)
			if err != nil {
				return errors.Wrap(err, "unable to find template")
			}

			// extract the required method from the large templateContext
			// we only need the method as context in this case
			var ctx template.Method
			for _, methCtx := range r.opts.Context.Service.Methods {
				if strings.Contains(missingFunction, methCtx.Name) {
					ctx = methCtx
				}
			}

			tpl := template.NewPartial(templateName, true)
			data, err := tpl.Render(r.box, ctx)
			if err != nil {
				return errors.Wrap(err, "failed to render partial")
			}

			writer := template.NewFileAppendWriter(r.opts.TargetFile, data)
			if err := writer.Write(); err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to append-write to %s", r.TargetPath()))
			}
			logrus.Infof("added missing function to %s: %s", r.opts.TargetFile, missingFunction)
		}
	}

	return nil
}

// templateFromFunction extracts the partial templateName of a encode/decode function
func (r *GrpcEncodeDecode) templateFromFunction(name string) (templateName string, err error) {
	if strings.Contains(name, "RequestDecoder") {
		return RequestDecoderPartial, nil
	}
	if strings.Contains(name, "RequestEncoder") {
		return RequestEncoderPartial, nil
	}
	if strings.Contains(name, "ResponseDecoder") {
		return ResponseDecoderPartial, nil
	}
	if strings.Contains(name, "ResponseEncoder") {
		return ResponseEncoderPartial, nil
	}

	return "", fmt.Errorf("function %s does not have a template associated", name)
}
