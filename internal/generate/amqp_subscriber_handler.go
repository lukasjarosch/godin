package generate

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/bundle"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/vetcher/go-astra/types"
)

type AmqpSubscriberHandler struct {
	BaseGenerator
	subscriber template.Subscriber
}

func NewAMQPSubscriberHandler(sub template.Subscriber, box packr.Box, serviceInterface *types.Interface, ctx template.Context, options ...Option) *AmqpSubscriberHandler {
	defaults := &Options{
		Context:    ctx,
		Overwrite:  true,
		IsGoSource: true,
		Template:   "amqp_subscriber_handler",
		TargetFile: fmt.Sprintf("internal/service/subscriber/%s", bundle.SubscriberFileName(sub.Subscription.Topic)),
	}

	for _, opt := range options {
		opt(defaults)
	}

	return &AmqpSubscriberHandler{
		BaseGenerator{
			box:   box,
			iface: serviceInterface,
			opts:  defaults,
		},
		sub,
	}
}

func (s *AmqpSubscriberHandler) Update() error {

	// overwrite the contet subscribers with the - to be generated - subscriber.
	s.opts.Context.Service.Subscriber = []template.Subscriber{s.subscriber}

	if s.TargetExists() {
		return nil
	}
	return s.GenerateFull()
}
