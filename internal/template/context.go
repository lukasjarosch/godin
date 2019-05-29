package template

import (
	"fmt"
	"strings"
)

type Context struct {
	Service Service
	Godin Godin
}

type Godin struct {
	Version string
	Commit string
	Build string
}

type Service struct {
	Name       string
	Namespace string
	Methods    []Method
	ImportPath string
}

type Variable struct {
	Name string
	Type string
}

type Method struct {
	// required for partials which do not have access to the Service struct
	ServiceName string
	Name string
	Params []Variable
	Returns []Variable
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
		list = append(list, fmt.Sprintf("%s %s", arg.Name, arg.Type))
	}

	return strings.Join(list, ", ")
}

func (m Method) ReturnList() string {
	var list []string

	for _, arg := range m.Returns {
		list = append(list, fmt.Sprintf("%s %s", arg.Name, arg.Type))
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


