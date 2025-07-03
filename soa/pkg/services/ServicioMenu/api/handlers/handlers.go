package handlers

import (
	"encoding/json"
	"net/http"
	"soa/pkg/services/ServicioMenu/core/usecase/backBD"
	"strconv"
)

type Server struct {
	Svc *backBD.ServicioMenu
}

func NewServer(svc *backBD.ServicioMenu) *Server {
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

// PLATOS
func (s *Server) InsertarPlato(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	type reqBody struct {
		NombrePlato string  `json:"nombre_plato"`
		Categoria   int     `json:"categoria"`
		Descripcion string  `json:"descripcion"`
		Precio      float64 `json:"precio"`
		Imagen      string  `json:"imagen"`
		Estado      bool    `json:"estado"`
	}
	var req reqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	err := s.Svc.InsertarPlato(req.NombrePlato, req.Categoria, req.Descripcion, req.Precio, req.Imagen, req.Estado)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}

func (s *Server) ListarPlatos(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	platos, err := s.Svc.SeleccionarPlatos("", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(platos)
}

func (s *Server) ActualizarPlato(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	type reqBody struct {
		IDPlato     int     `json:"id_plato"`
		NombrePlato string  `json:"nombre_plato"`
		Categoria   int     `json:"categoria"`
		Descripcion string  `json:"descripcion"`
		Precio      float64 `json:"precio"`
		Imagen      string  `json:"imagen"`
		Estado      bool    `json:"estado"`
	}
	var req reqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	err := s.Svc.ActualizarPlato(req.IDPlato, req.NombrePlato, req.Categoria, req.Descripcion, req.Precio, req.Imagen, req.Estado)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}

func (s *Server) EliminarPlato(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	id := r.URL.Query().Get("id")
	idInt, _ := strconv.Atoi(id)
	err := s.Svc.EliminarPlato(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}

// CATEGORIAS
func (s *Server) InsertarCategoria(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	type reqBody struct {
		Nombre string `json:"nombre"`
	}
	var req reqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	err := s.Svc.InsertarCategoria(req.Nombre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}

func (s *Server) ListarCategorias(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	cats, err := s.Svc.SeleccionarCategorias()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cats)
}

func (s *Server) ActualizarCategoria(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	type reqBody struct {
		ID     int    `json:"id"`
		Nombre string `json:"nombre"`
	}
	var req reqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	err := s.Svc.ActualizarCategoria(req.ID, req.Nombre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}

func (s *Server) EliminarCategoria(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	id := r.URL.Query().Get("id")
	idInt, _ := strconv.Atoi(id)
	err := s.Svc.EliminarCategoria(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}

// MENÚ SEMANAL
func (s *Server) InsertarMenuSemanal(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	type reqBody struct {
		IDMenu        string `json:"id_menu"`
		FechaDeInicio string `json:"fecha_inicio"`
		FechaFin      string `json:"fecha_fin"`
	}
	var req reqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	err := s.Svc.InsertarMenuSemanal(req.IDMenu, req.FechaDeInicio, req.FechaFin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}

func (s *Server) ListarMenusSemanales(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	menus, err := s.Svc.ListarMenusSemanales()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(menus)
}

func (s *Server) InsertarMenudia(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	type reqBody struct {
		IDMenu    string `json:"id_menu"`
		DiaSemana string `json:"dia_semana"`
	}
	var req reqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	err := s.Svc.InsertarMenudia(req.IDMenu, req.DiaSemana)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}

func (s *Server) ListarDiasDeMenu(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	idMenu := r.URL.Query().Get("id_menu")
	dias, err := s.Svc.ListarDiasDeMenu(idMenu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(dias)
}

func (s *Server) InsertarPlatoEnMenudia(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	type reqBody struct {
		IDDia      int  `json:"id_dia"`
		IDPlato    int  `json:"id_plato"`
		Cantidad   int  `json:"cantidad"`
		Disponible bool `json:"disponible"`
	}
	var req reqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	err := s.Svc.InsertarPlatoEnMenudia(req.IDDia, req.IDPlato, req.Cantidad, req.Disponible)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}

func (s *Server) ListarPlatosEnMenudia(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	idDia, _ := strconv.Atoi(r.URL.Query().Get("id_dia"))
	platos, err := s.Svc.ListarPlatosEnMenudia(idDia)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(platos)
}

// GET /menu/completo?id_menu=SEMANA_X
func (s *Server) ObtenerMenuSemanalCompleto(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	idMenu := r.URL.Query().Get("id_menu")
	if idMenu == "" {
		http.Error(w, "id_menu es requerido", http.StatusBadRequest)
		return
	}
	menu, err := s.Svc.ObtenerMenuSemanalCompleto(idMenu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menu)
}
