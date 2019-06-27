package amqp

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type SubscriberHandler func(delivery amqp.Delivery)

// Subscription defines all data required to setup an AMQP subscription
// All values, except the ctag are provided by the configuration or inferred by Godin.
type Subscription struct {
	Topic    string            `json:"topic"`
	Exchange string            `json:"exchange"`
	Queue    SubscriptionQueue `json:"queue"`
	AutoAck  bool              `json:"auto_ack"`
	ctag     string            `json:"-"` // generated
}

// SubscriptionQueue configures the queue on which the subscription runs.
type SubscriptionQueue struct {
	Name       string `json:"name"`
	Durable    bool   `json:"durable"`
	AutoDelete bool   `json:"auto_delete"`
	Exclusive  bool   `json:"exclusive"`
	NoWait     bool   `json:"no_wait"`
}

type handler struct {
	implementation SubscriberHandler
	done           chan error
	ctag           string
	routingKey     string
}

// Subscriber handles AMQP subscriptions.
type Subscriber struct {
	channel      *amqp.Channel
	subscription *Subscription
	handlers     []handler
}

// NewSubscriber returns an AMQP publisher
func NewSubscriber(channel *amqp.Channel, subscription *Subscription) Subscriber {

	// generate a unique ctag for the subscriber
	sub := strings.Replace(subscription.Topic, ".", "_", -1)
	ctag := fmt.Sprintf("%s_%s", sub, uuid.New().String())
	subscription.ctag = ctag

	return Subscriber{
		channel:      channel,
		subscription: subscription,
	}
}

// Subscribe will declare the queue defined in the Subscription, bind it to the exchange and start consuming
// by calling the handler in a goroutine.
func (c *Subscriber) Subscribe(handler SubscriberHandler) error {
	queue, err := c.channel.QueueDeclare(
		c.subscription.Queue.Name,
		c.subscription.Queue.Durable,
		c.subscription.Queue.AutoDelete,
		c.subscription.Queue.Exclusive,
		c.subscription.Queue.NoWait,
		nil,
	)
	if err != nil {
		return err
	}

	if err = c.channel.QueueBind(
		queue.Name,
		c.subscription.Topic,
		c.subscription.Exchange,
		c.subscription.Queue.NoWait,
		nil,
	); err != nil {
		return err
	}

	deliveries, err := c.channel.Consume(
		queue.Name,
		"",
		c.subscription.AutoAck,
		c.subscription.Queue.Exclusive,
		false,
		c.subscription.Queue.NoWait,
		nil,
	)
	if err != nil {
		return err
	}

	h := c.setHandler(handler)
	go c.Handler(deliveries, h)

	return nil
}

// setHandler installs a SubscriberHandler to use for this subscription.
func (c *Subscriber) setHandler(handlerImpl SubscriberHandler) handler {
	h := handler{
		routingKey:     c.subscription.Topic,
		ctag:           c.subscription.ctag,
		implementation: handlerImpl,
		done:           make(chan error),
	}
	c.handlers = append(c.handlers, h)

	return h
}

// Handler is started by Subscribe() as Goroutine. For each received AMQP delivery,
// it will call the implementation(delivery) to allow business logic for each delivery to run.
func (c *Subscriber) Handler(deliveries <-chan amqp.Delivery, h handler) {
	for d := range deliveries {
		h.implementation(d)
	}
	h.done <- nil
}

// Shutdown will cancel the subscriber by it's ctag. It needs to be registered
// to a shutdown handler.
func (c *Subscriber) Shutdown() error {
	for _, h := range c.handlers {
		if err := c.channel.Cancel(h.ctag, true); err != nil {
			return err
		}
		<-h.done
	}
	return nil
}
