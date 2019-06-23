package middleware

import (
	"context"
	"time"

	"github.com/lukasjarosch/godin/pkg/log"

	"github.com/lukasjarosch/godin/examples/user/internal/service"
	"github.com/lukasjarosch/godin/examples/user/internal/service/endpoint"
)

type loggingMiddleware struct {
	next   service.User
	logger log.Logger
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.User) service.User {
		return &loggingMiddleware{next, logger}
	}
}


// Create logs the request and response of the service.Create endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l loggingMiddleware) Create(ctx context.Context, username string, email string) (user *service.UserEntity, err error) {
	l.logger.Log(
		"endpoint", "Create",
		"request", endpoint.CreateRequest{
			Username: username,
			Email:    email,
		},
	)

	defer func(begin time.Time) {
		resp := endpoint.CreateResponse{User: user}

		l.logger.Log(
			"endpoint", "Create",
			"response", resp,
			"error", err,
			"success", err == nil,
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.Create(ctx, username, email)
}
