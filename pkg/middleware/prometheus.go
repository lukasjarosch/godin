package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"fmt"
)

var requestDuration = prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
	Name:    "grpc_request_duration_ms",
	Help:    "Request duration in milliseconds",
	Buckets: []float64{50, 100, 250, 500, 1000},
}, []string{"method"})

var requestsCurrent = prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
	Name: "grpc_requests_current",
	Help: "The current numer of gRPC requests by endpoint",
}, []string{"method"})

var requestStatus = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
	Name: "grpc_requests_total",
	Help: "The total number of gRPC requests and whether the business failed or not",
}, []string{"method", "success"})

// Prometheus adds basic RED metrics on all endpoints. The transport layer (gRPC) should also have metrics attached and
// will then take care of monitoring grpc endpoints including their status.
func Prometheus(methodName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			requestsCurrent.With("method", methodName).Add(1)

			defer func(begin time.Time) {
				requestDuration.With("method", methodName).Observe(time.Since(begin).Seconds())
				requestsCurrent.With("method", methodName).Add(-1)
				requestStatus.With("method", methodName, "success", fmt.Sprint(err == nil))
			}(time.Now())

			return next(ctx, request)
		}
	}
}
