{{ define "grpc_response_encoder" }}
// {{ .Response }}Encoder encodes the domain-level {{ .Response }} into a protobuf {{ .ProtobufResponse }}
func {{ .Response }}Encoder(response endpoint.{{ .Response }}) (pbResponse *pb.{{ .ProtobufResponse }}, err error) {
    // TODO: map 'response' to 'pbResponse' and return
    return pbResponse, err
}
{{ end }}
