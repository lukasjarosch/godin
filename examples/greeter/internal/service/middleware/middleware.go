package middleware

import (
	"yyy/internal/service"
)

type Middleware func(service service.Yyy) service.Yyy
