package Bdoperators

import (
	UserModels "ServicioUsuario/pkg/core/domain"
	"ServicioUsuario/pkg/core/internal"
	"ServicioUsuario/pkg/core/usecase/interfaces"
	"ServicioUsuario/pkg/repository"
	repoInterface "ServicioUsuario/pkg/repository/interfaces"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	crypton "github.com/Shauanth/Singleton_Encription_ServiceGolang/crypton"
	log "github.com/go-kit/log"
)

type ServicioUsuario struct {
	logger      log.Logger
	crud        repoInterface.UserRepository
	cryptConfig crypton.Config
}

// Corrige el tipo de retorno a interfaces.Service
func NuevoServicioUsuario(db *sql.DB, cryptConfig crypton.Config) interfaces.Service {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	crud := repository.NewUserRepository(db)
	return &ServicioUsuario{
		logger:      logger,
		crud:        crud,
		cryptConfig: cryptConfig,
	}
}

// Corrige la firma para que retorne valores
func (s *ServicioUsuario) Get(_ context.Context, filters ...internal.Filter) (internal.StatusCode, []UserModels.UsuarioBD, error) {
	// Implementar lógica de filtrado si es necesario
	// Por ahora retorna todos los usuarios
	return s.SeleccionarUsuarios("", nil)
}

// Corrige la firma para que retorne valores
func (s *ServicioUsuario) Status(_ context.Context, userId string) (internal.StatusCode, error) {
	// Implementar lógica de status si es necesario
	return internal.InProgress, nil
}

func (s *ServicioUsuario) ServiceStatus(_ context.Context) (int, error) {
	s.logger.Log("msg", "Checking status")
	return http.StatusOK, nil
}

func (s *ServicioUsuario) InsertarNuevoUsuario(nombres, apellidos, correo, telefono string, fechaNacimiento time.Time, contrasenia string, rol int) (internal.StatusCode, error) {
	// Encriptar la contraseña antes de guardar
	contraseniaEncriptada, err := crypton.Encrypt(contrasenia, s.cryptConfig)
	if err != nil {
		s.logger.Log("err", fmt.Sprintf("error al encriptar contraseña: %v", err))
		return internal.Error, fmt.Errorf("error al encriptar contraseña: %w", err)
	}
	datos := UserModels.UsuarioVariable{
		Nombres:         nombres,
		Apellidos:       apellidos,
		Correo:          correo,
		Telefono:        telefono,
		FechaNacimiento: fechaNacimiento,
		Contrasenia:     contraseniaEncriptada,
		Rol:             rol,
		EstadoAcceso:    true,
	}
	if err := s.crud.Insertar(`"Usuario"`, datos); err != nil {
		s.logger.Log("err", fmt.Sprintf("error al insertar usuario: %v", err))
		return internal.Error, err
	}
	return internal.InProgress, nil
}

// Corrige la firma y el uso de where
func (s *ServicioUsuario) ActualizarUsuario(idUsuario int, nombres, apellidos, correo, telefono string, fechaNacimiento time.Time, contrasenia string, rol int, estado bool) (internal.StatusCode, error) {
	// Encriptar la contraseña si se proporciona una nueva
	contraseniaEncriptada := contrasenia
	if contrasenia != "" {
		var err error
		contraseniaEncriptada, err = crypton.Encrypt(contrasenia, s.cryptConfig)
		if err != nil {
			s.logger.Log("err", fmt.Sprintf("error al encriptar contraseña: %v", err))
			return internal.Error, fmt.Errorf("error al encriptar contraseña: %w", err)
		}
	}
	datos := UserModels.UsuarioVariable{
		Nombres:         nombres,
		Apellidos:       apellidos,
		Correo:          correo,
		Telefono:        telefono,
		FechaNacimiento: fechaNacimiento,
		Contrasenia:     contraseniaEncriptada,
		Rol:             rol,
		EstadoAcceso:    estado,
	}
	where := "id_usuario = $1"
	if err := s.crud.Actualizar(`"Usuario"`, datos, where, idUsuario); err != nil {
		s.logger.Log("err", fmt.Sprintf("error al actualizar usuario: %v", err))
		return internal.Error, err
	}
	return internal.InProgress, nil
}

