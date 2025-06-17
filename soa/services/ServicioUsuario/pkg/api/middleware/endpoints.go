package middleware

/*
import (
	"ServicioUsuario/pkg/core/response"
	"ServicioUsuario/pkg/core/usecase/interfaces"
	"context"

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

// MakeStatusEndpoint creates an endpoint for status.
func MakeStatusEndpoint(svc interfaces.ServicioUsuario) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// Implement your status logic here
		// You may need to extract the string argument from the request, here using an empty string as a placeholder
		svc.Status(ctx, "")
		return "OK", nil
	}
}

func MakeGetEndpoint(svc interfaces.ServicioUsuario) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(response.GetRequest)
		// Convert []string to []internal.Filter
		var filters []internal.Filter
		for _, f := range req.Filters {
			filters = append(filters, internal.Filter{Value: f})
		}
		docs, err := svc.Get(ctx, filters...)
		if err != nil {
			// handle error appropriately, e.g. return nil, err
			return nil, err
		}
		return docs, nil
	}
}

// MakeServiceStatusEndpoint creates an endpoint for service status.
func MakeServiceStatusEndpoint(svc interfaces.ServicioUsuario) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// Implement your service status logic here
		status, err := svc.ServiceStatus(ctx)
		if err != nil {
			return nil, err
		}
		return status, nil
	}
}

// MakeUsuarioEndpoint creates an endpoint for usuario-related operations.
func MakeUsuarioEndpoint(svc interfaces.ServicioUsuario) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// Implement your usuario logic here. Adjust as needed.
		result, err := svc.Usuario(ctx, request)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
}
*/
