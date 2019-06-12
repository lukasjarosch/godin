package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"yyy/internal/service"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.Yyy
}

// Hello logs the request and response of the service.Hello endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l logMiddleware) Hello(ctx context.Context, name string) (greeting *service.Greeting, err error) {
	l.logger.Log(
		"endpoint", "Hello",
		"request", endpoint.HelloRequest{ctx, name},
	)

	defer func(begin time.Time) {
		resp := endpoint.HelloResponse{greeting}
		if err != nil {
			resp.Err = err
		}

		i.logger.Log(
			"endpoint", "Hello",
			"response", resp,
			"took", time.Since(begin),
		)
	}(time.Now())

	return i.next.Hello(ctx, name)
}

// Hello2 logs the request and response of the service.Hello2 endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l logMiddleware) Hello2(ctx context.Context, name string) (greeting service.Greeting, err error) {
	l.logger.Log(
		"endpoint", "Hello2",
		"request", endpoint.Hello2Request{ctx, name},
	)

	defer func(begin time.Time) {
		resp := endpoint.Hello2Response{greeting}
		if err != nil {
			resp.Err = err
		}

		i.logger.Log(
			"endpoint", "Hello2",
			"response", resp,
			"took", time.Since(begin),
		)
	}(time.Now())

	return i.next.Hello2(ctx, name)
}

// Hello3 logs the request and response of the service.Hello3 endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l logMiddleware) Hello3(ctx context.Context, name string) (greeting string, err error) {
	l.logger.Log(
		"endpoint", "Hello3",
		"request", endpoint.Hello3Request{ctx, name},
	)

	defer func(begin time.Time) {
		resp := endpoint.Hello3Response{greeting}
		if err != nil {
			resp.Err = err
		}

		i.logger.Log(
			"endpoint", "Hello3",
			"response", resp,
			"took", time.Since(begin),
		)
	}(time.Now())

	return i.next.Hello3(ctx, name)
}

// Hello4 logs the request and response of the service.Hello4 endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l logMiddleware) Hello4(ctx context.Context, name []string) (greeting []service.Greeting, err error) {
	l.logger.Log(
		"endpoint", "Hello4",
		"request", endpoint.Hello4Request{ctx, name},
	)

	defer func(begin time.Time) {
		resp := endpoint.Hello4Response{greeting}
		if err != nil {
			resp.Err = err
		}

		i.logger.Log(
			"endpoint", "Hello4",
			"response", resp,
			"took", time.Since(begin),
		)
	}(time.Now())

	return i.next.Hello4(ctx, name)
}

// Hello5 logs the request and response of the service.Hello5 endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l logMiddleware) Hello5(ctx context.Context, name []string) (greeting []*service.Greeting, err error) {
	l.logger.Log(
		"endpoint", "Hello5",
		"request", endpoint.Hello5Request{ctx, name},
	)

	defer func(begin time.Time) {
		resp := endpoint.Hello5Response{greeting}
		if err != nil {
			resp.Err = err
		}

		i.logger.Log(
			"endpoint", "Hello5",
			"response", resp,
			"took", time.Since(begin),
		)
	}(time.Now())

	return i.next.Hello5(ctx, name)
}

// Hello6 logs the request and response of the service.Hello6 endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l logMiddleware) Hello6(ctx context.Context, name []*service.Greeting) (greeting []*service.Greeting, err error) {
	l.logger.Log(
		"endpoint", "Hello6",
		"request", endpoint.Hello6Request{ctx, name},
	)

	defer func(begin time.Time) {
		resp := endpoint.Hello6Response{greeting}
		if err != nil {
			resp.Err = err
		}

		i.logger.Log(
			"endpoint", "Hello6",
			"response", resp,
			"took", time.Since(begin),
		)
	}(time.Now())

	return i.next.Hello6(ctx, name)
}

// Hello7 logs the request and response of the service.Hello7 endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l logMiddleware) Hello7(ctx context.Context, name *service.Greeting) (greeting []string, err error) {
	l.logger.Log(
		"endpoint", "Hello7",
		"request", endpoint.Hello7Request{ctx, name},
	)

	defer func(begin time.Time) {
		resp := endpoint.Hello7Response{greeting}
		if err != nil {
			resp.Err = err
		}

		i.logger.Log(
			"endpoint", "Hello7",
			"response", resp,
			"took", time.Since(begin),
		)
	}(time.Now())

	return i.next.Hello7(ctx, name)
}

// Hello8 logs the request and response of the service.Hello8 endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l logMiddleware) Hello8(ctx context.Context, name *[]service.Greeting) (greeting []string, err error) {
	l.logger.Log(
		"endpoint", "Hello8",
		"request", endpoint.Hello8Request{ctx, name},
	)

	defer func(begin time.Time) {
		resp := endpoint.Hello8Response{greeting}
		if err != nil {
			resp.Err = err
		}

		i.logger.Log(
			"endpoint", "Hello8",
			"response", resp,
			"took", time.Since(begin),
		)
	}(time.Now())

	return i.next.Hello8(ctx, name)
}

// Hello9 logs the request and response of the service.Hello9 endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l logMiddleware) Hello9(ctx context.Context, name *[]service.Greeting, foo string, bar string) (greeting []string, err error) {
	l.logger.Log(
		"endpoint", "Hello9",
		"request", endpoint.Hello9Request{ctx, name, foo, bar},
	)

	defer func(begin time.Time) {
		resp := endpoint.Hello9Response{greeting}
		if err != nil {
			resp.Err = err
		}

		i.logger.Log(
			"endpoint", "Hello9",
			"response", resp,
			"took", time.Since(begin),
		)
	}(time.Now())

	return i.next.Hello9(ctx, name, foo, bar)
}