// Code generated by Godin v{{ .Godin.Version }}; DO NOT EDIT.
package main

import (
    "fmt"
    "net"
    "os"
    "os/signal"
    "syscall"

	googleGrpc "google.golang.org/grpc"
	"github.com/go-godin/rabbitmq"

    pb "{{ .Protobuf.Package }}"
	"{{ .Service.Module }}/internal/amqp"
    "{{ .Service.Module }}/internal/service"

	"github.com/go-godin/log"
)

// getEnv get key environment variable if exist otherwise return defaultValue
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
	if rabbitmqConn.Channel, err = rabbitmqConn.NewChannel(); err != nil {
		logger.Error("failed to create AMQP channel", "err", err)
		os.Exit(-1)
	}

	return rabbitmqConn
}
{{ end }}

{{ if gt (len .Service.Subscriber) 0 }}
func initSubscriptions(logger log.Logger, svc service.{{ title .Service.Name}}, connection *rabbitmq.RabbitMQ) {
	subscriptions := amqp.Subscriptions(connection)
	{{- range .Service.Subscriber }}
	if err := subscriptions.{{ .Handler }}(logger, svc); err != nil {
		logger.Error("failed to create subscription", "err", err)
		os.Exit(-1)
	}
	{{- end }}
}
{{- end }}

func initDebugHttp(logger log.Logger) net.Listener {
	debugListener, err := net.Listen("tcp", DebugAddr)
	if err != nil {
		logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
		os.Exit(1)
	}
	return debugListener
}

