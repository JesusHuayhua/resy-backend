package middleware

import (
	"context"
	"soa/services/users/pkg/core/response"
	"soa/services/users/pkg/core/svc_internal"
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
		err := svc.Get(ctx, req.Filters...)
		if err != nil {
			return response.GetResponse{Err: err.Error()}, nil
		}
		return response.GetResponse{Err: "Success"}, nil
	}
}

func MakeStatusEndpoint(svc interfaces.ServicioUsuario) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(response.StatusRequest)
		status, err := svc.Status(ctx, req.TicketID)
		if status != svc_internal.Error {
			return response.StatusResponse{Status: status, Err: err.Error()}, nil
		}
		return response.StatusResponse{Status: status, Err: ""}, nil
	}
}

func MakeServiceStatusEndpoint(svc interfaces.ServicioUsuario) endpoint.Endpoint {

}

func MakeUsuarioEndpoint(svc interfaces.ServicioUsuario) endpoint.Endpoint {

}
