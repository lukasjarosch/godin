package module

import (
	"github.com/lukasjarosch/godin/internal/project"
	prompting "github.com/lukasjarosch/godin/internal/prompt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ProducerConfig struct {
	Exchange string
	Name string
}

type Producer struct {
}

func NewProducer() *Producer {
	return &Producer{}
}

func (p *Producer) Execute() error {
	logrus.Info("new amqp producer created")


	name, exchange, err := p.prompt()
	if err != nil {
	    return err
	}

	if err := p.configure(name, exchange); err != nil {
		return err
	}

	// TODO: implement producer module


	return nil
}

func (p *Producer) configure(name, exchange string) error {
	producers := viper.GetStringSlice("producers.registered")
	viper.Set("producers.registered", append(producers, name))
	viper.Set("producers." + name + ".exchange", exchange)
	project.SaveConfig()

	return nil
}

func (p *Producer) prompt() (name, exchange string, err error) {

	prompt := prompting.NewPrompt(
		"Enter the producer name (CamelCase)",
		"ExampleProducer",
		prompting.Validate(
			prompting.MinLengthThree(),
		),
	)
	name, err = prompt.Run()
	if err != nil {
		return "", "", err
	}

	prompt = prompting.NewPrompt(
		"Enter the AMQP exchange name ",
		"",
		prompting.Validate(
			prompting.MinLengthThree(),
		),
	)
	exchange, err = prompt.Run()
	if err != nil {
		return "", "", err
	}

	return name, exchange, nil
}
