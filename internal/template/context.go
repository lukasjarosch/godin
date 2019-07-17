package template

import (
	"fmt"
	"github.com/lukasjarosch/godin/internal/bundle/transport"
	"strings"
	"time"

	"github.com/go-godin/rabbitmq"
	"github.com/lukasjarosch/godin/internal/bundle"
	"github.com/mitchellh/mapstructure"

	"github.com/lukasjarosch/godin/internal"
	"github.com/lukasjarosch/godin/internal/parse"
	config "github.com/spf13/viper"
	"github.com/vetcher/go-astra/types"
)

type Context struct {
	Service  Service
	Godin    Godin
	Protobuf Protobuf
	Docker   Docker
	Project Project
}

// NewContextFromConfig will initialize the context will all the data from the configuration
// The context is not fully populated after this call, but all configuration values are accessible.
func NewContextFromConfig() Context {
	var subscribers []Subscriber

	// amqp subscribers
	sub := config.GetStringMap(bundle.SubscriberKey)
	if len(sub) > 0 {
		for x := range sub {
			s := &rabbitmq.Subscription{}
			mapstructure.Decode(sub[x], s)
			subscribers = append(subscribers, Subscriber{
				Subscription: *s,
				Handler:      bundle.SubscriberHandlerName(s.Topic),
			})
		}
	}

	ctx := Context{
		Service: Service{
			Name:              config.GetString("service.name"),
			Namespace:         config.GetString("service.namespace"),
			Module:            config.GetString("service.module"),
			LoggingMiddleware: config.GetBool("service.middleware.logging"),
			Subscriber:        subscribers,
			Transport:Transport{
				GRPC: config.GetBool("transport.grpc.enabled"),
				AMQP: config.GetBool(transport.AMQPTransportEnabledKey),
			},
		},
		Protobuf: Protobuf{
			Package: config.GetString("protobuf.package"),
			Service: config.GetString("protobuf.service"),
		},
		Project:Project{
			Created: config.GetTime("project.created"),
			Updated: config.GetTime("project.updated"),
		},
		Godin: Godin{
			Version: internal.Version,
			Build:   internal.Build,
			Commit:  internal.Commit,
		},
		Docker: Docker{
			Registry: config.GetString("docker.registry"),
		},
	}

	return ctx
}

// PopulateFromService will populate an existing Context with the available data from the parse service-file
func PopulateFromService(ctx Context, service *parse.Service) Context {
	serviceName := config.GetString("service.name")

	// map go-astra/types Method to our Method struct
	var methods []Method
	for _, meth := range service.Interface.Methods {

		// methods
		var params []Variable
		for _, arg := range meth.Args {
			params = append(params, Variable{
				Name: arg.Name,
				Type: arg.Type.String(),
			})
		}

		var returns []Variable
		for _, ret := range meth.Results {
			returns = append(returns, Variable{
				Name: ret.Name,
				Type: ret.Type.String(),
			})
		}

		methods = append(methods, Method{
			Name:             meth.Name,
			Comments:         meth.Docs,
			Params:           params,
			Returns:          returns,
			ServiceName:      serviceName,
			Request:          fmt.Sprintf("%sRequest", meth.Name),
			Response:         fmt.Sprintf("%sResponse", meth.Name),
			ProtobufRequest:  config.GetString(fmt.Sprintf("service.endpoints.%s.protobuf.request", meth.Name)),
			ProtobufResponse: config.GetString(fmt.Sprintf("service.endpoints.%s.protobuf.response", meth.Name)),
		})
	}

	ctx.Service.Methods = methods

	return ctx
}

type Godin struct {
	Version string
	Commit  string
	Build   string
}

type Project struct {
	Created time.Time
	Updated time.Time
}

type Service struct {
	Name              string
	Namespace         string
	Methods           []Method
	Module            string
	LoggingMiddleware bool
	Subscriber        []Subscriber
	Transport         Transport
}

type Protobuf struct {
	Package string
	Service string
}

type Transport struct {
	GRPC bool
	AMQP bool
}

type Docker struct {
	Registry string
}

type Subscriber struct {
	Handler      string
	Subscription rabbitmq.Subscription
}

