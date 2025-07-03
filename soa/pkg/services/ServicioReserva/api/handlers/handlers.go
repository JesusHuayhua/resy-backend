package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	ReservaModels "soa/pkg/services/ServicioReserva/core/domain"
	"soa/pkg/services/ServicioReserva/core/usecase/backBD"
	"strconv"
	"time"
)

type Server struct {
	Svc *backBD.ServicioReserva
}

func NewServer(svc *backBD.ServicioReserva) *Server {
	return &Server{Svc: svc}
}

func habilitarCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func OpcionesHandler(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	w.WriteHeader(http.StatusOK)
}

// POST /reservas
func (s *Server) InsertarReserva(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	type tempReserva struct {
		IDCliente        *int      `json:"id_cliente"`
		NombreCliente    string    `json:"nombre_cliente"`
		TelefonoCliente  string    `json:"telefono_cliente"`
		CorreoCliente    string    `json:"correo_cliente"`
		FechaReservada   time.Time `json:"fecha_reservada"`
		NumPersonas      int       `json:"num_personas"`
		EstadoReserva    string    `json:"estado_reserva"`
		Especificaciones string    `json:"especificaciones"`
	}
	var tempReq tempReserva
	if err := json.NewDecoder(r.Body).Decode(&tempReq); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	// Validar el estado de la reserva
	estadosPermitidos := map[string]bool{
		"Pendiente":  true,
		"Confirmada": true,
		"Cancelada":  true,
	}
	if !estadosPermitidos[tempReq.EstadoReserva] {
		http.Error(w, "Estado de reserva inválido. Valores permitidos: Pendiente, Confirmada, Cancelada", http.StatusBadRequest)
		return
	}
	// Convierte a la estructura final con NullTypes
	req := ReservaModels.ReservaData{
		FechaReservada:   tempReq.FechaReservada,
		NumPersonas:      tempReq.NumPersonas,
		EstadoReserva:    tempReq.EstadoReserva,
		Especificaciones: tempReq.Especificaciones,
		NombreCliente:    sql.NullString{String: tempReq.NombreCliente, Valid: tempReq.NombreCliente != ""},
		TelefonoCliente:  sql.NullString{String: tempReq.TelefonoCliente, Valid: tempReq.TelefonoCliente != ""},
		CorreoCliente:    sql.NullString{String: tempReq.CorreoCliente, Valid: tempReq.CorreoCliente != ""},
	}

	if tempReq.IDCliente != nil {
		req.IDCliente = sql.NullInt64{Int64: int64(*tempReq.IDCliente), Valid: true}
	} else {
		req.IDCliente = sql.NullInt64{Valid: false}
	}

	if err := s.Svc.InsertarReserva(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}

// GET /reservas
// ListarReservas permite filtrar reservas por idReserva, idCliente, fecha y estadoReserva
// Si no se proporcionan filtros, devuelve todas las reservas.
// Los filtros se pasan como parámetros de consulta en la URL, por ejemplo:
// /reservas?idReserva=RES1&idCliente=123&fecha=2023-10-01&estadoReserva=Pendiente
// Los filtros son opcionales, si no se proporcionan, se devolverán todas las reservas
// Si se proporciona un filtro, se aplicará a la consulta.
func (s *Server) ListarReservas(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	filtros := make(map[string]interface{})

	if idReserva := r.URL.Query().Get("idReserva"); idReserva != "" {
		filtros["id_reserva"] = idReserva
	}

	if idClienteStr := r.URL.Query().Get("idCliente"); idClienteStr != "" {
		idCliente, err := strconv.Atoi(idClienteStr)
		if err != nil {
			http.Error(w, "IDCliente inválido", http.StatusBadRequest)
			return
		}
		filtros["id_cliente"] = idCliente
	}
	if fechaStr := r.URL.Query().Get("fecha"); fechaStr != "" {
		fecha, err := time.Parse("2006-01-02", fechaStr)
		if err != nil {
			http.Error(w, "Formato de fecha inválido. Usa YYYY-MM-DD", http.StatusBadRequest)
			return
		}
		filtros["fecha_reservada"] = fecha
	}
	if estado := r.URL.Query().Get("estadoReserva"); estado != "" {
		filtros["estado_reserva"] = estado
	}

	reservas, err := s.Svc.ListarReservas(filtros)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservas)
}

