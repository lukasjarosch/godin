package middleware

import (
    "{{ .Service.Module }}/internal/service"
)

// Middleware defines the service-specific middleware type
type Middleware func(service service.{{ title .Service.Name }}) service.{{ title .Service.Name }}
