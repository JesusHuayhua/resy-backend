package middleware

import (
	"context"
	"soa/services/users/pkg/core/response"
	"soa/services/users/pkg/core/usecase/interfaces"

	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	GetEndpoint           endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
	StatusEndpoint        endpoint.Endpoint
	UsuarioEndpoint       endpoint.Endpoint
}

func NewEndpoints(svc interfaces.ServicioUsuario) Set {
	return Set{
		GetEndpoint:           MakeGetEndpoint(svc),
		StatusEndpoint:        MakeStatusEndpoint(svc),
		ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
		UsuarioEndpoint:       MakeUsuarioEndpoint(svc),
	}
}

func MakeGetEndpoint(svc interfaces.ServicioUsuario) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(response.GetRequest)
		docs, err := svc.Get(ctx, req.Filters...)
		if err != nil {

		}
		return GetResponse()
	}
}
