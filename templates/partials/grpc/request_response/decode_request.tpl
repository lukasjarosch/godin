{{ define "grpc_decode_request" }}
// Decode{{ .Request }} is used in the client and decodes a gRPC request into a domain-level request
func Decode{{ .Request }}(ctx context.Context, pbRequest interface{}) (request interface{}, err error) {
    if pbRequest == nil {
        return nil, errors.New("nil {{ .Request }}")
    }
    req := pbRequest.(*pb.{{ .ProtobufRequest }})
    request, err = {{ .Request }}Decoder(req)
    if err != nil {
        return nil, err
    }
    return request, nil
}
{{ end }}