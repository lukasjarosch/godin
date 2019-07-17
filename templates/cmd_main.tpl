// Code generated by Godin v{{ .Godin.Version }}; DO NOT EDIT.
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
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
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

	// initialize service including middleware
	var svc service.{{ title .Service.Name }}
	svc = usecase.NewServiceImplementation(logger)
	{{- if .Service.LoggingMiddleware }}
	svc = middleware.LoggingMiddleware(logger)(svc)
	{{- end }}

	// initialize endpoint and transport layers
	var (
		endpoints   = endpoint.Endpoints(svc, logger)
		grpcHandler = svcGrpc.NewServer(endpoints, logger)
		{{- if .Service.Transport.AMQP }}
		rabbitmqConn = initRabbitMQ(logger)
		{{- end }}
	)

	{{- if .Service.Transport.AMQP }}
	// setup AMQP subscriptions
	subscriptions := amqp.Subscriptions(rabbitmqConn.Channel)
	{{- range .Service.Subscriber }}
	if err := subscriptions.{{ .Handler }}(logger, svc); err != nil {
		logger.Error("failed to create subscription", "err", err)
		os.Exit(-1)
	}
	{{- end }}
	{{- end }}

	// serve gRPC server
	grpcServer := googleGrpc.NewServer(
	    googleGrpc.UnaryInterceptor(grpcPrometheus.UnaryServerInterceptor),
	)
	g.Add(initGrpc(grpcServer, grpcHandler, logger), func(error) {
		grpcServer.GracefulStop()
	})
	grpcPrometheus.Register(grpcServer)

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

// getEnv get key environment variable if exist otherwise return defalutValue
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// shutdownHandler to handle graceful shutdowns
func shutdownHandler(interruptChannel chan struct{}) func() error {
	return func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-interruptChannel:
			return nil
		}
	}
}

// initGrpc serve up GRPC
func initGrpc(grpcServer *googleGrpc.Server, handler pb.{{ .Protobuf.Service }}Server, logger log.Logger) func() error {
	grpcListener, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}

	return func() error {
		logger.Log("transport", "gRPC", "addr", GrpcAddr)
		pb.Register{{ .Protobuf.Service }}Server(grpcServer, handler)
		return grpcServer.Serve(grpcListener)
	}
}

{{ if .Service.Transport.AMQP }}
// initRabbitMQ will initialize the amqp connection and create a new channel
func initRabbitMQ(logger log.Logger) *rabbitmq.RabbitMQ {
	rabbitmqConn, err := rabbitmq.NewRabbitMQFromEnv()
	if err != nil {
		logger.Error("failed to initialize RabbitMQ connection", "err", err)
		os.Exit(-1)
	}
	if err := rabbitmqConn.Connect(); err != nil {
		logger.Error("failed to connect to RabbitMQ", "err", err)
		os.Exit(-1)
	}
	if err := rabbitmqConn.NewChannel(); err != nil {
		logger.Error("failed to create AMQP channel", "err", err)
		os.Exit(-1)
	}

	return rabbitmqConn
}
{{ end }}

func initDebugHttp(logger log.Logger) net.Listener {
	debugListener, err := net.Listen("tcp", DebugAddr)
	if err != nil {
		logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
		os.Exit(1)
	}
	return debugListener
}

