package middleware

import (
    "{{ .Service.Module }}/internal"
)

type Middleware func(service service.{{ title .Service.Name }}) service.{{ title .Service.Name }}