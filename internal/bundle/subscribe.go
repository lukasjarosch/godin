package bundle

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/lukasjarosch/godin/internal/godin"
	"github.com/lukasjarosch/godin/internal/prompt"
	"github.com/lukasjarosch/godin/pkg/amqp"
	config "github.com/spf13/viper"
)

const SubscriberKey = "transport.amqp.subscriber"

type subscriber struct {
	Subscription amqp.Subscription
	HandlerName  string `json:"handler_name"`
}

func InitializeSubscriber() (*subscriber, error) {
	sub := amqp.Subscription{}
	handlerName, err := promptValues(&sub)
	if err != nil {
		return nil, err
	}

	handlerName = strcase.ToSnake(handlerName)
	sub.ExchangeType = "durable" // currently not configurable

	confSub := config.GetStringMap(SubscriberKey)
	if _, ok := confSub[handlerName]; ok == true {
		return nil, fmt.Errorf("subscriber '%s' is already registered", handlerName)
	}
	confSub[handlerName] = sub
	config.Set(SubscriberKey, confSub)
	godin.SaveConfiguration()

	return &subscriber{
		Subscription: sub,
		HandlerName:  handlerName,
	}, nil
}

func promptValues(sub *amqp.Subscription) (handlerName string, err error) {
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
		"user-exchange",
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
		"user-created-queue",
		prompt.Validate(),
	)
	queue, err := p.Run()
	if err != nil {
		return "", err
	}
	sub.Queue = queue

	// HandlerName
	p = prompt.NewPrompt(
		"Name of the handler for this subscription (CamelCase)",
		"UserCreatedSubscriber",
		prompt.Validate(
			prompt.CamelCase(),
		),
	)
	handler, err := p.Run()
	if err != nil {
		return "", err
	}
	return handler, nil
}
