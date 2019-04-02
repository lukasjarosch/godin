package server

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

    pb "{{ .Spec.Service.API.Import }}"
	service "{{ .ModuleName }}/internal/{{ .ServiceName }}"
	godin "github.com/lukasjarosch/godin/pkg/grpc"

	"{{ .ModuleName }}/internal/config"
)

// server is a wrapper to hold all of our services stuff.
// Anything related to the transport-layer can be set up here
type {{ .ServiceName | camelcase }}Server struct {
	GRPC        *godin.Server
}

// NewServer constructs a new Server using your service as gRPC handler implementation.
func NewServer({{ deps_param_list }}) *{{ .ServiceName | camelcase }}Server {

	// setup the business logic the dependencies
	svc := service.New{{ .ServiceName | camelcase }}({{ deps_name_list }})

	// wrap the business logic inside the transport handler
	handler := New{{ .GrpcServiceName }}Handler(svc)

	// attach the handler to the gRPC server
	impl := func(g *grpc.Server) {
		pb.Register{{ .GrpcServiceName }}Server(g, handler)
	}

	// create the actual gRPC server
	// See pkg/server/server.go for the default options
	server := godin.NewServer(
		godin.Name("{{ .ServiceName }}"),
		godin.Implementation(impl),

		// Override config with env variables configured by our business domain
		godin.GrpcNetworkPort(config.GrpcPort),
	)

	return &{{ .ServiceName | camelcase }}Server{
		GRPC:        server,
	}
}