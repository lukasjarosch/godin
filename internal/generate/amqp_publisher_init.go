package generate

import (
	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/vetcher/go-astra/types"
)

type AmqpPublisher struct {
	BaseGenerator
}

func NewAMQPPublisher(box packr.Box, serviceInterface *types.Interface, ctx template.Context, options ...Option) *AmqpPublisher {
	defaults := &Options{
		Context:    ctx,
		Overwrite:  true,
		IsGoSource: true,
		Template:   "amqp_publisher_init",
		TargetFile: "internal/amqp/publishers.go",
	}

	for _, opt := range options {
		opt(defaults)
	}

	return &AmqpPublisher{
		BaseGenerator{
			box:   box,
			iface: serviceInterface,
			opts:  defaults,
		},
	}
}

// Update will call GenerateFull
func (s *AmqpPublisher) Update() error {
	return s.GenerateFull()
}
