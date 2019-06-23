package middleware

import (
	"github.com/lukasjarosch/godin/examples/user/internal/service"
)

type Middleware func(service service.User) service.User
