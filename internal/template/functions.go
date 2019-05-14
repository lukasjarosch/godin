package template

/*
func FunctionMap(data *Data) tpl.FuncMap {
	return tpl.FuncMap{
		"arg_list":           ArgumentList(data),
		"ret_list":           ReturnList(data),
		"enum_body":          EnumBody(data),
		"deps_param_list":    DependencyParameterList(data),
		"default_value_list": DefaultValueList(data),
		"has_dependency":     HasDependency(data),
		"deps_value_mapping": DependencyValueMapping(data),
		"deps_name_list":     DependencyNameList(data),
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
		return specification.ErrMethodUnspecified.Error()
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
		return specification.ErrMethodUnspecified.Error()
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
		return specification.ErrModelUnspecified.Error()
	}
}

func DependencyParameterList(data *Data) func() string {
	return func() string {
		var paramList []string
		for _, d := range data.Spec.ResolvedDependencies {
			format := "%s %s"

			paramList = append(paramList, fmt.Sprintf(format, d.Name(), d.Type()))
		}
		return strings.Join(paramList, ", ")
	}
}

func DefaultValueList(data *Data) func(vars []specification.Variable) string {
	return func(vars []specification.Variable) string {
		var list []string

		for _, v := range vars {
			list = append(list, v.DefaultValue(data.Spec))
		}

		return strings.Join(list, ", ")
	}
}

func HasDependency(data *Data) func(depType string) bool {
	return func(depType string) bool {
		for _, dep := range data.Spec.ResolvedDependencies {
			if dep.Name() == depType {
				return true
			}
		}
		return false
	}
}

func DependencyValueMapping(data *Data) func() string {
	return func() string {
		var format = "%s: %s,"
		var mappings []string

		for _, dep := range data.Spec.ResolvedDependencies {
			mappings = append(mappings, fmt.Sprintf(format, dep.Name(), dep.Name()))
		}
		return strings.Join(mappings, "\n")
	}
}

func DependencyNameList(data *Data) func() string {
	return func() string {
		var depList []string

		for _, dep := range data.Spec.ResolvedDependencies {
			depList = append(depList, dep.Name())
		}

		return strings.Join(depList, ", ")
	}
}
*/
