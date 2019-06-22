{{ define "grpc_response_decoder" }}
// {{ .Response }}Decoder maps the protobuf response of the gRPC transport layer onto the domain-level {{ .Response }}
func {{ .Response }}Decoder(pbResponse *pb.{{ .ProtobufResponse }}) (response endpoint.{{ .Response }}, err error) {
    // TODO: map 'pbResponse' to 'response' and return
    return response, err
}
{{ end }}
