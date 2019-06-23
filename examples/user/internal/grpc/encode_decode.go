package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	service "github.com/lukasjarosch/godin/examples/user"
	"github.com/lukasjarosch/godin/examples/user/endpoint"
	pb "godin.user"
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
// CreateRequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level CreateRequest
func CreateRequestDecoder(pbRequest *pb.CreateRequest) (request endpoint.CreateRequest, err error) {
	// TODO: map 'pbRequest' to 'request' and return
	return request, err
}

// CreateResponseEncoder encodes the domain-level CreateResponse into a protobuf CreateResponse
func CreateResponseEncoder(response endpoint.CreateResponse) (pbResponse *pb.CreateResponse, err error) {
	// TODO: map 'response' to 'pbResponse' and return
	return pbResponse, err
}

// CreateRequestEncoder encodes the domain-level CreateRequest into a protobuf CreateRequest
func CreateRequestEncoder(request endpoint.CreateRequest) (pbRequest *pb.CreateRequest, err error) {
	// TODO: map 'request' to 'pbRequest' and return
	return pbRequest, err
}

// CreateResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level CreateResponse
func CreateResponseDecoder(pbResponse *pb.CreateResponse) (response endpoint.CreateResponse, err error) {
	// TODO: map 'pbResponse' to 'response' and return
	return response, err
}
