
<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://raw.githubusercontent.com/lukasjarosch/godin/develop/gopher.png">
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
* Go (1.12)
* make
* [packr v1](https://github.com/gobuffalo/packr)(!)

### Installation
> There are no binary releases, yet :)
1. clone the repo and `cd` into it
2. install using `make install`
3. Done

### Examples
Examples are now hosted in a separate repository: [godin-examples](https://github.com/lukasjarosch/godin-examples)

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
in `godin.json`

````bash
.
├── cmd
│   └── greeter
├── Dockerfile
├── godin.json
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
in `godin.json`.

This is the structure `godin update` currently creates
```bash
.
├── cmd
│   └── user
│       └── main.go
├── Dockerfile
├── godin.json
├── go.mod
├── go.sum
├── internal
│   ├── grpc
│   │   ├── encode_decode.go
│   │   ├── request_response.go
│   │   └── server.go
│   └── service
│       ├── endpoint
│       │   ├── endpoints.go
│       │   ├── request_response.go
│       │   └── set.go
│       ├── middleware
│       │   ├── logging.go
│       │   └── middleware.go
│       ├── service.go
│       └── user
│           └── implementation.go
└── pkg
    └── grpc
```

Not all files will be generated fully every time `update` is executed. Below is a table to see which
files you'd better not edit :sweat_smile:.

| File                                             | Mode       |
|--------------------------------------------------|------------|
| `internal/service/middleware/middleware.go`        | full |
| `internal/service/middleware/logging.go`           | update     |
| `internal/service/endpoint/endpoints.go`           | update     |
| `internal/service/endpoint/request_response.go`           | full     |
| `internal/service/endpoint/set.go`           | full     |
| `internal/service/<serviceName>/implementation.go`  | update     |
| `internal/grpc/request_response.go` | full |
| `internal/grpc/server.go` | full |
| `internal/grpc/encode_decode.go` | update |


### configuration
After `godin init` is called, everything is based off your service definition plus the `godin.json`. 
An example of such a config file is listed below.

```json
{
  "godin": {
    "version": "0.3.0"
  },
  "project": {
    "created": "Sun, 23 Jun 2019 10:15:02 CEST",
    "updated": "Sun, 23 Jun 2019 12:14:26 CEST"
  },
  "protobuf": {
    "package": "github.com/lukasjarosch/godin-examples/user/api",
    "service": "UserService"
  },
  "service": {
    "endpoints": {
      "create": {
        "protobuf": {
          "request": "CreateRequest",
          "response": "CreateResponse"
        }
      },
      "delete": {
        "protobuf": {
          "request": "DeleteRequest",
          "response": "DeleteResponse"
        }
      },
      "get": {
        "protobuf": {
          "request": "GetRequest",
          "response": "GetResponse"
        }
      },
      "list": {
        "protobuf": {
          "request": "ListRequest",
          "response": "ListResponse"
        }
      }
    },
    "middleware": {
      "authorization": false,
      "caching": false,
      "logging": true,
      "recovery": true
    },
    "module": "github.com/lukasjarosch/godin-examples/user",
    "name": "user",
    "namespace": "godin"
  },
  "transport": {
    "grpc": {
      "enabled": true
    }
  }
}
```

### Middleware
There are three middleware levels: **service**, **endpoint** and **transport**

#### Service middleware
It's implementation specific and allows to define a middleware on domain-level.
Here you have full access to request and response data of the business logic (implementation.go).
By default, Godin will generate and maintain a logging middleware which logs the request and
response data, errors and the execution time of your business logic.

Other use-cases of service middlewares are: 
 * caching
 * authentication
 * monitoring of business metrics (users_logged_in, queue_size, ...)
 
#### Endpoint middleware
It's only provided by godin and is applicable to every endpoint.
The middleware does only have `interface{}` access to the request and response data. 
This makes the endpoint middleware a great place to put all the annoying stuff which most of 
the developers should not have to deal with, like:

* Endpoint instrumentation (prometheus)
* Circuit breaking
* Distributed tracing
* Rate limiting
* General logging (endpoint request duration)

Currently, Godin does not provide all of those middlewares. The following are implemented:
* [x] Prometheus endpoint instrumentation (exposed via `0.0.0.0:3000/metrics`). 
  - **endpoint_request_duration_ms**
  - **endpoint_requests_current**
  - **endpoint_requests_total**
* [x] General logging to capture endpoint request duration
* [ ] Circuit breaking
* [ ] Rate limiting
* [ ] Distributed tracing

#### Transport middleware
It's (obviously) transport specific. Godin is currently only providing a gRPC server
as that's it's main use-case. In the future, HTTP might follow.

Godin automatically registers the [go-grpc-prometheus](https://github.com/grpc-ecosystem/go-grpc-prometheus) interceptors.
Thus these metrics are also registered and available via `/metrics`.

#### Rename protobuf request/responses
By default, Godin will construct Protobuf requests and responses like this: `<EndpointName>Request` and `<EndpointName>Response`.
This might not always be correct as Godin does not parse the protobuf definition.
If you need to change the name of a request/response, proceed as follows:
* call `godin update` to generate the default names and also to ensure that the endpoint exists in `godin.json`
* refactor the app (search and replace) to use the new request for example. Search for: `*pb.<Endpoint>Request` and replace with your actual name. Then also replace the name in `godin.json`
* calling `godin update` will now work as expected again, but using your custom type name

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.
