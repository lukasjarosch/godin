{{- define "amqp_subscribe_decode" }}
// {{ .Handler }}Decoder will unmarshal the Body of the incoming Delivery into a protobuf message.
//
// Note: Godin will not regenerate this file, only append new functions. So if this file was already present when
// you added the subscriber, you need to fix the imports by adding:
//	 {{ untitle .Handler }}Proto "{{ .Protobuf.Import }}"
func {{ .Handler }}Decoder(delivery *rabbitmq.Delivery) (decoded interface{}, err error) {
	var event *{{ untitle .Handler }}Proto.{{ .Protobuf.Message }}

	if err := proto.Unmarshal(delivery.Body, event); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Delivery into *{{ untitle .Handler }}Proto.{{ .Protobuf.Message }}: %s", err)
	}

	return event, nil
}
{{- end }}
