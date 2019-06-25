package bundle

import (
	"github.com/lukasjarosch/godin/internal/prompt"
	"github.com/lukasjarosch/godin/pkg/amqp"
)

type subscriber struct {
	Subscription amqp.Subscription
	HandlerName  string `json:"handler_name"`
}

func InitializeSubscriber() (*subscriber, error) {
	sub := amqp.Subscription{}
	handlerName, err := promptValues(sub)
	if err != nil {
		return nil, err
	}

	return &subscriber{
		Subscription: sub,
		HandlerName:  handlerName,
	}, nil
}

func promptValues(sub amqp.Subscription) (handlerName string, err error) {
	// Topic
	p := prompt.NewPrompt(
		"AMQP subscription Topic",
		"user.created",
		prompt.Validate(),
	)
	topic, err := p.Run()
	if err != nil {
		return "", err
	}
	sub.Topic = topic

	// Exchange
	p = prompt.NewPrompt(
		"AMQP Exchange name",
		"user-Exchange",
		prompt.Validate(),
	)
	exchange, err := p.Run()
	if err != nil {
		return "", err
	}
	sub.Exchange = exchange

	// Queue
	p = prompt.NewPrompt(
		"AMQP Queue name which is bound to the Exchange",
		"user-created-Queue",
		prompt.Validate(),
	)
	queue, err := p.Run()
	if err != nil {
		return "", err
	}
	sub.Queue = queue

	// HandlerName
	p = prompt.NewPrompt(
		"Name of the handler for this subscription",
		"UserCreatedHandler",
		prompt.Validate(),
	)
	handler, err := p.Run()
	if err != nil {
		return "", err
	}
	return handler, nil
}
