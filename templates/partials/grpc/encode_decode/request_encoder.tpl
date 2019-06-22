{{ define "grpc_request_encoder" }}
// {{ .Name }}RequestEncoder encodes the domain-level {{ .Name }}Request into a protobuf {{ .Name }}Request
func {{ .Name }}RequestEncoder(request endpoint.{{ .Name }}Request) (pbRequest *pb.{{ .Name }}Request, err error) {
    // TODO: map 'request' to 'pbRequest' and return; adjust the protobuf types as required, they may not be correct
    return pbRequest, err
}
{{ end }}