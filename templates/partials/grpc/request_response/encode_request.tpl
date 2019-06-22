{{ define "grpc_encode_request" }}
// Encode{{ .Request }} is used in the server and encodes a domain-level request into a gRPC request
func Encode{{ .Request }}(ctx context.Context, request transport.{{ .Request }}) (pbRequest *pb.{{ .ProtobufRequest }}, err error) {
    if request == nil {
        return nil, errors.New("nil {{ .Request }}")
    }
    req := request.(endpoint.{{ .Request }})
    pbRequest, err := {{ .Request }}Encoder(req)
    if err != nil {
        return nil, err
    }
    return pbRequest, nil
}
{{ end }}
