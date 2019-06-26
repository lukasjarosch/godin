package transport

import (
	"fmt"

	"github.com/lukasjarosch/godin/internal/godin"
	"github.com/lukasjarosch/godin/internal/prompt"
	"github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
)

const AMQPTransportKey = "transport.amqp"
const AMQPDefaultAddressKey = "transport.amqp.default_address"

type amqpTransport struct {
	DefaultAddress string `json:"default_address"`
	Exchange       string `json:"exchange"`
	ExchangeType   string `json:"exchange_type"`
}

// InitializeAMQP will try to load the amqpTransport struct from the configuration.
// If that fails, the user will be prompted for the required values.
// The loaded or entered data is then saved to the config before returning.
func InitializeAMQP() (transport *amqpTransport, err error) {
	// Load AMQP from config, if that fails create it and prompt the user for the values
	transport, err = AMQPFromConfig()
	if err != nil {
		transport = &amqpTransport{}
		if err := promptAmqpTransportValues(transport); err != nil {
			return nil, fmt.Errorf("error while prompting AMQP transport values: %s", err)
		}
	}

	config.Set(AMQPDefaultAddressKey, transport.DefaultAddress)
	godin.SaveConfiguration()

	return transport, nil
}

// Try and load the amqpTransport from the configuration.
// Returns an error if the config key "transport.amqp" does not exist.
// If any sub-key is not set properly, it will not be checked as it's still the
// developers responsibility to provide sane values.
func AMQPFromConfig() (*amqpTransport, error) {
	if config.IsSet(AMQPTransportKey) {
		defaultAddress := config.GetString(AMQPDefaultAddressKey)
		logrus.Debug("AMQP transport is already configured, loading from config")

		return &amqpTransport{
			DefaultAddress: defaultAddress,
		}, nil
	}
	return nil, fmt.Errorf("amqp transport not configured")
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
