package handlers

import (
	ReservaModels "ServicioReserva/pkg/core/domain"
	"ServicioReserva/pkg/core/usecase/backBD"
	"database/sql"
	"encoding/json"
	"net/http"
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
func (s *Server) ListarReservas(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	// Cambia esto:
	reservas, err := s.Svc.ListarReservas("") // Sin segundo parámetro
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reservas)
}

// PUT /reservas?id=RES1
func (s *Server) ActualizarReserva(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Falta el parámetro id", http.StatusBadRequest)
		return
	}
	var req ReservaModels.ReservaData
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	if err := s.Svc.ActualizarReserva(id, req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
