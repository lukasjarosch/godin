{{ define "grpc_response_decoder" }}
// {{ .Name }}ResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level {{ .Name }}Response
func {{ .Name }}ResponseDecoder(pbResponse *pb.{{ .Name }}Response) (response endpoint.{{ .Name }}Response, err error) {
    // TODO: map 'pbResponse' to 'response' and return; adjust the protobuf types as required, they may not be correct
}
{{ end }}