// Implementa EliminarUsuario
func (s *ServicioUsuario) EliminarUsuario(id int) (internal.StatusCode, error) {
	if err := s.crud.Eliminar(`"Usuario"`, fmt.Sprintf("%d", id)); err != nil {
		s.logger.Log("err", fmt.Sprintf("error al eliminar usuario: %v", err))
		return internal.Error, err
	}
	return internal.InProgress, nil
}

func (s *ServicioUsuario) SeleccionarUsuarios(condicion string, args ...interface{}) (internal.StatusCode, []UserModels.UsuarioBD, error) {
	var usuarios []UserModels.UsuarioBD
	columnas := []string{
		"id_usuario", "nombres", "apellidos", "correo",
		"telefono", "fechanacimiento", "contrasenia", "rol", "estadoacceso",
	}
	var rows *sql.Rows
	var err error
	if condicion == "" {
		rows, err = s.crud.Seleccionar(`"Usuario"`, columnas, "", args...)
	} else {
		rows, err = s.crud.Seleccionar(`"Usuario"`, columnas, condicion, args...)
	}
	if err != nil {
		s.logger.Log("err", fmt.Sprintf("error en Select: %v", err))
		return internal.Error, nil, fmt.Errorf("error en Select: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var usuario UserModels.UsuarioBD
		err := rows.Scan(
			&usuario.IdUsuario,
			&usuario.DataUsuario.Nombres,
			&usuario.DataUsuario.Apellidos,
			&usuario.DataUsuario.Correo,
			&usuario.DataUsuario.Telefono,
			&usuario.DataUsuario.FechaNacimiento,
			&usuario.DataUsuario.Contrasenia,
			&usuario.DataUsuario.Rol,
			&usuario.DataUsuario.EstadoAcceso,
		)
		if err != nil {
			s.logger.Log("err", fmt.Sprintf("error al escanear fila: %v", err))
			return internal.Error, nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		usuarios = append(usuarios, usuario)
	}
	if err = rows.Err(); err != nil {
		s.logger.Log("err", fmt.Sprintf("error al escanear fila: %v", err))
		return internal.Error, nil, fmt.Errorf("error al escanear fila: %v", err)
	}
	return internal.InProgress, usuarios, nil
}

// Métodos auxiliares para roles (no están en la interfaz principal)
func (s *ServicioUsuario) InsertarNuevoRol(nombreRol string) error {
	datos := UserModels.Rol{
		NombreRol: nombreRol,
	}
	return s.crud.Insertar(`"Roles"`, datos)
}

func (s *ServicioUsuario) ActualizarRol(idRol int, nombreRol string) error {
	datos := UserModels.Rol{
		NombreRol: nombreRol,
	}
	where := "id_rol = $1"
	return s.crud.Actualizar(`"Roles"`, datos, where, idRol)
}

func (s *ServicioUsuario) SeleccionarRoles(condicion string, args ...interface{}) ([]UserModels.RolDB, error) {
	var roles []UserModels.RolDB
	columnas := []string{"id_rol", "nombrerol"}
	var rows *sql.Rows
	var err error
	if condicion == "" {
		rows, err = s.crud.Seleccionar(`"Roles"`, columnas, "", args...)
	} else {
		rows, err = s.crud.Seleccionar(`"Roles"`, columnas, condicion, args...)
	}
	if err != nil {
		s.logger.Log("err", fmt.Sprintf("error en Select: %v", err))
		return nil, fmt.Errorf("error en Select: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var rol UserModels.RolDB
		err := rows.Scan(
			&rol.ID,
			&rol.DataRol.NombreRol,
		)
		if err != nil {
			s.logger.Log("err", fmt.Sprintf("error al escanear fila: %v", err))
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		roles = append(roles, rol)
	}
	if err = rows.Err(); err != nil {
		s.logger.Log("err", fmt.Sprintf("error después de iterar filas: %v", err))
		return nil, fmt.Errorf("error después de iterar filas: %v", err)
	}
	return roles, nil
}
