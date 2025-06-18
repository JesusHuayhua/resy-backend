package endpoints

import (
	"context"
	"errors"
	"os"
	"soa/services/users/pkg/core/response"
	"soa/services/users/pkg/core/svc_internal"
	"soa/services/users/pkg/core/usecase/interfaces"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type Set struct {
	GetEndpoint           endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
	StatusEndpoint        endpoint.Endpoint
	UsuarioEndpoint       endpoint.Endpoint
}

func NewEndpoints(svc interfaces.UserService) Set {
	return Set{
		GetEndpoint:           MakeGetEndpoint(svc),
		StatusEndpoint:        MakeStatusEndpoint(svc),
		ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
		UsuarioEndpoint:       MakeUsuarioEndpoint(svc),
	}
}

func MakeGetEndpoint(svc interfaces.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(response.GetRequest)
		err := svc.Get(ctx, req.Filters...)
		if err != nil {
			return response.GetResponse{Err: err.Error()}, nil
		}
		return response.GetResponse{Err: "Success"}, nil
	}
}

func MakeStatusEndpoint(svc interfaces.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(response.StatusRequest)
		status, err := svc.Status(ctx, req.TicketID)
		if status != svc_internal.Error {
			return response.StatusResponse{Status: status, Err: err.Error()}, nil
		}
		return response.StatusResponse{Status: status, Err: ""}, nil
	}
}

func MakeServiceStatusEndpoint(svc interfaces.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(response.ServiceStatusRequest)
		code, err := svc.ServiceStatus(ctx)
		if err != nil {
			return response.ServiceStatusResponse{Code: code, Err: err.Error()}, nil
		}
		return response.ServiceStatusResponse{Code: code, Err: ""}, nil
	}
}

func MakeUsuarioEndpoint(svc interfaces.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(response.UsuarioRequest)
		code, err := svc.Usuario(ctx, req.TipoOp, req.Args)
		if err != nil {
			return response.UsuarioResponse{Code: code, Err: err.Error()}, nil
		}
		return response.UsuarioResponse{Code: code, Err: ""}, nil
	}
}

func (s *Set) Get(ctx context.Context, filters ...svc_internal.Filter) error {
	resp, err := s.GetEndpoint(ctx, response.GetRequest{Filters: filters})
	if err != nil {
		return err
	}
	getResp := resp.(response.GetResponse)
	if getResp.Err != "" {
		return errors.New(getResp.Err)
	}
	return nil
}

func (s *Set) ServiceStatus(ctx context.Context) (int, error) {
	resp, err := s.ServiceStatusEndpoint(ctx, response.ServiceStatusRequest{})
	svcStatusResp := resp.(response.ServiceStatusResponse)
	if err != nil {
		return svcStatusResp.Code, err
	}
	if svcStatusResp.Err != "" {
		return svcStatusResp.Code, errors.New(svcStatusResp.Err)
	}
	return svcStatusResp.Code, nil
}

func (s *Set) Status(ctx context.Context, ticketID string) (svc_internal.StatusCode, error) {
	resp, err := s.StatusEndpoint(ctx, response.StatusRequest{TicketID: ticketID})
	if err != nil {
		return svc_internal.Error, err
	}
	stsResp := resp.(response.StatusResponse)
	if stsResp.Err != "" {
		return svc_internal.Error, errors.New(stsResp.Err)
	}
	return stsResp.Status, nil
}

func (s *Set) Usuario(ctx context.Context, tipoOP int, args []svc_internal.Filter) (int, error) {
	resp, err := s.UsuarioEndpoint(ctx, response.UsuarioRequest{TipoOp: tipoOP, Args: args})
	userResponse := resp.(response.UsuarioResponse)
	if err != nil {
		return userResponse.Code, err
	}
	if userResponse.Err != "" {
		return userResponse.Code, errors.New(userResponse.Err)
	}
	return userResponse.Code, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
