{{ define "grpc_decode_response" }}
// Decode{{ .Response }} is used in the server and decodes a gRPC response into a domain-level response
func Decode{{ .Response }}(ctx context.Context, pbResponse interface{}) (response interface{}, err error) {
    if pbResponse == nil {
        return nil, errors.New("nil {{ .Response }}")
    }
    response, err := {{ .Response }}Decoder(res)
    if err != nil {
        return nil, err
    }
    return response, nil
}
{{ end }}
