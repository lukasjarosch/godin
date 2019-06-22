{{ define "grpc_request_encoder" }}
// {{ .Name }}RequestEncoder encodes the domain-level {{ .Name }}Request into a protobuf {{ .ProtobufRequest }}
func {{ .Name }}RequestEncoder(request endpoint.{{ .Name }}Request) (pbRequest *pb.{{ .ProtobufRequest }}, err error) {
    // TODO: map 'request' to 'pbRequest' and return
    return pbRequest, err
}
{{ end }}