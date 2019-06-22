{{ define "grpc_response_encoder" }}
// {{ .Name }}ResponseEncoder encodes the domain-level {{ .Name }}Response into a protobuf {{ .ProtobufResponse }}
func {{ .Name }}ResponseEncoder(response endpoint.{{ .Name }}Response) (pbResponse *pb.{{ .ProtobufResponse }}, err error) {
    // TODO: map 'response' to 'pbResponse' and return
    return pbResponse, err
}
{{ end }}
