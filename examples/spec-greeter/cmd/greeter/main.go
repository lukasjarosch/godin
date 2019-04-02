package main

import (
	"os"

	"os/signal"
	"syscall"

	cfg "github.com/lukasjarosch/godin/examples/spec-greeter/internal/config"
	svc "github.com/lukasjarosch/godin/examples/spec-greeter/internal/server"
	"github.com/lukasjarosch/godin/pkg/grpc"
	"github.com/sirupsen/logrus"
)

// Compile time variables are injected
var (
	GitCommit string
	GitBranch string
	BuildTime string
)

func main() {
	// perpare dependencies
	config := cfg.NewConfig()
	logger := initLogging(config.LogDebug)

	// setup our gRPC server
	server := svc.NewServer(config, logger)
	logger.WithFields(logrus.Fields{
		"instance":   server.GRPC.Options.ID,
		"git.commit": GitCommit,
		"git.branch": GitBranch,
		"build":      BuildTime,
	}).Infof("starting server: %s", server.GRPC.Options.Name)

	// graceful shutdown using signals (SIGINT and SIGTERM)
	// the GRPC shutdown handler will also take care of the prometheus HTTP server
	go shutdownHandler(server.GRPC)

	// HTTP server providing Prometheus metrics is included in the gRPC server
	server.GRPC.ServeMetrics()

	// finally: serve the gRPC server in the foreground
	if err := server.GRPC.Serve(); err != nil {
		logger.Fatal(err)
	}
}

// initLogging initializes a new zap productionLogger and returns the sugared logger
func initLogging(debug bool) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	if debug {
		logger.SetLevel(logrus.DebugLevel)
	}

	return logger
}

// wait for SIGINT or SIGTERM and then call Shutdown()
func shutdownHandler(service *grpc.Server) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT)
	signal.Notify(sigs, syscall.SIGTERM)
	logrus.Infof("signal: %v", <-sigs)

	service.Shutdown()
}
