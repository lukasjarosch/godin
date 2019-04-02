package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/lukasjarosch/godin-api-go/godin/greeter/v1beta1"
	service "github.com/lukasjarosch/godin/examples/spec-greeter/internal/greeter"
)

// greeterHandler is the transport-layer wrapper of our business-logic in the server package
// Everything concerning requests/responses belongs in here. Only conversion (business-model <-> protobuf) should happen here actually.
type greeterHandler struct {
	implementation *service.Greeter
}

func NewGreeterAPIHandler(implementation *service.Greeter) *greeterHandler {
	return &greeterHandler{
		implementation: implementation,
	}
}

// Hello is the gRPC handler for godin.greeter.v1beta1.Hello()
func (e *greeterHandler) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	// TODO: call e.implementation.Hello and return the response
	return nil, status.Error(codes.Unimplemented, "godin.greeter.v1beta1.GreeterAPI.Hello() unimplemented")
}
