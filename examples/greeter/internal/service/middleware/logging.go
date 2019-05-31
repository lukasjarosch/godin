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

func (l logMiddleware) Hello(ctx context.Context, name string) (greeting *service.Greeting, err error) {

	// log the request of Hello
	l.logger.Log(
		"endpoint", "Hello",
		"request", endpoint.HelloRequest{ctx, name},
	)

	// log the response of Hello
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

func (l logMiddleware) Hello2(ctx context.Context, name string) (greeting service.Greeting, err error) {

	// log the request of Hello2
	l.logger.Log(
		"endpoint", "Hello2",
		"request", endpoint.Hello2Request{ctx, name},
	)

	// log the response of Hello2
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

func (l logMiddleware) Hello3(ctx context.Context, name string) (greeting string, err error) {

	// log the request of Hello3
	l.logger.Log(
		"endpoint", "Hello3",
		"request", endpoint.Hello3Request{ctx, name},
	)

	// log the response of Hello3
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

func (l logMiddleware) Hello4(ctx context.Context, name []string) (greeting []service.Greeting, err error) {

	// log the request of Hello4
	l.logger.Log(
		"endpoint", "Hello4",
		"request", endpoint.Hello4Request{ctx, name},
	)

	// log the response of Hello4
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

func (l logMiddleware) Hello5(ctx context.Context, name []string) (greeting []*service.Greeting, err error) {

	// log the request of Hello5
	l.logger.Log(
		"endpoint", "Hello5",
		"request", endpoint.Hello5Request{ctx, name},
	)

	// log the response of Hello5
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

func (l logMiddleware) Hello6(ctx context.Context, name []*service.Greeting) (greeting []*service.Greeting, err error) {

	// log the request of Hello6
	l.logger.Log(
		"endpoint", "Hello6",
		"request", endpoint.Hello6Request{ctx, name},
	)

	// log the response of Hello6
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

func (l logMiddleware) Hello7(ctx context.Context, name *service.Greeting) (greeting []string, err error) {

	// log the request of Hello7
	l.logger.Log(
		"endpoint", "Hello7",
		"request", endpoint.Hello7Request{ctx, name},
	)

	// log the response of Hello7
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

func (l logMiddleware) Hello8(ctx context.Context, name *[]service.Greeting) (greeting []string, err error) {

	// log the request of Hello8
	l.logger.Log(
		"endpoint", "Hello8",
		"request", endpoint.Hello8Request{ctx, name},
	)

	// log the response of Hello8
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

func (l logMiddleware) Hello9(ctx context.Context, name *[]service.Greeting, foo string, bar string) (greeting []string, err error) {

	// log the request of Hello9
	l.logger.Log(
		"endpoint", "Hello9",
		"request", endpoint.Hello9Request{ctx, name, foo, bar},
	)

	// log the response of Hello9
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
