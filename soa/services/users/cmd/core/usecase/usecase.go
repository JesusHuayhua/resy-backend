package usecase

import (
	"soa/services/users/cmd/core/internal"
)

type GetRequest struct {
	Filters []internal.StatusCode `json:"data,omitempty"`
}

type StatusRequest struct {
	TicketID string `json:"ticketID,omitempty"`
}

type StatusResponse struct {
	Status int `json:"status"`
}

type ServiceStatusRequest struct {
}

type ServiceStatusResponse struct {
	Code int    `json:"status"`
	Err  string `json:"err, omitempty"`
}
