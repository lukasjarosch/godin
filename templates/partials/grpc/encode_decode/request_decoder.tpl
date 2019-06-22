{{ define "grpc_request_decoder" }}
// {{ .Name }}RequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level {{ .Name }}Request
func {{ .Name }}RequestDecoder(pbRequest *pb.{{ .ProtobufRequest }}) (request endpoint.{{ .Name }}Request, err error) {
    // TODO: map 'pbRequest' to 'request' and return
    return request, err
}
{{ end }}
