package generate

import (
	"fmt"
	"regexp"

	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/parse"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/vetcher/go-astra/types"
)

type GrpcRequestResponse struct {
	BaseGenerator
}

const (
	EncodeRequestPartial  = "grpc_encode_request"
	EncodeResponsePartial = "grpc_encode_response"
	DecodeRequestPartial  = "grpc_decode_request"
	DecodeResponsePartial = "grpc_decode_response"
)

func NewGrpcRequestResponse(box packr.Box, serviceInterface *types.Interface, ctx template.Context, options ...Option) *GrpcRequestResponse {
	defaults := &Options{
		Context:    ctx,
		Overwrite:  true,
		IsGoSource: true,
		Template:   "grpc_request_response",
		TargetFile: "internal/grpc/request_response.go",
	}

	for _, opt := range options {
		opt(defaults)
	}

	return &GrpcRequestResponse{
		BaseGenerator{
			box:   box,
			iface: serviceInterface,
			opts:  defaults,
		},
	}
}

// Update is disabled for this file, it will only proxy the call to GenerateFull()
func (r *GrpcRequestResponse) Update() error {
	if r.TargetExists() {
		return r.GenerateMissing()
	}
	return r.GenerateFull()
}

// GenerateMissing will generate all missing encode/decode functions and append them to the existing file
func (r *GrpcRequestResponse) GenerateMissing() error {
	implementation := parse.NewTransportRequestResponseParser(r.opts.TargetFile, r.iface)
	if err := implementation.Parse(); err != nil {
		return errors.Wrap(err, "RequestResponse.Parse")
	}

	if len(implementation.MissingFunctions) > 0 {
		for _, missingMethod := range implementation.MissingFunctions {
			templateName, err := r.templateFromFunction(missingMethod)
			if err != nil {
				return errors.Wrap(err, "unable to find template")
			}

			ctx := struct {
				Name string
			}{
				implementation.EndpointName(missingMethod),
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
			logrus.Infof("added missing function to %s: %s", r.opts.TargetFile, missingMethod)
		}
	}

	return nil
}

// templateFromFunction extracts the partial templateName of a encode/decode function
func (r *GrpcRequestResponse) templateFromFunction(name string) (templateName string, err error) {
	encodeRequest := regexp.MustCompile(`(Encode)\w+(Request)`)
	encodeResponse := regexp.MustCompile(`(Encode)\w+(Response)`)
	decodeRequest := regexp.MustCompile(`(Decode)\w+(Request)`)
	decodeResponse := regexp.MustCompile(`(Decode)\w+(Response)`)

	if encodeRequest.MatchString(name) {
		return EncodeRequestPartial, nil
	}
	if encodeResponse.MatchString(name) {
		return EncodeResponsePartial, nil
	}
	if decodeRequest.MatchString(name) {
		return DecodeRequestPartial, nil
	}
	if decodeResponse.MatchString(name) {
		return DecodeResponsePartial, nil
	}

	return "", fmt.Errorf("function %s does not have a template associated", name)
}
