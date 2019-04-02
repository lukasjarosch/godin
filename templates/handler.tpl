package server

import (
	"context"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"

    pb "{{ .Spec.Service.API.Import }}"
	service "{{ .ModuleName }}/internal/{{ .ServiceName }}"
)

// greeterHandler is the transport-layer wrapper of our business-logic in the server package
// Everything concerning requests/responses belongs in here. Only conversion (business-model <-> protobuf) should happen here actually.
type {{ .ServiceName }}Handler struct {
	implementation *service.{{ .ServiceName | camelcase }}
}

func New{{ .GrpcServiceName }}Handler(implementation *service.{{ .ServiceName | camelcase }}) *{{ .ServiceName }}Handler{
	return &{{ .ServiceName }}Handler{
		implementation: implementation,
	}
}

{{ $serviceName := .ServiceName }}
{{ $grpcService := .GrpcServiceName }}
{{ $apiPackage := .Spec.Service.API.Package }}
{{ range .Spec.Service.Methods }}
// {{ .Name }} is the gRPC handler for {{ $apiPackage }}.{{ .Name }}()
func (e *{{ $serviceName }}Handler) {{ .Name }}(ctx context.Context, request *pb.{{ .Name }}Request) (*pb.{{ .Name }}Response, error) {
    // TODO: call e.implementation.{{ .Name }} and return the response
    return nil, status.Error(codes.Unimplemented, "{{ $apiPackage }}.{{ $grpcService }}.{{ .Name }}() unimplemented")
}
{{- end }}

