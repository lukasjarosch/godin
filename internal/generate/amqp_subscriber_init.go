package generate

import (
	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/vetcher/go-astra/types"
)

type AmqpSubscriber struct {
	BaseGenerator
}

func NewAMQPSubscriber(box packr.Box, serviceInterface *types.Interface, ctx template.Context, options ...Option) *AmqpSubscriber {
	defaults := &Options{
		Context:    ctx,
		Overwrite:  true,
		IsGoSource: true,
		Template:   "amqp_subscriber_init",
		TargetFile: "internal/amqp/subscriptions.go",
	}

	for _, opt := range options {
		opt(defaults)
	}

	return &AmqpSubscriber{
		BaseGenerator{
			box:   box,
			iface: serviceInterface,
			opts:  defaults,
		},
	}
}

// Update will call GenerateFull
func (s *AmqpSubscriber) Update() error {
	return s.GenerateFull()
}
