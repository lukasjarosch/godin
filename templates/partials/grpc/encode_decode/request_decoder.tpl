{{ define "grpc_request_decoder" }}
// {{ .Name }}RequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level {{ .Name }}Request
func {{ .Name }}RequestDecoder(pbRequest *pb.{{ .Name }}Request) (request endpoint.{{ .Name }}Request, err error) {
    // TODO: map 'pbRequest' to 'request' and return; adjust the protobuf types as required, they may not be correct
}
{{ end }}