type Variable struct {
	Name string
	Type string
}

func (v Variable) resolveType(typ string, prefix string) (string, error) {
	if strings.HasPrefix(typ, prefix) {
		trimmed := strings.TrimLeft(v.Type, prefix)
		if types.IsBuiltinTypeString(trimmed) {
			return fmt.Sprintf("%s%s", prefix, trimmed), nil
		}

		return fmt.Sprintf("%s%s.%s", prefix, "service", trimmed), nil
	}

	return "", fmt.Errorf("type %s cannot be resolved with prefix %s", typ, prefix)
}

// ResolveType resolves the type to use inside a template. It covers different combinations which should suffice most cases.
func (v Variable) ResolveType() string {
	prefixes := []string{"[]*", "*[]", "[]", "*"}

	if types.IsBuiltinTypeString(v.Type) {
		return v.Type
	}

	if strings.Contains(v.Type, ".") {
		return v.Type
	}

	for _, prefix := range prefixes {
		typ, err := v.resolveType(v.Type, prefix)
		if err == nil {
			return typ
		}
	}

	return fmt.Sprintf("%s.%s", "service", v.Type)
}

// NilValue returns the appropriate nil value for the variable
func (v Variable) NilValue() string {
	switch v.Type {
	case "error":
		return "nil"
	case "string":
		return "\"\""
	case "boolean":
		return "false"
	case "int",
		"int16",
		"int32",
		"int64",
		"float32",
		"float64":
		return "0"
	default:
		return "nil"
	}
}

type Method struct {
	// required for partials which do not have access to the Service struct
	ServiceName      string
	Comments         []string
	Name             string
	Params           []Variable
	Returns          []Variable
	Request          string
	Response         string
	ProtobufRequest  string
	ProtobufResponse string
}

func (m Method) RequestName() string {
	return fmt.Sprintf("%sRequest", m.Name)
}

func (m Method) ResponseName() string {
	return fmt.Sprintf("%sResponse", m.Name)
}

func (m Method) ParamList() string {
	var list []string

	for _, arg := range m.Params {
		_, ok := types.BuiltinTypes[arg.Type]
		if !ok {
			list = append(list, fmt.Sprintf("%s %s", arg.Name, arg.ResolveType()))
		} else {
			list = append(list, fmt.Sprintf("%s %s", arg.Name, arg.Type))
		}
	}

	return strings.Join(list, ", ")
}

func (m Method) ReturnList() string {
	var list []string

	for _, arg := range m.Returns {
		_, ok := types.BuiltinTypes[arg.Type]
		if !ok {
			list = append(list, fmt.Sprintf("%s %s", arg.Name, arg.ResolveType()))
		} else {
			list = append(list, fmt.Sprintf("%s %s", arg.Name, arg.Type))
		}
	}

	return strings.Join(list, ", ")
}

// ReturnImplementationMissing will produce a list of
// nil returns and an error. For example, the return params are (name string, err error)
// the method will return "nil, fmt.Errorf("NOT IMPLEMENTED")"
func (m Method) ReturnImplementationMissing() string {
	var list []string
	for _, ret := range m.Returns {
		if ret.Type == "error" {
			list = append(list, "domain.ErrNotImplemented")
		} else {
			list = append(list, ret.NilValue())
		}
	}

	return strings.Join(list, ", ")
}

// ReturnVariableList returns all return variable names as  comma-separated string
func (m Method) ReturnVariableList() string {
	var list []string

	for _, v := range m.Returns {
		list = append(list, v.Name)
	}

	return strings.Join(list, ",")
}

func MethodFromType(function *types.Function) Method {

	var args []Variable
	for _, arg := range function.Args {
		args = append(args, Variable{
			Name: arg.Name,
			Type: arg.Type.String(),
		})
	}

	var returns []Variable
	for _, ret := range function.Results {
		returns = append(returns, Variable{
			Name: ret.Name,
			Type: ret.Type.String(),
		})
	}

	return Method{
		Name:     function.Name,
		Params:   args,
		Returns:  returns,
		Comments: function.Docs,
	}
}
