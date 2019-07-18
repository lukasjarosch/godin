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

type AmqpEncodeDecode struct {
	BaseGenerator
}

func NewAMQPEncodeDecode(box packr.Box, serviceInterface *types.Interface, ctx template.Context, options ...Option) *AmqpEncodeDecode {
	defaults := &Options{
		Context:    ctx,
		Overwrite:  true,
		IsGoSource: true,
		Template:   "amqp_encode_decode",
		TargetFile: "internal/amqp/encode_decode.go",
	}

	for _, opt := range options {
		opt(defaults)
	}

	return &AmqpEncodeDecode{
		BaseGenerator{
			box:   box,
			iface: serviceInterface,
			opts:  defaults,
		},
	}
}

// Update will call GenerateFull
func (s *AmqpEncodeDecode) Update() error {
	if s.TargetExists() {
		return s.GenerateMissing()
	}
	return s.GenerateFull()
}

func (s *AmqpEncodeDecode) GenerateMissing() error {
	var publishers []string
	var subscribers []string

	for _, pub := range s.opts.Context.Service.Publisher {
		publishers = append(publishers, pub.Name)
	}

	for _, sub := range s.opts.Context.Service.Subscriber {
		subscribers = append(subscribers, sub.Handler)
	}

	implementation := parse.NewAmqpEncodeDecodeParser(s.opts.TargetFile, publishers, subscribers)
	if err := implementation.Parse(); err != nil {
		return errors.Wrap(err, "Parse")
	}

	if len(implementation.MissingFunctions) > 0 {
		for _, missingFunction := range implementation.MissingFunctions {
			templateName, err := s.templateFromFunction(missingFunction)
			if err != nil {
				return errors.Wrap(err, "unable to find template")
			}

			// extract the correct publisher / subscriber
			// encoders are used by publishers, decoders are used by subscribers
			// e.g. UserCreatedEncoder => UserCreated; it's the required config key and thus the template context to use
			if strings.Contains(missingFunction, "Decoder") {
				// TODO: subscriber
			} else if strings.Contains(missingFunction, "Encoder") {
				var ctx template.Publisher

				name := strings.Replace(missingFunction, "Encoder", "", 1)

				for _, registeredPublisher := range s.opts.Context.Service.Publisher {
					if registeredPublisher.Name == name {
						ctx = registeredPublisher
						break
					}
				}

				tpl := template.NewPartial(templateName, true)
				data, err := tpl.Render(s.box, ctx)
				if err != nil {
					return errors.Wrap(err, "failed to render partial")
				}

				writer := template.NewFileAppendWriter(s.opts.TargetFile, data)
				if err := writer.Write(); err != nil {
					return errors.Wrap(err, fmt.Sprintf("failed to append-write to %s", s.TargetPath()))
				}

				logrus.Infof("added missing publish encoder to %s: %s", s.opts.TargetFile, missingFunction)

			} else {
				logrus.Debugf("ignoring function %s", missingFunction)
				continue
			}
		}
	}
	return nil
}

// templateFromFunction extracts the partial templateName of a encode/decode function
func (r *AmqpEncodeDecode) templateFromFunction(name string) (templateName string, err error) {
	if strings.Contains(name, "Encoder") {
		return "amqp_publish_encode", nil
	}

	return "", fmt.Errorf("function %s does not have a template associated", name)
}
