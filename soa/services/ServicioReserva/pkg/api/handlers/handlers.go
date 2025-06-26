package handlers

import (
	ReservaModels "ServicioReserva/pkg/core/domain"
	"ServicioReserva/pkg/core/usecase/backBD"
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
	var req ReservaModels.ReservaData
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	// Parse fecha si es string
	// req.FechaReservada = ...
	if err := s.Svc.InsertarReserva(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}

// GET /reservas
func (s *Server) ListarReservas(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	reservas, err := s.Svc.ListarReservas("", nil)
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
