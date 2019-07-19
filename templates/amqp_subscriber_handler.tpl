package subscriber

import (
	"context"
	"github.com/go-godin/log"
	"github.com/go-godin/rabbitmq"
	grpcMetadata "github.com/go-godin/grpc-metadata"
	"github.com/pkg/errors"

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
// {{ .Handler }}Subscriber is responsible of handling all incoming AMQP messages with routing key '{{ .Subscription.Topic }}'
// It might seem overly complicated at first, but the design is on purpose. You WANT to have access to the Delivery,
// thus it would not make sense to use a middleware for Decoding it into a DAO or domain-level object as you would
// loose access to the Delivery.
func {{ .Handler }}Subscriber(logger log.Logger, usecase service.{{ title $serviceName }}, decoder rabbitmq.SubscriberDecoder) rabbitmq.SubscriptionHandler {
	return func(ctx context.Context, delivery *rabbitmq.Delivery) {
		logger = logger.With(string(grpcMetadata.RequestID), ctx.Value(string(grpcMetadata.RequestID)))

		event, err := decode{{ .Handler }}(delivery, decoder, logger)
		if err != nil {
		    return
		}

		_ = event // remove this line, it just keeps the compiler calm until you start using the event :)

		// TODO: Handle {{ .Subscription.Topic }} subscription here
		/*
			If you want to NACK the delivery, use `delivery.NackDelivery()` instead of Nack().
			This will ensure that the prometheus amqp_nack_counter is increased.
		*/

		_ = delivery.Ack(false)
    }
}

// decode{{ .Handler }} cleans up the actual handler by providing a cleaner interface for decoding incoming {{ .Protobuf.Message }} deliveries.
// It will also take care of logging errors and handling metrics.
func decode{{ .Handler }}(delivery *rabbitmq.Delivery, decoder rabbitmq.SubscriberDecoder, logger log.Logger) (*{{ untitle .Handler }}Proto.{{ .Protobuf.Message }}, error) {
	event, err := decoder(delivery)
	if err != nil {
		if err2 := delivery.NackDelivery(false, false, "{{ .Subscription.Topic }}"); err2 != nil {
			err = errors.Wrap(err, err2.Error())
		}
		delivery.IncrementTransportErrorCounter("{{ .Subscription.Topic }}")
		logger.Error("failed to decode {{ .Protobuf.Message }}", "err", err)
		return nil, err
	}
	logger.Debug("decoded {{ .Protobuf.Message }}", "event", event)

	return event.(*{{ untitle .Handler }}Proto.{{ .Protobuf.Message }}), nil
}
{{ end }}

