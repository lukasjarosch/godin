// Code generated by Godin v{{ .Godin.Version }}.
package main

import (
    "fmt"
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"

	"github.com/oklog/oklog/pkg/group"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	googleGrpc "google.golang.org/grpc"
	"github.com/go-godin/grpc-interceptor"
	"github.com/go-godin/rabbitmq"

    pb "{{ .Protobuf.Package }}"
	"{{ .Service.Module }}/internal/amqp"
    svcGrpc "{{ .Service.Module }}/internal/grpc"
    "{{ .Service.Module }}/internal/service"
    "{{ .Service.Module }}/internal/service/usecase"
    "{{ .Service.Module }}/internal/service/middleware"
    "{{ .Service.Module }}/internal/service/endpoint"

	"github.com/go-godin/log"
)

var DebugAddr = getEnv("DEBUG_ADDRESS", "0.0.0.0:3000")
var GrpcAddr = getEnv("GRPC_ADDRESS", "0.0.0.0:50051")

// group to manage the lifecycle for goroutines
var g group.Group

func main() {
    logger := log.NewLoggerFromEnv()

    // setup AMQP connections (separate publish and subscribe connections to reduce AMQP pressure)
	{{- if gt (len .Service.Subscriber) 0 }}
	rabbitmqSubConn := initRabbitMQ(logger)
	defer rabbitmqSubConn.Close()
	{{- end }}
	{{- if gt (len .Service.Publisher) 0 }}
	rabbitmqPubConn := initRabbitMQ(logger)
	defer rabbitmqPubConn.Close()
	{{- end }}
	{{ if gt (len .Service.Publisher) 0 }}
	// init publishers
	publishers := amqp.Publishers(rabbitmqPubConn, logger)
	{{- end }}

	// initialize service including middleware
	var svc service.{{ title .Service.Name }}
	{{/*
		Depending on which bundles are activated, the parameter list of the service implementation is going to change.
		This is a tricky bit as all combinations need to be taken care of.
	*/}}
	{{- if gt (len .Service.Publisher) 0 }}
	svc = usecase.NewServiceImplementation(logger, publishers)
	{{- else }}
	svc = usecase.NewServiceImplementation(logger)
	{{- end }}

	{{- if .Service.LoggingMiddleware }}
	svc = middleware.LoggingMiddleware(logger)(svc)
	{{- end }}

	// initialize endpoint and transport layers
	var (
		endpoints   = endpoint.Endpoints(svc, logger)
		grpcHandler = svcGrpc.NewServer(endpoints, logger)
	)

	// serve gRPC server
	grpcServer := googleGrpc.NewServer(
	    googleGrpc.UnaryInterceptor(grpc_interceptor.UnaryInterceptor),
	)
	g.Add(initGrpc(grpcServer, grpcHandler, logger), func(error) {
		grpcServer.GracefulStop()
	})

	// serve debug http server (prometheus)
	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener := initDebugHttp(logger)
	g.Add(func() error {
		logger.Log("transport", "debug/HTTP", "addr", DebugAddr)
		return http.Serve(debugListener, http.DefaultServeMux)
	}, func(error) {
		debugListener.Close()
	})

	// Wait for SIGINT or SIGTERM and stop gracefully
	cancelInterrupt := make(chan struct{})
	g.Add(shutdownHandler(cancelInterrupt), func(e error) {
		close(cancelInterrupt)
	})

	// run
	if err := g.Run(); err != nil {
		logger.Log("fatal", err)
		os.Exit(1)
	}
}

