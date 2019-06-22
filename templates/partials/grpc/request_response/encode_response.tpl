{{ define "grpc_encode_response" }}
// Encode{{ .Response }} is used in the client and encodes a domain-level response into a gRPC response
func Encode{{ .Response }}(ctx context.Context, response interface{}) (pbResponse interface{}, err error) {
    if response == nil {
        return nil, errors.New("nil {{ .Response }}")
    }
    res := response.(endpoint.{{ .Response }})
    pbResponse, err := {{ .Response }}Encoder(res)
    if err != nil {
        return nil, err
    }
    return pbResponse, nil
}
{{ end }}
