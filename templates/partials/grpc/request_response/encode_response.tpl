{{ define "grpc_encode_response" }}
// Encode{{ .Name }}Response is used in the client and encodes a domain-level response into a gRPC response
func Encode{{ .Name }}Response(ctx context.Context, response interface{}) (pbResponse interface{}, err error) {
    if response == nil {
        return nil, errors.New("nil {{ .Name }}Response")
    }
    res := response.(endpoint.{{ .Name }}Response)
    pbResponse, err := {{ .Name }}ResponseEncoder(res)
    if err != nil {
        return nil, err
    }
    return pbResponse, nil
}
{{ end }}
