package template

import (
	tpl "text/template"
	"strings"
	"fmt"
	"github.com/lukasjarosch/godin/internal/specification"
)

func FunctionMap(data *Data) tpl.FuncMap {
	return tpl.FuncMap{
		"arg_list": ArgumentList(data),
		"ret_list": ReturnList(data),
		"enum_body": EnumBody(data),
	}
}

// ArgumentList returns the argument list of a given function
// The data is extracted from the specification
func ArgumentList(data *Data) func(method string) string {
	return func(method string) string {
		for _, meth := range data.Spec.Service.Methods {
			if meth.Name == method {
				var argList []string

				for _, arg := range meth.Arguments {
					argList = append(argList, arg.String())
				}

				return strings.Join(argList, ", ")
			}
		}
		return "UNSPECIFIED METHOD"
	}
}

func ReturnList(data *Data) func(method string) string {
	return func(method string) string {
		for _, meth := range data.Spec.Service.Methods {
			if meth.Name == method {
				var retList []string

				for _, ret := range meth.Returns {
					retList = append(retList, ret.String())
				}

				return strings.Join(retList, ", ")
			}
		}
		return "UNSPECIFIED METHOD"
	}
}

func EnumBody(data *Data) func(enum specification.Enumeration) string {
	return func(spec specification.Enumeration) string {
		format := "%s = %d"

		for _, e := range data.Spec.Models.Enums {
			if e.Name == spec.Name {
				var body []string

				for i, field := range e.Fields {
					body = append(body, fmt.Sprintf(format, field, i))
				}
				return strings.Join(body, "\n")
			}
		}
		return "UNSPECIFIED MODEL"
	}
}
