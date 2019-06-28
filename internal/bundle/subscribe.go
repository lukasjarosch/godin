package bundle

import (
	"fmt"
	"strings"

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
	err := promptValues(&sub)
	if err != nil {
		return nil, err
	}

	handlerName := strings.Replace(sub.Topic, ".", "_", -1)
	handlerName = strings.ToLower(handlerName)

	// defaults
	sub.AutoAck = false
	sub.Queue.Durable = true
	sub.Queue.NoWait = false
	sub.Queue.Exclusive = false
	sub.Queue.AutoDelete = false

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

func promptValues(sub *amqp.Subscription) (err error) {
	// Topic
	p := prompt.NewPrompt(
		"AMQP subscription Topic",
		"user.created",
		prompt.Validate(),
	)
	topic, err := p.Run()
	if err != nil {
		return err
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
		return err
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
		return err
	}
	sub.Queue.Name = queue

	return  nil
}
