package transport

import (
	"context"
	"soa/services/users/pkg/api/middleware/endpoints"
	"soa/services/users/pkg/api/middleware/protobuf"
	"soa/services/users/pkg/core/response"
	"soa/services/users/pkg/core/svc_internal"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	get           grpctransport.Handler
	status        grpctransport.Handler
	usuario       grpctransport.Handler
	serviceStatus grpctransport.Handler
}

func NewGRPCServer(ep endpoints.Set) protobuf.UsuarioServer {
	return &grpcServer{
		get: grpctransport.NewServer(
			ep.GetEndpoint,
			decodeGRPCGetRequest,
			decodeGRPCGetResponse,
		),
		status: grpctransport.NewServer(
			ep.StatusEndpoint,
			decodeGRPCStatusRequest,
			decodeGRPCStatusResponse,
		),
		usuario: grpctransport.NewServer(
			ep.UsuarioEndpoint,
			decodeGRPCUsuarioRequest,
			decodeGRPCUsuarioResponse,
		),
		serviceStatus: grpctransport.NewServer(
			ep.ServiceStatusEndpoint,
			decodeGRPCServiceStatusRequest,
			decodeGRPCServiceStatusResponse,
		),
	}
}

func (g *grpcServer) Get(ctx context.Context, r *protobuf.GetRequest) (*protobuf.GetReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*protobuf.GetReply), nil
}

func (g *grpcServer) ServiceStatus(ctx context.Context, r *protobuf.ServiceStatusRequest) (*protobuf.ServiceStatusReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*protobuf.ServiceStatusReply), nil
}

func (g *grpcServer) Status(ctx context.Context, r *protobuf.StatusRequest) (*protobuf.StatusReply, error) {
	_, rep, err := g.status.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*protobuf.StatusReply), nil
}

func (g *grpcServer) Usuario(ctx context.Context, r *protobuf.UsuarioRequest) (*protobuf.UsuarioReply, error) {
	_, rep, err := g.usuario.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*protobuf.UsuarioReply), nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*protobuf.GetRequest)
	var filters []svc_internal.Filter
	for _, f := range req.Filters {
		filters = append(filters, svc_internal.Filter{Key: f.Key, Value: f.Value})
	}
	return response.GetRequest{Filters: filters}, nil
}

func decodeGRPCStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*protobuf.StatusRequest)
	return response.StatusRequest{TicketID: req.TicketID}, nil
}

func toInternalFilters(in []*protobuf.Filter) []svc_internal.Filter {
	out := make([]svc_internal.Filter, 0, len(in))
	for _, f := range in {
		if f == nil {
			continue
		}
		out = append(out, svc_internal.Filter{
			Key:   f.Key,
			Value: f.Value,
		})
	}
	return out
}

func decodeGRPCUsuarioRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*protobuf.UsuarioRequest)
	return response.GenericRequest{TipoOp: int(req.TipoOp), Args: toInternalFilters(req.Args)}, nil
}

func decodeGRPCServiceStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return response.ServiceStatusRequest{}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*protobuf.GetReply)
	return response.GetResponse{Err: reply.Err}, nil
}

func decodeGRPCStatusResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*protobuf.StatusReply)
	return response.StatusResponse{Status: svc_internal.StatusCode(reply.Status), Err: reply.Err}, nil
}

func decodeGRPCUsuarioResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*protobuf.UsuarioReply)
	//Nota: reply.Err no esta actualizado, gRPC se avanzara despues.
	return response.GenericResponse{Code: int(reply.Code), Data: reply.Err}, nil
}

func decodeGRPCServiceStatusResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*protobuf.ServiceStatusReply)
	return response.ServiceStatusResponse{Code: int(reply.Code), Err: reply.Err}, nil
}
