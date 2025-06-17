package protobuf

import "context"

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
