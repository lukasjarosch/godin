{{ define "grpc_decode_request" }}
// Decode{{ .Name }}Request is used in the client and decodes a gRPC request into a domain-level request
func Decode{{ .Name }}Request(ctx context.Context, pbRequest interface{}) (request interface{}, err error) {
    if pbRequest == nil {
        return nil, errors.New("nil {{ .Name }}Request")
    }
    request, err := {{ .Name }}RequestDecoder(req)
    if err != nil {
        return nil, err
    }
    return request, nil
}
{{ end }}