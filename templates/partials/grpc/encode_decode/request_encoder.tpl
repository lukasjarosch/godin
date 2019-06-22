{{ define "grpc_request_encoder" }}
// {{ .Request }}Encoder encodes the domain-level {{ .Request }} into a protobuf {{ .ProtobufRequest }}
func {{ .Request }}Encoder(request endpoint.{{ .Request }}) (pbRequest *pb.{{ .ProtobufRequest }}, err error) {
    // TODO: map 'request' to 'pbRequest' and return
    return pbRequest, err
}
{{ end }}