// PUT /reservas?id=RES1
func (s *Server) ActualizarReserva(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	// Maneja la solicitud OPTIONS (pre-vuelo CORS)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	// Obtiene el ID de la reserva de los parámetros de la URL
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Falta el parámetro id", http.StatusBadRequest)
		return
	}

	// Define una estructura temporal para la decodificación del JSON.
	// Todos los campos son punteros para distinguir entre "no enviado" y "valor vacío".
	type tempReserva struct {
		IDCliente        *int       `json:"id_cliente"`
		NombreCliente    *string    `json:"nombre_cliente"`
		TelefonoCliente  *string    `json:"telefono_cliente"`
		CorreoCliente    *string    `json:"correo_cliente"`
		FechaReservada   *time.Time `json:"fecha_reservada"`
		NumPersonas      *int       `json:"num_personas"`
		EstadoReserva    *string    `json:"estado_reserva"`
		Especificaciones *string    `json:"especificaciones"`
	}
	var tempReq tempReserva
	// Decodifica el cuerpo de la solicitud JSON en la estructura temporal
	if err := json.NewDecoder(r.Body).Decode(&tempReq); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	// Valida el estado de la reserva si se proporciona en la solicitud
	if tempReq.EstadoReserva != nil {
		estadosPermitidos := map[string]bool{
			"Pendiente":  true,
			"Confirmada": true,
			"Cancelada":  true,
		}
		if !estadosPermitidos[*tempReq.EstadoReserva] {
			http.Error(w, "Estado de reserva inválido. Valores permitidos: Pendiente, Confirmada, Cancelada", http.StatusBadRequest)
			return
		}
	}
	// Crea una instancia de ReservaModels.ReservaData para pasar al servicio.
	// Solo los campos que se proporcionaron en el JSON se asignarán.
	var req ReservaModels.ReservaData
	// Asigna campos no nulos directamente si están presentes
	if tempReq.FechaReservada != nil {
		req.FechaReservada = *tempReq.FechaReservada
	}
	if tempReq.NumPersonas != nil {
		req.NumPersonas = *tempReq.NumPersonas
	}
	if tempReq.EstadoReserva != nil {
		req.EstadoReserva = *tempReq.EstadoReserva
	}
	if tempReq.Especificaciones != nil {
		req.Especificaciones = *tempReq.Especificaciones
	}
	// Maneja campos que usan sql.NullString
	if tempReq.NombreCliente != nil {
		req.NombreCliente = sql.NullString{String: *tempReq.NombreCliente, Valid: true}
	} else {
		// Si el campo no se envió, se marca como inválido.
		// Esto le indica a la capa de servicio que no debe actualizar este campo.
		req.NombreCliente = sql.NullString{Valid: false}
	}
	if tempReq.TelefonoCliente != nil {
		req.TelefonoCliente = sql.NullString{String: *tempReq.TelefonoCliente, Valid: true}
	} else {
		req.TelefonoCliente = sql.NullString{Valid: false}
	}
	if tempReq.CorreoCliente != nil {
		req.CorreoCliente = sql.NullString{String: *tempReq.CorreoCliente, Valid: true}
	} else {
		req.CorreoCliente = sql.NullString{Valid: false}
	}
	// Maneja el campo IDCliente que usa sql.NullInt64
	if tempReq.IDCliente != nil {
		req.IDCliente = sql.NullInt64{Int64: int64(*tempReq.IDCliente), Valid: true}
	} else {
		req.IDCliente = sql.NullInt64{Valid: false}
	}
	// Llama al servicio para actualizar la reserva
	if err := s.Svc.ActualizarReserva(id, req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Envía una respuesta de éxito
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}

// DELETE /reservas?id=RES1
func (s *Server) EliminarReserva(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Falta el parámetro id", http.StatusBadRequest)
		return
	}
	if err := s.Svc.EliminarReserva(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}

func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "up", "time": time.Now()})
}
