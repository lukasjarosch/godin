package template

import (
	"fmt"
	"strings"

	"github.com/lukasjarosch/godin/internal"
	"github.com/lukasjarosch/godin/internal/parse"
	config "github.com/spf13/viper"
	"github.com/vetcher/go-astra/types"
)

type Context struct {
	Service Service
	Godin   Godin
}

// NewContextFromConfig will initialize the context will all the data from the configuration
// The context is not fully populated after this call, but all configuration values are accessible.
func NewContextFromConfig() Context {
	ctx := Context{
		Service: Service{
			Name:      config.GetString("service.name"),
			Namespace: config.GetString("service.namespace"),
			Module:    config.GetString("service.module"),
		},
		Godin: Godin{
			Version: internal.Version,
			Build:   internal.Build,
			Commit:  internal.Commit,
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
			Name:        meth.Name,
			Comments:    meth.Docs,
			Params:      params,
			Returns:     returns,
			ServiceName: serviceName,
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

type Service struct {
	Name      string
	Namespace string
	Methods   []Method
	Module    string
}

type Variable struct {
	Name string
	Type string
}

// ResolveType resolves the type to use inside a template. It covers different combinations which should suffice most cases.
func (v Variable) ResolveType() string {

	if strings.Contains(v.Type, ".") {
		return v.Type
	}

	if strings.HasPrefix(v.Type, "[]*") {
		return fmt.Sprintf("[]*%s.%s", "service", strings.TrimLeft(v.Type, "[]*"))
	}

	if strings.HasPrefix(v.Type, "*[]") {
		return fmt.Sprintf("*[]%s.%s", "service", strings.TrimLeft(v.Type, "[]*"))
	}

	if strings.HasPrefix(v.Type, "*") {
		return fmt.Sprintf("*%s.%s", "service", strings.TrimLeft(v.Type, "*"))
	}

	if strings.HasPrefix(v.Type, "[]") {
		return fmt.Sprintf("[]%s.%s", "service", strings.TrimLeft(v.Type, "[]"))
	}

	return fmt.Sprintf("%s.%s", "service", v.Type)
}

type Method struct {
	// required for partials which do not have access to the Service struct
	ServiceName string
	Comments    []string
	Name        string
	Params      []Variable
	Returns     []Variable
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
		Name: function.Name,
		Params: args,
		Returns: returns,
		Comments: function.Docs,
	}
}