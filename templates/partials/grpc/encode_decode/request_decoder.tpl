{{ define "grpc_request_decoder" }}
// {{ .Request }}Decoder maps the protobuf request of the gRPC transport layer onto the domain-level {{ .Request }}
func {{ .Request }}Decoder(pbRequest *pb.{{ .ProtobufRequest }}) (request endpoint.{{ .Request }}, err error) {
    // TODO: map 'pbRequest' to 'request' and return
    return request, err
}
{{ end }}
