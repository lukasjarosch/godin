package subscriber

import (
	"context"
	"github.com/go-godin/log"
	"github.com/go-godin/rabbitmq"
	grpc_metadata "github.com/go-godin/grpc-metadata"

    "{{ .Service.Module }}/internal/service"
    {{- range .Service.Subscriber }}
    {{ untitle .Handler }}Proto "{{ .Protobuf.Import }}"
    {{- end }}
)

{{- $serviceName := title .Service.Name -}}

{{/*
Loop over all subscribers. The slice is faked, tho. The generator will replace the slice with the current
subscriber only, so we can safely assume that there is only one element in the slice.
*/}}
{{ range .Service.Subscriber }}
// {{ .Handler }} is responsible of handling all incoming AMQP messages with routing key '{{ .Subscription.Topic }}'
func {{ .Handler }}Subscriber(logger log.Logger, usecase service.{{ title $serviceName }}, decoder rabbitmq.SubscriberDecoder) rabbitmq.SubscriptionHandler {
	return func(ctx context.Context, delivery *rabbitmq.Delivery) {
		// the requestId is injected into the context and should be attached on every log
		logger = logger.With(string(grpc_metadata.RequestID), ctx.Value(string(grpc_metadata.RequestID)))

		event, err := decoder(delivery)
		event = event.({{ untitle .Handler }}Proto.{{ .Protobuf.Message }})
		if err != nil {
		    logger.Error("failed to decode '{{ .Subscription.Topic }}' event", "err", err)
		    delivery.NackDelivery(false, false, "{{ .Subscription.Topic }}")
		    delivery.IncrementTransportErrorCounter("{{ .Subscription.Topic }}")
		    return
		}

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

