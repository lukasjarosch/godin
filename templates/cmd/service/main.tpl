package main

import (
	"os"

	"github.com/go-kit/kit/log"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewJSONLogger(os.Stderr)
		logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)
	}

	logger.Log("fatal", "this godin project does not yet provide any functionality")
	os.Exit(1)

    /*
	var svc service.Example
	svc = example.NewExampleService(logger)
	svc = middleware.NewLogMiddleware(logger)(svc)
	*/

	// endpoints := endpoint.Endpoints(svc)

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