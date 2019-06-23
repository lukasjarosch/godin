package user

import (
	"context"
	"github.com/lukasjarosch/godin/examples/user/internal/service"
	"github.com/lukasjarosch/godin/pkg/log"
)

type serviceImplementation struct {
	logger log.Logger
}

func NewServiceImplementation(logger log.Logger) *serviceImplementation {
	return &serviceImplementation{
		logger: logger,
	}
}

// Create will create a new user and return it.
func (s *serviceImplementation) Create(ctx context.Context, username string, email string) (user *service.User, err error) {
}
