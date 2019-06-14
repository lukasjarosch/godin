{{ define "grpc_response_encoder" }}
// {{ .Name }}ResponseEncoder encodes the domain-level {{ .Name }}Response into a protobuf {{ .Name }}Response
func {{ .Name }}ResponseEncoder(response endpoint.{{ .Name }}Response) (pbResponse *pb.{{ .Name }}Response, err error) {
    // TODO: map 'response' to 'pbResponse' and return; adjust the protobuf types as required, they may not be correct
}
{{ end }}
