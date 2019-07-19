package bundle

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/go-godin/rabbitmq"
	"github.com/lukasjarosch/godin/internal/godin"
	"github.com/lukasjarosch/godin/internal/prompt"
	config "github.com/spf13/viper"
)

const SubscriberKey = "transport.amqp.subscriber"

type SubscriberConfiguration struct {
	RabbitMQ    rabbitmq.Subscription
	HandlerName string `json:"handler_name" mapstructure:"handler_name"`
	Protobuf    struct {
		GoModule    string `json:"go_module" mapstructure:"go_module"`
		MessageName string `json:"message_name" mapstructure:"message_name"`
	}
}

type subscriber struct {
	Configuration SubscriberConfiguration
}

func InitializeSubscriber() (*subscriber, error) {
	cfg := SubscriberConfiguration{}
	err := promptSubscriberValues(&cfg)
	if err != nil {
		return nil, err
	}

	subscriberKey := strings.Replace(cfg.RabbitMQ.Topic, ".", "_", -1)
	subscriberKey = strings.ToLower(subscriberKey)

	cfg.HandlerName = strcase.ToCamel(subscriberKey)
	cfg.RabbitMQ.AutoAck = false
	cfg.RabbitMQ.Queue.Durable = true
	cfg.RabbitMQ.Queue.NoWait = false
	cfg.RabbitMQ.Queue.Exclusive = false
	cfg.RabbitMQ.Queue.AutoDelete = false

	confSub := config.GetStringMap(SubscriberKey)
	if _, ok := confSub[subscriberKey]; ok == true {
		return nil, fmt.Errorf("subscriber '%s' is already registered", subscriberKey)
	}
	confSub[subscriberKey] = cfg
	config.Set(SubscriberKey, confSub)
	godin.SaveConfiguration()

	return &subscriber{
		Configuration: cfg,
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

func promptSubscriberValues(cfg *SubscriberConfiguration) (err error) {
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
	cfg.RabbitMQ.Topic = topic

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
	cfg.RabbitMQ.Exchange = exchange

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
	cfg.RabbitMQ.Queue.Name = queue

	// Prefetch count
	p = prompt.NewPrompt(
		"Prefetch count",
		"10",
		prompt.Validate(prompt.PositiveInteger()),
	)
	prefetchCount, err := p.Run()
	if err != nil {
		return err
	}
	count, _ := strconv.Atoi(prefetchCount)
	cfg.RabbitMQ.PrefetchCount = count

	// Protobuf module
	p = prompt.NewPrompt(
		"Import of the target protobuf",
		"github.com/user/some-other-service-protobuf",
		prompt.Validate(),
	)
	protoModule, err := p.Run()
	if err != nil {
		return err
	}
	cfg.Protobuf.GoModule = protoModule

	// Protobuf message
	p = prompt.NewPrompt(
		"Name of the protobuf message of the subscribed event",
		"AnotherServiceMessage",
		prompt.Validate(),
	)
	protoMessage, err := p.Run()
	if err != nil {
		return err
	}
	cfg.Protobuf.MessageName = protoMessage

	return nil
}
