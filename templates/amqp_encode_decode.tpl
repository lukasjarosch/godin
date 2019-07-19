package amqp

import (
    "fmt"

    "github.com/go-godin/rabbitmq"

    pb "{{ .Protobuf.Package }}"
)

{{- if gt (len .Service.Publisher) 0 }}
{{- range .Service.Publisher }}
{{- template "amqp_publish_encode" . -}}
{{- end }}
{{- end }}

{{- if gt (len .Service.Subscriber) 0 }}
{{- range .Service.Subscriber }}
{{- template "amqp_subscribe_decode" . -}}
{{- end }}
{{- end }}

