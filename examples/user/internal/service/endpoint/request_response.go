// Code generated by Godin v0.3.0; DO NOT EDIT.

package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"

	service "github.com/lukasjarosch/godin/examples/user"
)

type (
	CreateRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	CreateResponse struct {
		User *UserEntity `json:"user"`
		Err  error       `json:"-"`
	}
)

// Implement the Failer interface for all responses
func (resp CreateResponse) Failed() error { return r.Err }