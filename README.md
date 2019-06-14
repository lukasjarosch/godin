
<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/lukasjarosch/godin">
    <img src="gopher.png" alt="Logo" width="250" height="250">
  </a>

# Godin - An opinionated toolkit for [go-kit](https://github.com/go-kit/kit) microservices
> Work in progress, be careful.

After starting your journey into the realm of microservices with Go, you sooner or later start to get lazy and copy&paste stuff.
Things get horrbile to maintain and all the joy of programming escapes.

**Godin** to the rescue! It helps you to create, update and maintain microservices using [go-kit](https://github.com/go-kit/kit/)
as foundation.

A quick summary of what Godin can do:
* [x] initialize a new microservice project including Makefile, Dockerfile and Kubernetes deployment/service manifests
* [x] generate a running go-kit microservice application based on a single service-interface.
* [x] update all necessary files with one command `godin update` and all new functions in the service interface are prepared 
* [x] remove tedious copy&paste work when writing request/response de-/encoders
* [ ] add new bundles to quickly up the game (mysql, mongodb, pub/sub, ...)
* [ ] pre-generate tests


### Built With
* [cobra](https://github.com/spf13/cobra)
* [viper](https://github.com/spf13/viper)
* [go-kit](https://github.com/go-kit/kit/)
* [go-astra](https://github.com/vetcher/go-astra)


<!-- GETTING STARTED -->
## Getting Started
### Prerequisites
Well, Go (1.12) must be installed, and `make`.

### Installation
> There are no binary releases, yet :)
1. clone the repo and `cd` into it
2. install using `make install`
3. Done

## Usage

### Available commands
|      |      |
|------|------|
| **init** | initialize a new godin project in the current directory |
| **update** | update the project in the current directory and regenerate necessary stuff |
| **version** | print godin version |
| **help** | Help about any command |


### godin init
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

### godin update
After writing out the service interface in `internal/service/service.go`, calling `godin update` will
generate everything required to serve a gRPC server. Middleware will automatically be added based on the project configuration 
in `godin.toml`.

This is the structure `godin update` currently creates
```bash
.
├── cmd
│   └── greeter
├── Dockerfile
├── godin.toml
├── go.mod
├── go.sum
├── internal
│   ├── grpc
│   │   ├── encode_decode.go
│   │   └── request_response.go
│   └── service
│       ├── endpoint
│       │   ├── request_response.go
│       │   └── set.go
│       ├── middleware
│       │   ├── logging.go
│       │   └── middleware.go
│       ├── service.go
│       └── greeter
│           └── implementation.go
└── pkg
    └── grpc
```

Not all files will be generated fully every time `update` is executed. Below is a table to see which
files you'd better not edit :sweat_smile:.

| File                                             | Mode       |
|--------------------------------------------------|------------|
| `internal/service/endpoint/request_response.go`    | full |
| `internal/service/endpoint/set.go`                 | full |
| `internal/service/middleware/middleware.go`        | full |
| `internal/service/middleware/logging.go`           | update     |
| `internal/service/<serviceName>/implementation.go`  | update     |
| `internal/grpc/request_response.go` | full |
| `internal/grpc/encode_decode.go` | update |


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
    
[grpc]
  enabled = true
```

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.
