// Code generated by Godin {{ .Godin.Version }}; DO NOT EDIT.
package amqp

import (
	"context"
	"github.com/pkg/errors"

	"github.com/go-godin/log"
	"github.com/go-godin/rabbitmq"
	"github.com/go-godin/grpc-metadata"
	"github.com/go-godin/log"
)

type PublisherSet struct {
	logger log.Logger
    {{- range .Service.Publisher }}
    {{ untitle .Name }} rabbitmq.Publisher
	{{- end }}
}

// Publishers will initialize all AMQP publishers. Each publisher will receive it's own channel to work on
// to reduce TCP pressure in times of high load.
func Publishers(conn *rabbitmq.RabbitMQ, loggger log.Logger) PublisherSet {
	{{- range .Service.Publisher }}
	{{ untitle .Name }}Channel, _ := conn.NewChannel()
	{{- end }}

	return PublisherSet{
		logger:loggger,
		{{- range .Service.Publisher }}
		{{ untitle .Name }}: rabbitmq.NewPublisher(
			{{ untitle .Name }}Channel,
			&rabbitmq.Publishing{
				Topic: "{{ .Publishing.Topic }}",
				Exchange: "{{ .Publishing.Exchange }}",
				DeliveryMode: {{ .Publishing.DeliveryMode }},
			},
		),
		{{- end }}
	}
}

{{- $serviceName := title .Service.Name -}}


{{- range .Service.Publisher }}
func (pub PublisherSet) Publish{{ .Name }}(ctx context.Context, event interface{}) error {
	logger := pub.logger.With("topic", "{{ .Publishing.Topic }}", "exchange", "{{ .Publishing.Exchange }}", "requestId", grpc_metadata.GetRequestID(ctx))

	evt, err := {{ .Name }}Encoder(event)
	if err != nil {
		logger.Error("failed to encode event", "event", event)
		return errors.Wrap(err, "failed to encode event")
	}

	if err := pub.{{ untitle .Name }}.Publish(ctx, evt); err != nil {
		logger.Error("failed to publish event", "err", err)
		return errors.Wrap(err, "failed to publish to '{{ .Publishing.Topic }}'")
	}

	logger.Info("published '{{ .Publishing.Topic }}' event")
	return nil
}
{{- end }}