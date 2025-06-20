// Package handlers proporciona los controladores HTTP para la gestión de usuarios y roles.
package handlers

import (
	BDoperators "ServicioUsuario/pkg/core/usecase/backBD"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	Svc *BDoperators.ServicioUsuario // Cambiado a puntero para evitar copia innecesaria
}

func NewServer(svc *BDoperators.ServicioUsuario) *Server {
	return &Server{Svc: svc}
}

// habilitarCORS configura los encabezados CORS para permitir solicitudes de cualquier origen.
func habilitarCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

// Handler global para OPTIONS
func OpcionesHandler(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	w.WriteHeader(http.StatusOK)
}

// POST /usuarios
// InsertarUsuario maneja la creación de un nuevo usuario a través de una petición HTTP POST.
// Decodifica el cuerpo de la petición, valida los datos y llama al servicio correspondiente.
func (s *Server) InsertarUsuario(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	type reqBody struct {
		Nombres         string `json:"nombres"`
		Apellidos       string `json:"apellidos"`
		Correo          string `json:"correo"`
		Telefono        string `json:"telefono"`
		Direccion       string `json:"direccion"`
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
		req.Nombres, req.Apellidos, req.Correo, req.Telefono, req.Direccion, fecha, req.Contrasenia, req.Rol,
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
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	_, usuarios, err := s.Svc.SeleccionarUsuarios("", []interface{}{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// PUT /usuarios/{id}
func (s *Server) ActualizarUsuario(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	type reqBody struct {
		Id              int    `json:"id"`
		Nombres         string `json:"nombres"`
		Apellidos       string `json:"apellidos"`
		Correo          string `json:"correo"`
		Telefono        string `json:"telefono"`
		Direccion       string `json:"direccion"`
		FechaNacimiento string `json:"fechanacimiento"`
		Contrasenia     string `json:"contrasenia"`
		Rol             int    `json:"rol"`
		EstadoAcceso    bool   `json:"estadoacceso"`
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
	status, err := s.Svc.ActualizarUsuario(req.Id, req.Nombres, req.Apellidos, req.Correo, req.Telefono, req.Direccion, fecha, req.Contrasenia, req.Rol, req.EstadoAcceso)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": status})
}

// DELETE /usuarios/{id}
func (s *Server) EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Falta el parámetro id", http.StatusBadRequest)
		return
	}
	var idInt int
	fmt.Sscanf(id, "%d", &idInt)
	status, err := s.Svc.EliminarUsuario(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": status})
}

// POST /roles
func (s *Server) InsertarRol(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	type reqBody struct {
		NombreRol string `json:"nombrerol"`
	}
	var req reqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	status, err := s.Svc.InsertarNuevoRol(req.NombreRol)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": status})
}

// PUT /roles/{id}
func (s *Server) ActualizarRol(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Falta el parámetro id", http.StatusBadRequest)
		return
	}
	type reqBody struct {
		Id        int    `json:"id"`
		NombreRol string `json:"nombrerol"`
	}
	var req reqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	status, err := s.Svc.ActualizarRol(req.Id, req.NombreRol)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": status})
}

// DELETE /roles/{id}
func (s *Server) EliminarRol(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Falta el parámetro id", http.StatusBadRequest)
		return
	}
	var idInt int
	fmt.Sscanf(id, "%d", &idInt)
	status, err := s.Svc.EliminarRol(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": status})
}

// GET /roles
func (s *Server) ListarRoles(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	_, roles, err := s.Svc.SeleccionarRoles("", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}

// POST /recuperar
func (s *Server) IniciarRecuperacionPassword(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	type reqBody struct {
		Correo string `json:"correo"`
	}
	var req reqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	token, err := s.Svc.IniciarRecuperacionPassword(req.Correo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"token": token})
}

// POST /recuperar/confirmar
func (s *Server) RecuperarPassword(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	type reqBody struct {
		Correo        string `json:"correo"`
		Token         string `json:"token"`
		NuevaPassword string `json:"nuevacontrasenia"`
	}
	var req reqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	err := s.Svc.RecuperarPassword(req.Correo, req.Token, req.NuevaPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}

// POST /login
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	type reqBody struct {
		Correo      string `json:"correo"`
		Contrasenia string `json:"contrasenia"`
	}
	var req reqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	ok, usuario, err := s.Svc.Login(req.Correo, req.Contrasenia)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"acceso": ok,
	}
	if ok {
		resp["usuario"] = usuario
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
