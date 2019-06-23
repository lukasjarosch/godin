package middleware

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
)

func RequestFrequency(frequency metrics.Counter, methodName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func() {
				frequency.With("success", fmt.Sprint(err == nil), "method", methodName).Add(1)
			}()
			return next(ctx, request)
		}
	}
}
