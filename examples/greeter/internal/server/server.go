package server

import (
	greeter "github.com/lukasjarosch/godin-api-go/godin/greeter/v1beta1"
	"github.com/lukasjarosch/godin/examples/greeter/internal/config"
	service "github.com/lukasjarosch/godin/examples/greeter/internal/greeter"
	godin "github.com/lukasjarosch/godin/pkg/grpc"
	"github.com/lukasjarosch/godin/pkg/http"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// server is a wrapper to hold all of our services stuff.
// Anything related to the transport-layer can be set up here
type server struct {
	GRPC        *godin.Server
	HTTPGateway *http.Server
}

// NewServer constructs a new Server using your service as gRPC handler implementation.
func NewServer(config *config.Config, logger *logrus.Logger) *server {

	// setup the business logic with it's dependencies
	svc := service.NewGreeterAPI(config, logger)

	// wrap our business logic in the transport handler
	handler := NewGreeterAPIHandler(svc)

	// attach our business logic to the gRPC server
	impl := func(g *grpc.Server) {
		// TODO: register the server using your protobuf stub
		greeter.RegisterGreeterAPIServer(g, handler)
	}

	// create the actual gRPC server
	// See pkg/server/server.go for the default options
	server := godin.NewServer(
		godin.Name("greeter"),
		godin.Implementation(impl),

		// Override config with env variables configured by our business domain
		godin.GrpcNetworkPort(config.GrpcPort),
	)

	return &exampleServiceServer{
		GRPC:        server,
		HTTPGateway: gatewayServer,
	}
}
