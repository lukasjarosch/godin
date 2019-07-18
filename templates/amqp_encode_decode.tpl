package amqp

import (
    "github.com/go-godin/rabbitmq"
    pb "{{ .Protobuf.Package }}"
)

{{- if gt (len .Service.Subscriber) 0 }}
type SubcriptionDecoder func(delivery *rabbitmq.Delivery) (decoded interface{}, err error)
{{- end }}

{{- if gt (len .Service.Publisher) 0 }}
{{- range .Service.Publisher }}
{{- template "amqp_publish_encode" . }}
{{- end }}
{{- end }}

