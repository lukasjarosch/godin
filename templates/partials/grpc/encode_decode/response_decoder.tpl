{{ define "grpc_response_decoder" }}
// {{ .Name }}ResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level {{ .Name }}Response
func {{ .Name }}ResponseDecoder(pbResponse *pb.{{ .ProtobufResponse }}) (response endpoint.{{ .Name }}Response, err error) {
    // TODO: map 'pbResponse' to 'response' and return
    return response, err
}
{{ end }}
