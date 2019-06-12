<p align="center">
  <img src="https://raw.githubusercontent.com/lukasjarosch/godin/develop/gopher.png">
</p>

# Godin - An opinionated toolkit for [go-kit](https://github.com/go-kit/kit) microservices
> Work in progress

## godin init
Calling `godin init` will generate the following project structure in the CWD. First you need to answer
a few prompts to configure the project. You only have to do this once, after that the configuration is stored
in `godin.toml`

````bash
.
├── cmd
│   └── greeter
├── Dockerfile
├── godin.toml
├── go.mod
├── internal
│   ├── grpc
│   └── service
│       ├── endpoint
│       ├── greeter
│       ├── middleware
│       │   └── middleware.go
│       └── service.go
└── pkg
    └── grpc
````

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


### configuration
After `godin init` is called, everything is based off your service definition plus the `godin.toml`. 
An example of such a config file is listed below.

```toml
[godin]
  version = "0.3.0"

[project]
  created = "Wed, 12 Jun 2019 19:13:15 CEST"

[protobuf]
  package = "godin.greeter"
  service = "GreeterService"

[service]
  module = "github.com/lukasjarosch/greeter"
  name = "greeter"
  namespace = "godin"

  [service.middleware]
    authorization = false
    caching = false
    logging = true
    recovery = true

```
