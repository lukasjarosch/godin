package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	service "yyy"
	pb "yyy/api"
	"yyy/endpoint"
)

// ----------------[ ERRORS ]----------------

// EncodeError encodes domain-level errors into gRPC transport-level errors
func EncodeError(err error) error {
	switch err {
	default:
		return status.Error(codes.Unknown, err.Error())
	}
	return err
}

// ----------------[ MAPPING FUNCS ]----------------

// TODO: this is a nice spot for convenience mapping functions :)

// ----------------[ ENCODER / DECODER ]----------------
// HelloRequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level HelloRequest
func HelloRequestDecoder(pbRequest *pb.HelloRequest) (request endpoint.HelloRequest, err error) {
	// TODO: map 'pbRequest' to 'request' and return; adjust the protobuf types as required, they may not be correct
	return request, err
}

// HelloResponseEncoder encodes the domain-level HelloResponse into a protobuf HelloResponse
func HelloResponseEncoder(response endpoint.HelloResponse) (pbResponse *pb.HelloResponse, err error) {
	// TODO: map 'response' to 'pbResponse' and return; adjust the protobuf types as required, they may not be correct
	return pbResponse, err
}

// HelloRequestEncoder encodes the domain-level HelloRequest into a protobuf HelloRequest
func HelloRequestEncoder(request endpoint.HelloRequest) (pbRequest *pb.HelloRequest, err error) {
	// TODO: map 'request' to 'pbRequest' and return; adjust the protobuf types as required, they may not be correct
	return pbRequest, err
}

// HelloResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level HelloResponse
func HelloResponseDecoder(pbResponse *pb.HelloResponse) (response endpoint.HelloResponse, err error) {
	// TODO: map 'pbResponse' to 'response' and return; adjust the protobuf types as required, they may not be correct
	return response, err
}

// Hello2RequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level Hello2Request
func Hello2RequestDecoder(pbRequest *pb.Hello2Request) (request endpoint.Hello2Request, err error) {
	// TODO: map 'pbRequest' to 'request' and return; adjust the protobuf types as required, they may not be correct
	return request, err
}

// Hello2ResponseEncoder encodes the domain-level Hello2Response into a protobuf Hello2Response
func Hello2ResponseEncoder(response endpoint.Hello2Response) (pbResponse *pb.Hello2ULTRAResponse, err error) {
	// TODO: map 'response' to 'pbResponse' and return; adjust the protobuf types as required, they may not be correct
	return pbResponse, err
}

// Hello2RequestEncoder encodes the domain-level Hello2Request into a protobuf Hello2Request
func Hello2RequestEncoder(request endpoint.Hello2Request) (pbRequest *pb.Hello2Request, err error) {
	// TODO: map 'request' to 'pbRequest' and return; adjust the protobuf types as required, they may not be correct
	return pbRequest, err
}

// Hello2ResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level Hello2Response
func Hello2ResponseDecoder(pbResponse *pb.Hello2ULTRAResponse) (response endpoint.Hello2Response, err error) {
	// TODO: map 'pbResponse' to 'response' and return; adjust the protobuf types as required, they may not be correct
	return response, err
}

// Hello3RequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level Hello3Request
func Hello3RequestDecoder(pbRequest *pb.Hello3Request) (request endpoint.Hello3Request, err error) {
	// TODO: map 'pbRequest' to 'request' and return; adjust the protobuf types as required, they may not be correct
	return request, err
}

// Hello3ResponseEncoder encodes the domain-level Hello3Response into a protobuf Hello3Response
func Hello3ResponseEncoder(response endpoint.Hello3Response) (pbResponse *pb.Hello3Response, err error) {
	// TODO: map 'response' to 'pbResponse' and return; adjust the protobuf types as required, they may not be correct
	return pbResponse, err
}

// Hello3RequestEncoder encodes the domain-level Hello3Request into a protobuf Hello3Request
func Hello3RequestEncoder(request endpoint.Hello3Request) (pbRequest *pb.Hello3Request, err error) {
	// TODO: map 'request' to 'pbRequest' and return; adjust the protobuf types as required, they may not be correct
	return pbRequest, err
}

