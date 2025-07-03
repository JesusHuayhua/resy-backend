package response

type GetRequest struct {
	Filters []string `json:"status,omitempty"`
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
	Err  string `json:"err,omitempty"`
}
