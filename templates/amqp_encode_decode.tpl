package amqp

import (
    pb "{{ .Protobuf.Package }}"
)


{{- if gt (len .Service.Publisher) 0 }}
{{- range .Service.Publisher }}
{{- template "amqp_publish_encode" . }}
{{- end }}
{{- end }}
