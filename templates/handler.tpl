package server

import (
	"context"

	greeter "github.com/lukasjarosch/godin-api-go/godin/greeter/v1beta1"
	service "{{ .ModuleName }}/internal/{{ .ServiceName }}"
)

// TODO: Implement all handlers for your gRPC service

// greeterAPIHandler is the transport-layer wrapper of our business-logic in the server package
// Everything concerning requests/responses belongs in here. Only conversion (business-model <-> protobuf) should happen here actually.
type {{ .ServiceName }}APIHandler struct {
	implementation *service.{{ .GrpcServiceName }}
}

func New{{ .GrpcServiceName }}Handler(implementation *service.{{ .GrpcServiceName }}) *{{ .ServiceName }}APIHandler{
	return &{{ .ServiceName }}APIHandler{
		implementation: implementation,
	}
}

func (e *{{ .ServiceName }}APIHandler) Greeting(ctx context.Context, request *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	greeting, err := e.implementation.Hello(request.Name)
	if err != nil {
		return nil, err
	}

	return &greeter.GreetingResponse{Greeting: greeting}, nil
}

