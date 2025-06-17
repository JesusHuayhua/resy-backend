package response

import "soa/services/users/pkg/core/svc_internal"

type GetRequest struct {
	Filters []svc_internal.Filter `json:"status,omitempty"`
}

type StatusRequest struct {
	TicketID string `json:"ticketID,omitempty"`
}

type GetResponse struct {
	Err string `json:"err,omitempty"`
}

type StatusResponse struct {
	Status svc_internal.StatusCode `json:"status"`
	Err    string                  `json:"err,omitempty"`
}

type UsuarioResponse struct {
	ticketID string `json:"ticketID"`
	option   string `json:"option"`
}

type ServiceStatusRequest struct {
}

type ServiceStatusResponse struct {
	Code int    `json:"status"`
	Err  string `json:"err, omitempty"`
}