// Hello3ResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level Hello3Response
func Hello3ResponseDecoder(pbResponse *pb.Hello3Response) (response endpoint.Hello3Response, err error) {
	// TODO: map 'pbResponse' to 'response' and return; adjust the protobuf types as required, they may not be correct
	return response, err
}

// Hello4RequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level Hello4Request
func Hello4RequestDecoder(pbRequest *pb.Hello4Request) (request endpoint.Hello4Request, err error) {
	// TODO: map 'pbRequest' to 'request' and return; adjust the protobuf types as required, they may not be correct
	return request, err
}

// Hello4ResponseEncoder encodes the domain-level Hello4Response into a protobuf Hello4Response
func Hello4ResponseEncoder(response endpoint.Hello4Response) (pbResponse *pb.Hello4Response, err error) {
	// TODO: map 'response' to 'pbResponse' and return; adjust the protobuf types as required, they may not be correct
	return pbResponse, err
}

// Hello4RequestEncoder encodes the domain-level Hello4Request into a protobuf Hello4Request
func Hello4RequestEncoder(request endpoint.Hello4Request) (pbRequest *pb.Hello4Request, err error) {
	// TODO: map 'request' to 'pbRequest' and return; adjust the protobuf types as required, they may not be correct
	return pbRequest, err
}

// Hello4ResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level Hello4Response
func Hello4ResponseDecoder(pbResponse *pb.Hello4Response) (response endpoint.Hello4Response, err error) {
	// TODO: map 'pbResponse' to 'response' and return; adjust the protobuf types as required, they may not be correct
	return response, err
}

// Hello5RequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level Hello5Request
func Hello5RequestDecoder(pbRequest *pb.Hello5Request) (request endpoint.Hello5Request, err error) {
	// TODO: map 'pbRequest' to 'request' and return; adjust the protobuf types as required, they may not be correct
	return request, err
}

// Hello5ResponseEncoder encodes the domain-level Hello5Response into a protobuf Hello5Response
func Hello5ResponseEncoder(response endpoint.Hello5Response) (pbResponse *pb.Hello5Response, err error) {
	// TODO: map 'response' to 'pbResponse' and return; adjust the protobuf types as required, they may not be correct
	return pbResponse, err
}

// Hello5RequestEncoder encodes the domain-level Hello5Request into a protobuf Hello5Request
func Hello5RequestEncoder(request endpoint.Hello5Request) (pbRequest *pb.Hello5Request, err error) {
	// TODO: map 'request' to 'pbRequest' and return; adjust the protobuf types as required, they may not be correct
	return pbRequest, err
}

// Hello5ResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level Hello5Response
func Hello5ResponseDecoder(pbResponse *pb.Hello5Response) (response endpoint.Hello5Response, err error) {
	// TODO: map 'pbResponse' to 'response' and return; adjust the protobuf types as required, they may not be correct
	return response, err
}

// Hello6RequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level Hello6Request
func Hello6RequestDecoder(pbRequest *pb.Hello6Request) (request endpoint.Hello6Request, err error) {
	// TODO: map 'pbRequest' to 'request' and return; adjust the protobuf types as required, they may not be correct
	return request, err
}

// Hello6ResponseEncoder encodes the domain-level Hello6Response into a protobuf Hello6Response
func Hello6ResponseEncoder(response endpoint.Hello6Response) (pbResponse *pb.Hello6Response, err error) {
	// TODO: map 'response' to 'pbResponse' and return; adjust the protobuf types as required, they may not be correct
	return pbResponse, err
}

// Hello6RequestEncoder encodes the domain-level Hello6Request into a protobuf Hello6Request
func Hello6RequestEncoder(request endpoint.Hello6Request) (pbRequest *pb.Hello6Request, err error) {
	// TODO: map 'request' to 'pbRequest' and return; adjust the protobuf types as required, they may not be correct
	return pbRequest, err
}

