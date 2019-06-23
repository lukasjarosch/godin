// Code generated by Godin v{{ .Godin.Version }}; DO NOT EDIT.
package main

import (
    "fmt"
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/oklog/oklog/pkg/group"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	googleGrpc "google.golang.org/grpc"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

    pb "{{ .Protobuf.Package }}"
    svcGrpc "{{ .Service.Module }}/internal/grpc"
    "{{ .Service.Module }}/internal/service"
    "{{ .Service.Module }}/internal/service/{{ .Service.Name }}"
    "{{ .Service.Module }}/internal/service/middleware"
    "{{ .Service.Module }}/internal/service/endpoint"

	"github.com/lukasjarosch/godin/pkg/log"
)

var DebugAddr = getEnv("DEBUG_ADDRESS", "0.0.0.0:3000")
var GrpcAddr = getEnv("GRPC_ADDRESS", "0.0.0.0:50051")

// group to manage the lifecycle for goroutines
var g group.Group

func main() {
    logger := log.New()

	// initialize service layer
	var svc service.{{ title .Service.Name }}
	svc = {{ .Service.Name }}.NewServiceImplementation(logger)
	{{- if .Service.LoggingMiddleware }}
	svc = middleware.LoggingMiddleware(logger)(svc)
	{{- end }}
	//TODO: svc = middleware.AuthorizationMiddleware(logger)(svc)
	//TODO: svc = middleware.RecoveringMiddleware(logger)(svc)

	// initialize prometheus
	requestDuration, requestFrequency := initMetrics()
	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())

	// initialize endpoint and transport layers
	var (
		endpoints   = endpoint.Endpoints(svc, requestDuration, requestFrequency)
		grpcHandler = svcGrpc.NewServer(endpoints, logger)
	)

	// serve debug http server (prometheus)
	debugListener := initDebugHttp(logger)
	g.Add(func() error {
		logger.Log("transport", "debug/HTTP", "addr", DebugAddr)
		return http.Serve(debugListener, http.DefaultServeMux)
	}, func(error) {
		debugListener.Close()
	})

	// serve gRPC server
	grpcServer := googleGrpc.NewServer()
	g.Add(initGrpc(grpcServer, grpcHandler, logger), func(error) {
		grpcServer.GracefulStop()
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

func initDebugHttp(logger log.Logger) net.Listener {
	debugListener, err := net.Listen("tcp", DebugAddr)
	if err != nil {
		logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
		os.Exit(1)
	}
	return debugListener
}

// initMetrics initializes all prometheus metrics for the transport layer
//
// contact_request_duration_seconds Summary/Histogram
// contact_request_count_total Counter
func initMetrics() (duration metrics.Histogram, frequency metrics.Counter) {
	{
		duration = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "{{ .Service.Namespace }}",
			Subsystem: "{{ .Service.Namespace }}",
			Name:      "request_duration_seconds",
			Help:      "Request duration in seconds.",
		}, []string{"method", "success"})
	}
	{
		frequency = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "{{ .Service.Namespace }}",
			Subsystem: "{{ .Service.Namespace }}",
			Name:      "request_count_total",
			Help:      "the total amount of requests served",
		}, []string{"method", "success"})
	}

	return duration, frequency
}
