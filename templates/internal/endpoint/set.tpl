package endpoint

import (
	"{{ .Service.Name }}/internal"
)

// Set bundles all of the service's endpoints. Also, Set implements the ExampleService interface.
// That way a Set can also be used just like the service which get's handy when you need access to the service while
// keeping all applied middlewares. This is particularly useful if you want to trigger business logic from a consumer,
// skipping the transport layer.
type Set struct {}

// Endpoints initializes the Set with all endpoints including their middleware
func Endpoints(svc service.{{ title .Service.Name }}) Set {
	return Set{}
}
