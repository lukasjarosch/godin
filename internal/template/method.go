package template

import (
	"io/ioutil"
	"path"

	"bytes"
	"text/template"

	"github.com/lukasjarosch/godin/internal/specification"
	"fmt"
	"strings"
)

type MethodPartial interface {
	Comments() []string
	ArgList() string
	ReturnList() string
	Name() string
	Receiver() string
}

type methodTemplate struct {
	method specification.ServiceMethod
	service specification.Service
}

func NewMethodTemplate(service specification.Service, method specification.ServiceMethod) *methodTemplate {
	return &methodTemplate{
		method: method,
		service:service,
	}
}

func (m *methodTemplate) Render() (string, error) {
	templatePath := path.Join(".", "templates", "partials", "method.tpl")
	templateData, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return "", err
	}

	tpl, err := template.New(templatePath).Parse(string(templateData))

	var out bytes.Buffer
	err = tpl.Execute(&out, m)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func (m *methodTemplate) Comments() []string {
	var c []string

	for _, comment := range m.method.Comments {
		c = append(c, comment)
	}

	return c
}

func (m *methodTemplate) ArgList() string {
	var argList []string

	for _, arg := range m.method.Arguments {
		argList = append(argList, arg.String())
	}

	return strings.Join(argList, ", ")
}

func (m *methodTemplate) ReturnList() string {
	var retList []string

	for _, ret := range m.method.Returns {
		retList = append(retList, ret.String())
	}

	return strings.Join(retList, ", ")
}

func (m *methodTemplate) Name() string {
	return m.method.Name
}

func (m *methodTemplate) Receiver() string {
	return fmt.Sprintf("svc *%s", strings.Title(m.service.Name))
}