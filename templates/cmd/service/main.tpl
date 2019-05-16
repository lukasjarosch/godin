package main

import (
	"os"

	"github.com/go-kit/kit/log"
	"{{ .Service.Name }}/internal"
	"{{ .Service.Name }}/internal/service/{{ .Service.Name }}"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewJSONLogger(os.Stderr)
		logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)
	}

    // setup service implementation
	var svc service.{{ title .Service.Name }}
	svc = {{ .Service.Name }}.NewImplementation(logger)

	endpoints := endpoint.Endpoints(svc)

	// grpcHandler := grpc.NewGrpcServer(endpoints, logger)

    /*
	grpcServer := googleGrpc.NewServer()
	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}
	*/

    /*
	logger.Log("transport", "gRPC", "addr", ":50051")
	api.RegisterExampleServiceServer(grpcServer, grpcHandler)
	if err := grpcServer.Serve(grpcListener); err != nil {
		panic(err)
	}
	*/
}