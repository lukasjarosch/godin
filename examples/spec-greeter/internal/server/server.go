package server

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/lukasjarosch/godin-api-go/godin/greeter/v1beta1"
	service "github.com/lukasjarosch/godin/examples/spec-greeter/internal/greeter"
	godin "github.com/lukasjarosch/godin/pkg/grpc"

	"github.com/lukasjarosch/godin/examples/spec-greeter/internal/config"
)

// server is a wrapper to hold all of our services stuff.
// Anything related to the transport-layer can be set up here
type GreeterServer struct {
	GRPC *godin.Server
}

// NewServer constructs a new Server using your service as gRPC handler implementation.
func NewServer(config *config.Config, logger *logrus.Logger) *GreeterServer {

	// setup the business logic the dependencies
	svc := service.NewGreeter(config, logger)

	// wrap the business logic inside the transport handler
	handler := NewGreeterAPIHandler(svc)

	// attach the handler to the gRPC server
	impl := func(g *grpc.Server) {
		pb.RegisterGreeterAPIServer(g, handler)
	}

	// create the actual gRPC server
	// See pkg/server/server.go for the default options
	server := godin.NewServer(
		godin.Name("greeter"),
		godin.Implementation(impl),

		// Override config with env variables configured by our business domain
		godin.GrpcNetworkPort(config.GrpcPort),
	)

	return &GreeterServer{
		GRPC: server,
	}
}
