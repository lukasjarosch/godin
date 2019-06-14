{{ define "grpc_decode_response" }}
// Decode{{ .Name }}Response is used in the server and decodes a gRPC response into a domain-level response
func Decode{{ .Name }}Response(ctx context.Context, pbResponse interface{}) (pbResponse interface{}, err error) {
    if pbResponse == nil {
        return nil, errors.New("nil {{ .Name }}Response")
    }
    response, err := {{ .Name }}ResponseDecoder(res)
    if err != nil {
        return nil, err
    }
    return response, nil
}
{{ end }}
