// Respuesta de servicio de usuario
// Este paquete define las estructuras de respuesta y solicitud para el servicio de usuario.
// Estas estructuras son utilizadas para manejar las solicitudes y respuestas del servicio de usuario,
// incluyendo la obtención de estados de tickets y solicitudes de estado del servicio.
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
