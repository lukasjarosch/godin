package service

{{ $service := title .Service.Name }}

// TODO: {{ $service }} documentation
type {{ $service }} interface {
    // TODO: specify the service endpoints or use 'godin add' to add new endpoints
}

// Application errors
var (
    ErrNotImplemented = errors.New("endpoint not implemented")
)