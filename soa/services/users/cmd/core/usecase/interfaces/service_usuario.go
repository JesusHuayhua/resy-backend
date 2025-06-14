package interfaces

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	UserModels "soa/services/users/cmd/core/domain"
	"soa/services/users/cmd/core/internal"
	"soa/services/users/cmd/core/usecase"
	"soa/services/users/cmd/repository"
	repoInterface "soa/services/users/cmd/repository/interfaces"
	"time"

	log "github.com/go-kit/log"
)

type ServicioUsuario struct{}

var (
	logger log.Logger
	crud   repoInterface.UserRepository
)

func nuevoServicio() usecase.Service {
	return &ServicioUsuario{}
}

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	crud = repository.NuevoCRUD()
}

func (s1 *ServicioUsuario) Get(_ context.Context, userId string) {

}

func (s2 *ServicioUsuario) Status(_ context.Context, userId string) {

}

func (s3 *ServicioUsuario) ServiceStatus(_ context.Context) (int, error) {
	logger.Log("Checking status")
	return http.StatusOK, nil
}

func (s4 *ServicioUsuario) InsertarNuevoUsuario(nombres string, apellidos string, correo string, telefono string, fechaNacimiento time.Time, contrasenia string, rol int) (internal.StatusCode, error) {
	datos := UserModels.UsuarioVariable{
		Nombres:         nombres,
		Apellidos:       apellidos,
		Correo:          correo,
		Telefono:        telefono,
		FechaNacimiento: fechaNacimiento,
		Contrasenia:     contrasenia,
		Rol:             rol,
		EstadoAcceso:    true,
	}
	return internal.InProgress, crud.Insertar(`"Usuario"`, datos)
}

func (s4 *ServicioUsuario) ActualizarUsuario(idUsuario int, nombres string, apellidos string, correo string, telefono string, fechaNacimiento time.Time, contrasenia string, rol int, estado bool) (internal.StatusCode, error) {
	datos := UserModels.UsuarioVariable{
		Nombres:         nombres,
		Apellidos:       apellidos,
		Correo:          correo,
		Telefono:        telefono,
		FechaNacimiento: fechaNacimiento,
		Contrasenia:     contrasenia,
		Rol:             rol,
		EstadoAcceso:    estado,
	}
	where := "id_usuario = $9"
	return internal.InProgress, crud.Actualizar(`"Usuario"`, datos, where, idUsuario)
}

func (s5 *ServicioUsuario) SeleccionarUsuarios(condicion string, args ...interface{}) (internal.StatusCode, []UserModels.UsuarioVariable, error) {
	var usuarios []UserModels.UsuarioVariable
	columnas := []string{"id_usuario", "nombres", "apellidos", "correo", "telefono",
		"fechanacimiento", "contrasenia", "rol", "estadoacceso",
	}
	var rows *sql.Rows
	var err error
	if condicion == "" {
		rows, err = crud.Seleccionar(`"Usuario"`, columnas, "", args...)
	} else {
		rows, err = crud.Seleccionar(`"Usuario"`, columnas, condicion, args...)
	}
	if err != nil {
		return internal.Error, nil, fmt.Errorf("error en Select: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var usuario UserModels.UsuarioBD
		err := rows.Scan(&usuario.IdUsuario,
			&usuario.U.Nombres,
			&usuario.U.Apellidos,
			&usuario.U.Correo, &usuario.U.Telefono, &usuario.U.FechaNacimiento,
			&usuario.U.Contrasenia,
			&usuario.U.Rol,
			&usuario.U.EstadoAcceso,
		)
		if err != nil {
			return internal.Error, nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		usuarios = append(usuarios, usuario.U)
	}
	if err = rows.Err(); err != nil {
		return internal.Error, nil, fmt.Errorf("error al escanear fila: %v", err)
	}
	return internal.InProgress, usuarios, nil
}

func (s6 *ServicioUsuario) InsertarNuevoRol(nombreRol string) {
	datos := UserModels.Rol{
		NombreRol: nombreRol,
	}
	crud.Insertar(`"Roles"`, datos)
}

func (s7 *ServicioUsuario) ActualizarRol(idRol int, nombreRol string) {
	datos := UserModels.Rol{
		NombreRol: nombreRol,
	}
	where := "id_rol = $2"
	crud.Actualizar(`"Roles"`, datos, where, idRol)
}

func (s8 *ServicioUsuario) SeleccionarRoles(condicion string, args ...interface{}) ([]UserModels.RolDB, error) {
	var roles []UserModels.RolDB
	columnas := []string{"id_rol", "nombrerol"}
	var rows *sql.Rows
	var err error
	if condicion == "" {
		rows, err = crud.Seleccionar(`"Roles"`, columnas, "", args...)
	} else {
		rows, err = crud.Seleccionar(`"Roles"`, columnas, condicion, args...)
	}
	if err != nil {
		return nil, fmt.Errorf("error en Select: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var rol UserModels.RolDB
		err := rows.Scan(
			&rol.ID,
			&rol.NombreRol,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		roles = append(roles, rol)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error después de iterar filas: %v", err)
	}
	return roles, nil
}
