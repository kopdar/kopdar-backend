package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/kopdar/kopdar-backend/internal/user"
)

type Set struct {
	MakeFindAllEndpoint endpoint.Endpoint
}

func New(userService user.Service) *Set {
	var findAllEndpoint endpoint.Endpoint
	{
		findAllEndpoint = MakeFindAllEndpoint(userService)
	}

	return &Set{
		MakeFindAllEndpoint: findAllEndpoint,
	}
}

func MakeFindAllEndpoint(s user.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.FindAll(ctx)
	}
}
