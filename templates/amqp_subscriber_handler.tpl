package subscriber

import (
	"context"
	"github.com/go-godin/log"
	"github.com/go-godin/rabbitmq"
	grpc_metadata "github.com/go-godin/grpc-metadata"

    "{{ .Service.Module }}/internal/service"
)

{{- $serviceName := title .Service.Name -}}

{{/*
Loop over all subscribers. The slice is faked, tho. The generator will replace the slice with the current
subscriber only, so we can safely assume that there is only one element in the slice.
*/}}
{{ range .Service.Subscriber }}
// {{ .Handler }} is responsible of handling all incoming AMQP messages with routing key '{{ .Subscription.Topic }}'
func {{ .Handler }}(logger log.Logger, usecase service.{{ title $serviceName }}) rabbitmq.SubscriptionHandler {
	return func(ctx context.Context, delivery *rabbitmq.Delivery) {
		// the requestId is injected into the context and should be attached on every log
		logger = logger.With(ctx.Value(string(grpc_metadata.RequestID)), ctx.Value(string(grpc_metadata.RequestID)))

		// TODO: Handle {{ .Subscription.Topic }} subscription
		/*
			If you want to NACK the delivery, use `delivery.NackDelivery()` instead of Nack().
			This will ensure that the prometheus amqp_nack_counter is increased.

			Godins delivery wrapper also provides a `delivery.IncrementTransportErrorCounter()` method to grant
			you access to the amqp_transport_error metric. Call it if the message is incomplete or cannot
			be unmarshalled for any reason.
		*/

		_ = delivery.Ack(false)
    }
}
{{ end }}

{{/*

func UserCreatedHandler(logger log.Logger, usecase service.Hello) rabbitmq.SubscriptionHandler {
	return func(ctx context.Context, delivery *rabbitmq.Delivery) {
		logger = logger.With("requestId", ctx.Value("requestId"))
		logger.Info("calling usecase.Hello")

		// TODO: unmarshal
		// delivery.IncrementTransportErrorCounter("user.created")

		greeting, err := usecase.Hello(ctx, "derp")
		if err != nil {
		    logger.Error("failed to call usecase.Hello()", "err", err)
			delivery.NackDelivery(false, false, "user.created")
		    return
		}

		logger.Info(greeting)
		delivery.Ack(false)
	}
}
*/}}
