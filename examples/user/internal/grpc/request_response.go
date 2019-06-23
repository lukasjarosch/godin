// Code generated by Godin v0.3.0; DO NOT EDIT.

package grpc

import (
	"context"
	"errors"

	service "github.com/lukasjarosch/godin/examples/user"
	"github.com/lukasjarosch/godin/examples/user/internal/service/endpoint"
)

// EncodeCreateRequest is used in the server and encodes a domain-level request into a gRPC request
func EncodeCreateRequest(ctx context.Context, request interface{}) (pbRequest interface{}, err error) {
	if request == nil {
		return nil, errors.New("nil CreateRequest")
	}
	req := request.(endpoint.CreateRequest)
	pbRequest, err := CreateRequestEncoder(req)
	if err != nil {
		return nil, err
	}
	return pbRequest, nil
}

// DecodeCreateResponse is used in the server and decodes a gRPC response into a domain-level response
func DecodeCreateResponse(ctx context.Context, pbResponse interface{}) (response interface{}, err error) {
	if pbResponse == nil {
		return nil, errors.New("nil CreateResponse")
	}
	response, err := CreateResponseDecoder(res)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// EncodeCreateResponse is used in the client and encodes a domain-level response into a gRPC response
func EncodeCreateResponse(ctx context.Context, response interface{}) (pbResponse interface{}, err error) {
	if response == nil {
		return nil, errors.New("nil CreateResponse")
	}
	res := response.(endpoint.CreateResponse)
	pbResponse, err := CreateResponseEncoder(res)
	if err != nil {
		return nil, err
	}
	return pbResponse, nil
}

// DecodeCreateRequest is used in the client and decodes a gRPC request into a domain-level request
func DecodeCreateRequest(ctx context.Context, rbRequest interface{}) (request interface{}, err error) {
	if pbRequest == nil {
		return nil, errors.New("nil CreateRequest")
	}
	request, err := CreateRequestDecoder(req)
	if err != nil {
		return nil, err
	}
	return request, nil
}
