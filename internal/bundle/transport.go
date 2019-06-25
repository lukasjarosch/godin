package bundle

import (
	"fmt"

	"github.com/lukasjarosch/godin/internal/prompt"
	"github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
	"github.com/lukasjarosch/godin/internal/godin"
)

const AmqpTransportKey = "transport.amqp"
const AmqpExchangeKey = "transport.amqp.exchange"
const AmqpExchangeTypeKey = "transport.amqp.exchange_type"
const AmqpDefaultAddress = "transport.amqp.default_address"

type amqpTransport struct {
	DefaultAddress string `json:"default_address"`
	Exchange       string `json:"exchange"`
	ExchangeType   string `json:"exchange_type"`
}

func InitializeAMQPTransport() (transport *amqpTransport, err error) {
	if config.IsSet(AmqpTransportKey) {
		exchange := config.GetString(AmqpExchangeKey)
		exchangeType := config.GetString(AmqpExchangeTypeKey)
		defaultAddress := config.GetString(AmqpDefaultAddress)

		logrus.Debug("AMQP transport is already configured, loading from config")

		return &amqpTransport{
			DefaultAddress: defaultAddress,
			Exchange:       exchange,
			ExchangeType:   exchangeType,
		}, nil
	}

	transport = &amqpTransport{}
	if err := promptAmqpTransportValues(transport); err != nil {
		return nil, fmt.Errorf("error while prompting AMQP transport values: %s", err)
	}

	config.Set(AmqpExchangeKey, transport.Exchange)
	config.Set(AmqpExchangeTypeKey, transport.ExchangeType)
	config.Set(AmqpDefaultAddress, transport.DefaultAddress)
	godin.SaveConfiguration()

	return transport, nil
}

func promptAmqpTransportValues(transport *amqpTransport) (err error) {
	// default AMQP server address
	p := prompt.NewPrompt(
		"AMQP default server address to use if AMQP_URL is not provided",
		"amqp://username:password@host:port/vhost",
		prompt.Validate(),
	)
	transport.DefaultAddress, err = p.Run()
	if err != nil {
		return err
	}

	return nil
}
