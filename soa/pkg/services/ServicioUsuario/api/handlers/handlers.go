// Package handlers proporciona los controladores HTTP para la gestión de usuarios y roles.
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	BDoperators "soa/pkg/services/ServicioUsuario/core/usecase/backBD"
	"time"
)

// Server representa el servidor HTTP que maneja las solicitudes relacionadas con usuarios y roles.
type Server struct {
	Svc *BDoperators.ServicioUsuario // Cambiado a puntero para evitar copia innecesaria
}

// NewServer crea una nueva instancia de Server con el servicio de usuario proporcionado.
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
// Se espera que el cuerpo de la petición contenga los campos necesarios para crear un usuario.
// El campo FechaNacimiento debe estar en formato "YYYY-MM-DD".
// Si la inserción es exitosa, se devuelve un JSON con el estado de la operación.
// Si ocurre un error al decodificar el JSON o al insertar el usuario, se devuelve
// un error HTTP 400 o 500 respectivamente.
// La función habilita CORS para permitir solicitudes desde diferentes orígenes.
// Si la petición es OPTIONS, se responde con un estado 200 OK para manejar CORS.
// La función espera un cuerpo de petición en formato JSON con los siguientes campos:
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
// ListarUsuarios maneja la obtención de usuarios a través de una petición HTTP GET.
// Permite filtrar por ID de usuario a través de parámetros de consulta.
// Si se proporciona el parámetro id_usuario, se filtran los usuarios por ese ID.
// Si no se proporciona, se devuelven todos los usuarios.
// La respuesta se devuelve en formato JSON.
// Si ocurre un error al obtener los usuarios, se devuelve un error HTTP 500.
// Si la petición es OPTIONS, se responde con un estado 200 OK para manejar CORS.
// Se espera que la petición tenga un parámetro de consulta opcional "id_usuario" para filtrar los resultados.
// Si el parámetro "id_usuario" está presente, se utiliza para filtrar los usuarios
// y se devuelve una lista de usuarios que coinciden con ese ID.
// Si el parámetro no está presente, se devuelven todos los usuarios.
// La respuesta se envía en formato JSON con el encabezado "Content-Type" establecido a "application/json".
// Si ocurre un error al obtener los usuarios, se devuelve un error HTTP 500 con el
// mensaje de error correspondiente.
// Si la petición es OPTIONS, se responde con un estado 200 OK para manejar CORS
func (s *Server) ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	queryParams := r.URL.Query()
	idUsuario := queryParams.Get("id_usuario")

	var condicion string
	var args []interface{}

	if idUsuario != "" {
		condicion = "id_usuario = $1" // Cambiado ? por $1
		args = append(args, idUsuario)
	}

	_, usuarios, err := s.Svc.SeleccionarUsuarios(condicion, args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// PUT /usuarios/{id}
// ActualizarUsuario maneja la actualización de un usuario existente a través de una petición HTTP PUT.
// Decodifica el cuerpo de la petición, valida los datos y llama al servicio correspondiente.
// Se espera que el cuerpo de la petición contenga los campos necesarios para actualizar el usuario.
// El campo FechaNacimiento debe estar en formato "YYYY-MM-DD".
// El ID del usuario a actualizar se obtiene del cuerpo de la petición.
// Si la actualización es exitosa, se devuelve un JSON con el estado de la operación.
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
// EliminarUsuario maneja la eliminación de un usuario existente a través de una petición HTTP DELETE.
// El ID del usuario a eliminar se obtiene de los parámetros de consulta.
// Si el ID no se proporciona, se devuelve un error HTTP 400.
// Si la eliminación es exitosa, se devuelve un JSON con el estado de la operación.
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
// InsertarRol maneja la creación de un nuevo rol a través de una petición HTTP POST.
// Decodifica el cuerpo de la petición, valida los datos y llama al servicio correspondiente.
// Se espera que el cuerpo de la petición contenga el nombre del rol.
// Si la inserción es exitosa, se devuelve un JSON con el estado de la operación
// Si ocurre un error al decodificar el JSON o al insertar el rol, se devuelve
// un error HTTP 400 o 500 respectivamente.
// La función habilita CORS para permitir solicitudes desde diferentes orígenes.
// Si la petición es OPTIONS, se responde con un estado 200 OK para manejar CORS
// La función espera un cuerpo de petición en formato JSON con el siguiente campo:
// - nombrerol: el nombre del rol a insertar.
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
// EliminarRol maneja la eliminación de un rol existente a través de una petición HTTP DELETE.
// El ID del rol a eliminar se obtiene de los parámetros de consulta.
// Si el ID no se proporciona, se devuelve un error HTTP 400.
// Si la eliminación es exitosa, se devuelve un JSON con el estado de la operación.

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
// ListarRoles maneja la obtención de roles a través de una petición HTTP GET.
// Permite filtrar por nombre de rol a través de parámetros de consulta.
// Si se proporciona el parámetro nombrerol, se filtran los roles por ese nombre.
// Si no se proporciona, se devuelven todos los roles.
// La respuesta se devuelve en formato JSON.
// Si ocurre un error al obtener los roles, se devuelve un error HTTP 500.
// Si la petición es OPTIONS, se responde con un estado 200 OK para manejar CORS.
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
// RecuperarPassword maneja la recuperación de la contraseña a través de una petición HTTP POST.
// Decodifica el cuerpo de la petición, valida los datos y llama al servicio correspondiente.
// Se espera que el cuerpo de la petición contenga el correo, el token de recuperación y
// la nueva contraseña.
// Si la recuperación es exitosa, se devuelve un JSON con el estado de la operación.
// Si ocurre un error al decodificar el JSON o al recuperar la contraseña, se devuelve
// un error HTTP 400 o 500 respectivamente.
// La función habilita CORS para permitir solicitudes desde diferentes orígenes.
// Si la petición es OPTIONS, se responde con un estado 200 OK para manejar CORS
// La función espera un cuerpo de petición en formato JSON con los siguientes campos:
// - correo: el correo electrónico del usuario.
// - token: el token de recuperación de contraseña.
// - nuevacontrasenia: la nueva contraseña que se desea establecer.
// Si la recuperación es exitosa, se devuelve un JSON con el estado "ok"
// Si la recuperación falla, se devuelve un error HTTP 400 con el mensaje de error correspondiente.
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
// Login maneja el inicio de sesión de un usuario a través de una petición HTTP POST.
// Decodifica el cuerpo de la petición, valida los datos y llama al servicio correspondiente.
// Se espera que el cuerpo de la petición contenga el correo y la contraseña del usuario.
// Si el inicio de sesión es exitoso, se devuelve un JSON con el estado de acceso
// y los datos del usuario.
// Si ocurre un error al decodificar el JSON o al iniciar sesión, se devuelve
// un error HTTP 400 o 500 respectivamente.
// La función habilita CORS para permitir solicitudes desde diferentes orígenes.
// Si la petición es OPTIONS, se responde con un estado 200 OK para manejar CORS.
// La función espera un cuerpo de petición en formato JSON con los siguientes campos:
// - correo: el correo electrónico del usuario.
// - contrasenia: la contraseña del usuario.
// Si el inicio de sesión es exitoso, se devuelve un JSON con los siguientes campos:
// - acceso: un booleano que indica si el inicio de sesión fue exitoso.
// - usuario: un objeto que contiene los datos del usuario si el inicio de sesión fue exitoso.
// Si el inicio de sesión falla, se devuelve un JSON con acceso: false.
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

// POST /recuperar/verificar
// VerificarTokenRecuperacion maneja la verificación del token de recuperación de contraseña a través de una petición HTTP POST.
// Decodifica el cuerpo de la petición, valida los datos y llama al servicio correspondiente.
// Se espera que el cuerpo de la petición contenga el correo y el token de recuperación.
// Si la verificación es exitosa, se devuelve un JSON con el estado de la operación
// Si ocurre un error al decodificar el JSON o al verificar el token, se devuelve
// un error HTTP 400 o 500 respectivamente.
// La función habilita CORS para permitir solicitudes desde diferentes orígenes.
// Si la petición es OPTIONS, se responde con un estado 200 OK para manejar CORS
// La función espera un cuerpo de petición en formato JSON con los siguientes campos:
// - correo: el correo electrónico del usuario.
// - token: el token de recuperación de contraseña.
// Si la verificación es exitosa, se devuelve un JSON con el estado "ok"
// Si la verificación falla, se devuelve un error HTTP 400 con el mensaje de error correspondiente.
// Si la verificación es exitosa, se devuelve un JSON con el estado "ok".
func (s *Server) VerificarTokenRecuperacion(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	type reqBody struct {
		Correo string `json:"correo"`
		Token  string `json:"token"`
	}
	var req reqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	err := s.Svc.VerificarTokenRecuperacion(req.Correo, req.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}

// POST /recuperar/actualizar
// ActualizarPasswordRecuperacion maneja la actualización de la contraseña de recuperación a través de una petición HTTP POST.
// Decodifica el cuerpo de la petición, valida los datos y llama al servicio correspondiente.
// Se espera que el cuerpo de la petición contenga el correo y la nueva contraseña.
// Si la actualización es exitosa, se devuelve un JSON con el estado de la operación
// Si ocurre un error al decodificar el JSON o al actualizar la contraseña, se devuelve
// un error HTTP 400 o 500 respectivamente.
// La función habilita CORS para permitir solicitudes desde diferentes orígenes.
// Si la petición es OPTIONS, se responde con un estado 200 OK para manejar CORS
// La función espera un cuerpo de petición en formato JSON con los siguientes campos:
// - correo: el correo electrónico del usuario.
// - nuevacontrasenia: la nueva contraseña que se desea establecer.
// Si la actualización es exitosa, se devuelve un JSON con el estado "ok"
// Si la actualización falla, se devuelve un error HTTP 400 con el mensaje de error correspondiente.
// Si la actualización es exitosa, se devuelve un JSON con el estado "ok".
func (s *Server) ActualizarPasswordRecuperacion(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	type reqBody struct {
		Correo        string `json:"correo"`
		NuevaPassword string `json:"nuevacontrasenia"`
	}
	var req reqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	err := s.Svc.ActualizarPasswordRecuperacion(req.Correo, req.NuevaPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
}
