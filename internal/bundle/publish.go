package bundle

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/go-godin/rabbitmq"
	"github.com/lukasjarosch/godin/internal/godin"
	config "github.com/spf13/viper"
	"github.com/streadway/amqp"

	"github.com/lukasjarosch/godin/internal/prompt"
)

const PublisherKey = "transport.amqp.publisher"

type publisher struct {
	Publishing rabbitmq.Publishing
}


func InitializePublisher() (*publisher, error) {
	pub := rabbitmq.Publishing{}
	protoMessage, err := promptPublisherValues(&pub)
	if err != nil {
		return nil, err
	}

	pub.DeliveryMode = amqp.Persistent
	pub.ProtobufMessage = protoMessage

	// unique publisher key
	publisherKey := strings.Replace(pub.Topic, ".", "_", -1)
	publisherKey = strings.ToLower(publisherKey)

	// configure
	confPub := config.GetStringMap(PublisherKey)
	if _, exists := confPub[publisherKey]; exists == true {
		return nil, fmt.Errorf("publisher on '%s' is already registered", pub.Topic)
	}
	confPub[publisherKey] = pub
	config.Set(PublisherKey, confPub)
	godin.SaveConfiguration()

	return &publisher{
		Publishing: pub,
	}, nil
}

func PublisherName(topic string) string {
	name := strings.Replace(topic, ".", "_", -1)
	name = strings.ToLower(name)
	name = strcase.ToCamel(name)

	return name
}

func promptPublisherValues(pub *rabbitmq.Publishing) (protoMessage string, err error) {
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
	pub.Topic = topic

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
	pub.Exchange = exchange

	// Protobuf message name
	p = prompt.NewPrompt(
		"Which protobuf message is going to be sent?",
		"UserCreatedEvent",
		prompt.Validate(),
	)
	protoMessage, err = p.Run()
	if err != nil {
		return "", err
	}

	return protoMessage, nil
}
