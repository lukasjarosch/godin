{{ define "grpc_encode_request" }}
// Encode{{ .Name }}Request is used in the server and encodes a domain-level request into a gRPC request
func Encode{{ .Name }}Request(ctx context.Context, request interface{}) (pbRequest interface{}, err error) {
    if request == nil {
        return nil, errors.New("nil {{ .Name }}Request")
    }
    req := request.(endpoint.{{ .Name }}Request)
    pbRequest, err := {{ .Name }}RequestEncoder(req)
    if err != nil {
        return nil, err
    }
    return pbRequest, nil
}
{{ end }}
