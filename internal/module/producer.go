package module

import (
	"fmt"
	"strings"

	prompting "github.com/lukasjarosch/godin/internal/prompt"
	"github.com/sirupsen/logrus"
)

type ProducerConfig struct {
	Exchange string
	Name     string
}

type Producer struct {
}

func NewProducer() *Producer {
	return &Producer{}
}

func (p *Producer) Execute() error {
	logrus.Info("new amqp producer created")

	name, exchange, filename, err := p.prompt()
	if err != nil {
		return err
	}

	logrus.Infof("creating AMQP producer '%s' on exchange '%s': %s", name, exchange, filename)

	// TODO: implement producer module

	return nil
}

func (p *Producer) prompt() (name, exchange, filename string, err error) {

	prompt := prompting.NewPrompt(
		"Enter the producer name (CamelCase)",
		"ExampleProducer",
		prompting.Validate(
			prompting.MinLengthThree(),
		),
	)
	name, err = prompt.Run()
	if err != nil {
		return "", "", "", err
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
		return "", "", "", err
	}

	prompt = prompting.NewPrompt(
		"Enter the filename of the producer:",
		fmt.Sprintf("%s.go", strings.ToLower(name)),
		prompting.Validate(
			prompting.MinLengthThree(),
			prompting.GoSuffix(),
		),
	)
	filename, err = prompt.Run()
	if err != nil {
		return "", "" ,"", err
	}

	return name, exchange, filename, nil
}
