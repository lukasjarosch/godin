package bundle

import (
	"fmt"
	"github.com/iancoleman/strcase"
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

	subscriberKey := strings.Replace(sub.Topic, ".", "_", -1)
	subscriberKey = strings.ToLower(subscriberKey)

	// defaults
	sub.AutoAck = false
	sub.Queue.Durable = true
	sub.Queue.NoWait = false
	sub.Queue.Exclusive = false
	sub.Queue.AutoDelete = false

	confSub := config.GetStringMap(SubscriberKey)
	if _, ok := confSub[subscriberKey]; ok == true {
		return nil, fmt.Errorf("subscriber '%s' is already registered", subscriberKey)
	}
	confSub[subscriberKey] = sub
	config.Set(SubscriberKey, confSub)
	godin.SaveConfiguration()

	return &subscriber{
		Subscription: sub,
		HandlerName:  subscriberKey,
	}, nil
}

// SubscriberHandlerName returns the handler name of the subscriber based on the subscription topic
// The topic 'user.created' will be handled by the 'UserCreatedSubscriber'
func SubscriberHandlerName(topic string) string {
	name := strings.ToLower(topic)
	name = strings.Replace(name, ".", "_", -1)
	name += "_subscriber"
	name = strcase.ToCamel(name)

	return name
}

// SubscriberFileName assembles the target .go file into which the handler is going to be generated.
func SubscriberFileName(topic string) string {
	name := strings.ToLower(topic)
	name = strings.Replace(name, ".", "_", -1)
	return fmt.Sprintf("%s.go", name)
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
