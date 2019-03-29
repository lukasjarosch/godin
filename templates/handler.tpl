package server

import (
	"context"

	greeter "github.com/lukasjarosch/godin-api-go/godin/greeter/v1beta1/"
	service "{{ .ModuleName }}/internal/{{ .ServiceName }}"
)

// TODO: Implement all handlers for your gRPC service

// greeterAPIHandler is the transport-layer wrapper of our business-logic in the server package
// Everything concerning requests/responses belongs in here. Only conversion (business-model <-> protobuf) should happen here actually.
type greeterAPIHandler struct {
	implementation *service.GreeterService
}

func NewGreeterAPIHandler(implementation *service.GreeterService) *greeterAPIHandler{
	return &greeterAPIHandler{
		implementation: implementation,
	}
}

func (e *greeterAPIHandler) Greeting(ctx context.Context, request *greeter.GreetingRequest) (*greeter.GreetingResponse, error) {
	greeting, err := e.implementation.Greeting(request.Name)
	if err != nil {
		return nil, err
	}

	return &greeter.GreetingResponse{Greeting: greeting}, nil
}

