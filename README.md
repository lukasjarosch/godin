<p align="center">
  <img src="https://raw.githubusercontent.com/lukasjarosch/godin/develop/gopher.png">
</p>

# Godin - An opinionated toolkit for [go-kit](https://github.com/go-kit/kit) microservices
> Work in progress

## godin update
| File                                             | Mode       |
|--------------------------------------------------|------------|
| `internal/service/endpoint/request_response.go`    | full |
| `internal/service/endpoint/set.go`                 | full |
| `internal/service/middleware/middleware.go`        | full |
| `internal/service/middleware/logging.go`           | update     |
| `internal/service/<serviceName>/implementation.go`  | update     |

### Protobufs
Protocol buffers are used to describe the APIs of an application.
The API definitions are the heart of a microservice architecture, so a workflow
to organize and automate protobufs should already exist.

I have a basic workflow running here: [lukasjarosch/godin-protobuf](https://github.com/lukasjarosch/godin-protobuf)