// Hello6ResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level Hello6Response
func Hello6ResponseDecoder(pbResponse *pb.Hello6Response) (response endpoint.Hello6Response, err error) {
	// TODO: map 'pbResponse' to 'response' and return; adjust the protobuf types as required, they may not be correct
	return response, err
}

// Hello7RequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level Hello7Request
func Hello7RequestDecoder(pbRequest *pb.Hello7Request) (request endpoint.Hello7Request, err error) {
	// TODO: map 'pbRequest' to 'request' and return; adjust the protobuf types as required, they may not be correct
	return request, err
}

// Hello7ResponseEncoder encodes the domain-level Hello7Response into a protobuf Hello7Response
func Hello7ResponseEncoder(response endpoint.Hello7Response) (pbResponse *pb.Hello7Response, err error) {
	// TODO: map 'response' to 'pbResponse' and return; adjust the protobuf types as required, they may not be correct
	return pbResponse, err
}

// Hello7RequestEncoder encodes the domain-level Hello7Request into a protobuf Hello7Request
func Hello7RequestEncoder(request endpoint.Hello7Request) (pbRequest *pb.Hello7Request, err error) {
	// TODO: map 'request' to 'pbRequest' and return; adjust the protobuf types as required, they may not be correct
	return pbRequest, err
}

// Hello7ResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level Hello7Response
func Hello7ResponseDecoder(pbResponse *pb.Hello7Response) (response endpoint.Hello7Response, err error) {
	// TODO: map 'pbResponse' to 'response' and return; adjust the protobuf types as required, they may not be correct
	return response, err
}

// Hello8RequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level Hello8Request
func Hello8RequestDecoder(pbRequest *pb.Hello8Request) (request endpoint.Hello8Request, err error) {
	// TODO: map 'pbRequest' to 'request' and return; adjust the protobuf types as required, they may not be correct
	return request, err
}

// Hello8ResponseEncoder encodes the domain-level Hello8Response into a protobuf Hello8Response
func Hello8ResponseEncoder(response endpoint.Hello8Response) (pbResponse *pb.Hello8Response, err error) {
	// TODO: map 'response' to 'pbResponse' and return; adjust the protobuf types as required, they may not be correct
	return pbResponse, err
}

// Hello8RequestEncoder encodes the domain-level Hello8Request into a protobuf Hello8Request
func Hello8RequestEncoder(request endpoint.Hello8Request) (pbRequest *pb.Hello8Request, err error) {
	// TODO: map 'request' to 'pbRequest' and return; adjust the protobuf types as required, they may not be correct
	return pbRequest, err
}

// Hello8ResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level Hello8Response
func Hello8ResponseDecoder(pbResponse *pb.Hello8Response) (response endpoint.Hello8Response, err error) {
	// TODO: map 'pbResponse' to 'response' and return; adjust the protobuf types as required, they may not be correct
	return response, err
}

// Hello9RequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level Hello9Request
func Hello9RequestDecoder(pbRequest *pb.Hello9Request) (request endpoint.Hello9Request, err error) {
	// TODO: map 'pbRequest' to 'request' and return; adjust the protobuf types as required, they may not be correct
	return request, err
}

// Hello9ResponseEncoder encodes the domain-level Hello9Response into a protobuf Hello9Response
func Hello9ResponseEncoder(response endpoint.Hello9Response) (pbResponse *pb.Hello9Response, err error) {
	// TODO: map 'response' to 'pbResponse' and return; adjust the protobuf types as required, they may not be correct
	return pbResponse, err
}

// Hello9RequestEncoder encodes the domain-level Hello9Request into a protobuf Hello9Request
func Hello9RequestEncoder(request endpoint.Hello9Request) (pbRequest *pb.Hello9Request, err error) {
	// TODO: map 'request' to 'pbRequest' and return; adjust the protobuf types as required, they may not be correct
	return pbRequest, err
}

// Hello9ResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level Hello9Response
func Hello9ResponseDecoder(pbResponse *pb.Hello9Response) (response endpoint.Hello9Response, err error) {
	// TODO: map 'pbResponse' to 'response' and return; adjust the protobuf types as required, they may not be correct
	return response, err
}
