// Code generated by Godin v0.3.0; DO NOT EDIT.

package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	godinMiddleware "github.com/lukasjarosch/godin/pkg/middleware"
	service "yyy"
)

type Set struct {
	HelloEndpoint  endpoint.Endpoint
	Hello2Endpoint endpoint.Endpoint
	Hello3Endpoint endpoint.Endpoint
	Hello4Endpoint endpoint.Endpoint
	Hello5Endpoint endpoint.Endpoint
	Hello6Endpoint endpoint.Endpoint
	Hello7Endpoint endpoint.Endpoint
	Hello8Endpoint endpoint.Endpoint
	Hello9Endpoint endpoint.Endpoint
}

func Endpoints(service service.yyy, duration metrics.Histogram, frequency metrics.Counter) Set {

	var hello endpoint.Endpoint
	{
		hello = HelloEndpoint(service)
		hello = godinMiddleware.LatencyMiddleware(duration, "Hello")(hello)
		hello = godinMiddleware.RequestFrequency(frequency, "Hello")(hello)
	}
	var hello2 endpoint.Endpoint
	{
		hello2 = Hello2Endpoint(service)
		hello2 = godinMiddleware.LatencyMiddleware(duration, "Hello2")(hello2)
		hello2 = godinMiddleware.RequestFrequency(frequency, "Hello2")(hello2)
	}
	var hello3 endpoint.Endpoint
	{
		hello3 = Hello3Endpoint(service)
		hello3 = godinMiddleware.LatencyMiddleware(duration, "Hello3")(hello3)
		hello3 = godinMiddleware.RequestFrequency(frequency, "Hello3")(hello3)
	}
	var hello4 endpoint.Endpoint
	{
		hello4 = Hello4Endpoint(service)
		hello4 = godinMiddleware.LatencyMiddleware(duration, "Hello4")(hello4)
		hello4 = godinMiddleware.RequestFrequency(frequency, "Hello4")(hello4)
	}
	var hello5 endpoint.Endpoint
	{
		hello5 = Hello5Endpoint(service)
		hello5 = godinMiddleware.LatencyMiddleware(duration, "Hello5")(hello5)
		hello5 = godinMiddleware.RequestFrequency(frequency, "Hello5")(hello5)
	}
	var hello6 endpoint.Endpoint
	{
		hello6 = Hello6Endpoint(service)
		hello6 = godinMiddleware.LatencyMiddleware(duration, "Hello6")(hello6)
		hello6 = godinMiddleware.RequestFrequency(frequency, "Hello6")(hello6)
	}
	var hello7 endpoint.Endpoint
	{
		hello7 = Hello7Endpoint(service)
		hello7 = godinMiddleware.LatencyMiddleware(duration, "Hello7")(hello7)
		hello7 = godinMiddleware.RequestFrequency(frequency, "Hello7")(hello7)
	}
	var hello8 endpoint.Endpoint
	{
		hello8 = Hello8Endpoint(service)
		hello8 = godinMiddleware.LatencyMiddleware(duration, "Hello8")(hello8)
		hello8 = godinMiddleware.RequestFrequency(frequency, "Hello8")(hello8)
	}
	var hello9 endpoint.Endpoint
	{
		hello9 = Hello9Endpoint(service)
		hello9 = godinMiddleware.LatencyMiddleware(duration, "Hello9")(hello9)
		hello9 = godinMiddleware.RequestFrequency(frequency, "Hello9")(hello9)
	}

	return Set{
		HelloEndpoint:  helloEndpoint,
		Hello2Endpoint: hello2Endpoint,
		Hello3Endpoint: hello3Endpoint,
		Hello4Endpoint: hello4Endpoint,
		Hello5Endpoint: hello5Endpoint,
		Hello6Endpoint: hello6Endpoint,
		Hello7Endpoint: hello7Endpoint,
		Hello8Endpoint: hello8Endpoint,
		Hello9Endpoint: hello9Endpoint,
	}
}