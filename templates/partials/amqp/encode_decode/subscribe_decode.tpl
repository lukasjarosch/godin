{{- define "amqp_subscribe_decode" }}
// {{ .Handler }}Decoder
func {{ .Handler }}Decoder(delivery *rabbitmq.Delivery) (decoded interface{}, err error) {
	return nil, fmt.Errorf("{{ .Handler }}Decoder is not implemented")
}
{{- end }}
