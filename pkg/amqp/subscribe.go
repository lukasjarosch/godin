package amqp

import (
	"github.com/streadway/amqp"
)

type SubscriberHandler func(delivery amqp.Delivery)

type Subscription struct {
	Topic          string `json:"topic"`
	Exchange       string `json:"exchange"`
	Queue          string `json:"queue"`
}

type handler struct {
	implementation SubscriberHandler
	done           chan error
	ctag           string
	routingKey     string
}

type Subscriber struct {
	channel  *amqp.Channel
	exchange string
	handlers []handler
}

// NewSubscriber returns an AMQP publisher
func NewSubscriber(channel *amqp.Channel, exchange string) Subscriber {
	return Subscriber{
		channel:  channel,
		exchange: exchange,
	}
}

func (c *Subscriber) Subscribe(queueName, routingKey string, ctag string, handler SubscriberHandler) error {
	queue, err := c.channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	if err = c.channel.QueueBind(queue.Name, routingKey, c.exchange, false, nil); err != nil {
		return err
	}

	deliveries, err := c.channel.Consume(queue.Name, ctag, false, false, false, false, nil)
	if err != nil {
		return err
	}

	h := c.addHandler(routingKey, ctag, handler)
	go c.Handler(deliveries, h)

	return nil
}

func (c *Subscriber) addHandler(routingKey string, ctag string, handlerImpl SubscriberHandler) handler {
	h := handler{
		routingKey:     routingKey,
		ctag:           ctag,
		implementation: handlerImpl,
		done:           make(chan error),
	}
	c.handlers = append(c.handlers, h)

	return h
}

func (c *Subscriber) Handler(deliveries <-chan amqp.Delivery, h handler) {
	for d := range deliveries {
		h.implementation(d)
	}
	h.done <- nil
}

func (c *Subscriber) Shutdown() error {
	for _, h := range c.handlers {
		if err := c.channel.Cancel(h.ctag, true); err != nil {
			return err
		}
		<-h.done
	}
	return nil
}
