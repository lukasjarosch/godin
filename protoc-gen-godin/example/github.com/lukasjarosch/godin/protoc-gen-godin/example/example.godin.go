// Code generated by protoc-gen-godin. DO NOT EDIT.
// source: example.proto

package example



// EncodeFooRequest encodes a domain-level FooRequest into the protobuf message 'godin.example.FooResponse'
// The actual value-mapping is done in FooRequestEncoder()
// This method is used in the gRPC server
func EncodeFooRequest(ctx context.Context, request interface{}) (interface{}, error) {
    if request == nil {
        return nil, errors.New("nil FooRequest")
    }
    req := request.(transport.FooRequest)
    pbRequest, err := FooRequestEncoder(req)
    if err != nil {
        return nil, err
    }

    return pbRequest, nil
}

// DecodeFooResponse decodes a protobuf 'godin.example.FooResponse' into a domain-level FooResponse
// The actual value-mapping is done in FooResponseDecoder()
// This method is used in the gRPC server
func DecodeFooResponse(ctx context.Context, pbResponse interface{}) (interface{}, error) {
    godin.example.FooResponse
}
// EncodeBarRequest encodes a domain-level BarRequest into the protobuf message 'godin.example.BarResponse'
// The actual value-mapping is done in BarRequestEncoder()
// This method is used in the gRPC server
func EncodeBarRequest(ctx context.Context, request interface{}) (interface{}, error) {
    if request == nil {
        return nil, errors.New("nil BarRequest")
    }
    req := request.(transport.BarRequest)
    pbRequest, err := BarRequestEncoder(req)
    if err != nil {
        return nil, err
    }

    return pbRequest, nil
}

// DecodeBarResponse decodes a protobuf 'godin.example.BarResponse' into a domain-level BarResponse
// The actual value-mapping is done in BarResponseDecoder()
// This method is used in the gRPC server
func DecodeBarResponse(ctx context.Context, pbResponse interface{}) (interface{}, error) {
    godin.example.BarResponse
}

