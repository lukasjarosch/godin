{{- define "amqp_publish_encode" }}
// {{ .Name }}Encoder is called just before publishing an event to '{{ .Publishing.Topic }}' and encodes
// the DAO to protobuf.
func {{ .Name }}Encoder(event interface{}) (*pb.{{ .Publishing.ProtobufMessage }}, error) {
var encoded pb.{{ .Publishing.ProtobufMessage }}

// TODO: map to protobuf

return &encoded, nil
}
{{- end }}
