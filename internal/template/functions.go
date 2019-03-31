package template

import (
	tpl "text/template"
	"strings"
)

func FunctionMap(data *Data) tpl.FuncMap {
	return tpl.FuncMap{
		"arg_list": ArgumentList(data),
		"ret_list": ReturnList(data),
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
		return "UNSPECIFIED"
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
		return "UNSPECIFIED"
	}
}
