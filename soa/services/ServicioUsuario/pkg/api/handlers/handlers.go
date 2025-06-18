package handlers

import (
	BDoperators "ServicioUsuario/pkg/core/usecase/backBD"
	"encoding/json"
	"net/http"
	"time"
)

type Server struct {
	Svc *BDoperators.ServicioUsuario // Cambiado a puntero para evitar copia innecesaria
}

func NewServer(svc *BDoperators.ServicioUsuario) *Server {
	return &Server{Svc: svc}
}

// POST /usuarios
func (s *Server) InsertarUsuario(w http.ResponseWriter, r *http.Request) {
	type reqBody struct {
		Nombres         string `json:"nombres"`
		Apellidos       string `json:"apellidos"`
		Correo          string `json:"correo"`
		Telefono        string `json:"telefono"`
		FechaNacimiento string `json:"fechanacimiento"`
		Contrasenia     string `json:"contrasenia"`
		Rol             int    `json:"rol"`
	}
	var req reqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	fecha, err := time.Parse("2006-01-02", req.FechaNacimiento)
	if err != nil {
		http.Error(w, "Fecha inválida", http.StatusBadRequest)
		return
	}
	status, err := s.Svc.InsertarNuevoUsuario(
		req.Nombres, req.Apellidos, req.Correo, req.Telefono, fecha, req.Contrasenia, req.Rol,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": status})
}

// GET /usuarios
func (s *Server) ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	_, usuarios, err := s.Svc.SeleccionarUsuarios("", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}
