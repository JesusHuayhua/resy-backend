package protobuf

import (
	"context"

	grpc "google.golang.org/grpc"
)

type StatusReply_Status int32

var StatusReply_Status_name = map[int32]string{
	0: "PENDING",
	1: "STARTED",
	2: "IN_PROGRESS",
	3: "FINISHED",
	4: "FAILED",
}

type UsuarioRequest struct {
	TicketID string `protobuf:"bytes,1,opt,name=ticketID" json:"ticketID,omitempty"`
	Option   string `protobuf:"bytes,2,opt,name=mark" json:"mark,omitempty"`
}

type UsuarioReply struct {
	Code int64  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Err  string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

type GetRequest_Filters struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

type ServiceStatusRequest struct {
}
type ServiceStatusReply struct {
	Code int64  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Err  string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

type GetRequest struct {
	Filters []*GetRequest_Filters `protobuf:"bytes,1,rep,name=filters" json:"filters,omitempty"`
}
type GetReply struct {
	Err string `protobuf:"bytes,2,opt,name=Err,json=err" json:"Err,omitempty"`
}
type StatusRequest struct {
	TicketID string `protobuf:"bytes,1,opt,name=ticketID" json:"ticketID,omitempty"`
}
type StatusReply struct {
	Status StatusReply_Status `protobuf:"varint,1,opt,name=status,enum=pb.StatusReply_Status" json:"status,omitempty"`
	Err    string             `protobuf:"bytes,2,opt,name=Err,json=err" json:"Err,omitempty"`
}

type UsuarioServer interface {
	Get(context.Context, *GetRequest) (*GetReply, error)
	Usuario(context.Context, *UsuarioRequest) (*UsuarioReply, error)
	Status(context.Context, *StatusRequest) (*StatusReply, error)
	ServiceStatus(context.Context, *ServiceStatusRequest) (*ServiceStatusReply, error)
}

var _UsuarioSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UsuarioSvc",
	HandlerType: (*UsuarioServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _UsuarioSvc_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usersvc.proto",
}

func _UsuarioSvc_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsuarioServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Watermark/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsuarioServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func RegisterUsuarioServer(s *grpc.Server, srv UsuarioServer) {
	s.RegisterService(&_UsuarioSvc_serviceDesc, srv)
}
