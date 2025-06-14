package interfaces

import (
	"context"
	"fmt"
	"net/http"
	"os"
	UserModels "soa/services/users/cmd/core/domain"
	"soa/services/users/cmd/core/internal"
	"time"

	log "github.com/go-kit/log"
)

type RolVariable struct {
	NombreRol string `db:"nombrerol"`
}

type UsuarioVariable struct {
	Nombres         string    `db:"nombres"`
	Apellidos       string    `db:"apellidos"`
	Correo          string    `db:"correo"`
	Telefono        string    `db:"telefono"`
	FechaNacimiento time.Time `db:"fechanacimiento"`
	Contrasenia     string    `db:"contrasenia"`
	Rol             int       `db:"rol"`
	EstadoAcceso    bool      `db:"estadoacceso"`
}

type ServicioUsuario struct{}

var logger log.Logger

func nuevoServicio() core.Service {
	return &ServicioUsuario{}
}

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
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
	//sql.DB
	//crud := database.NuevoCRUD(db)
	datos := UsuarioVariable{
		Nombres:         nombres,
		Apellidos:       apellidos,
		Correo:          correo,
		Telefono:        telefono,
		FechaNacimiento: fechaNacimiento,
		Contrasenia:     contrasenia,
		Rol:             rol,
		EstadoAcceso:    true,
	}
	datos.Contrasenia = "aea"
	//return crud.Insertar(`"Usuario"`, datos)
	return internal.InProgress, nil
}

func (s4 *ServicioUsuario) ActualizarUsuario(idUsuario int, nombres string, apellidos string, correo string, telefono string, fechaNacimiento time.Time, contrasenia string, rol int, estado bool) (internal.StatusCode, error) {
	//crud := database.NuevoCRUD(db)
	datos := UsuarioVariable{
		Nombres:         nombres,
		Apellidos:       apellidos,
		Correo:          correo,
		Telefono:        telefono,
		FechaNacimiento: fechaNacimiento,
		Contrasenia:     contrasenia,
		Rol:             rol,
		EstadoAcceso:    estado,
	}
	//where := "id_usuario = $9"
	//return crud.Actualizar(`"Usuario"`, datos, where, idUsuario)
	datos.Contrasenia = "12"
	return internal.InProgress, nil
}

func (s5 *ServicioUsuario) SeleccionarUsuarios(condicion string, args ...interface{}) (internal.StatusCode, []UserModels.Usuario, error) {
	//crud := database.NuevoCRUD(db)
	var usuarios []UserModels.Usuario
	columnas := []string{"id_usuario", "nombres", "apellidos", "correo", "telefono",
		"fechanacimiento", "contrasenia", "rol", "estadoacceso",
	}
	//var rows *sql.Rows
	var err error
	if condicion == "" {
		//	rows, err = crud.Seleccionar(`"Usuario"`, columnas, "", args...)
	} else {
		//	rows, err = crud.Seleccionar(`"Usuario"`, columnas, condicion, args...)
	}
	if err != nil {
		//	return nil, fmt.Errorf("error en Select: %v", err)
	}
	//defer rows.Close()
	for rows.Next() {
		var usuario UserModels.Usuario
		err := rows.Scan(&usuario.IdUsuario, &usuario.Nombres, &usuario.Apellidos, &usuario.Correo, &usuario.Telefono, &usuario.FechaNacimiento,
			&usuario.Contrasenia,
			&usuario.Rol,
			&usuario.EstadoAcceso,
		)
		if err != nil {
			//return nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		usuarios = append(usuarios, usuario)
	}
	if err = rows.Err(); err != nil {
		return internal.Error, nil, fmt.Errorf("error al escanear fila: %v", err)
	}
	return internal.InProgress, usuarios, nil
}

func (s6 *ServicioUsuario) InsertarNuevoRol(nombreRol string) {
	//crud := database.NuevoCRUD(db)
	datos := UserModels.Rol{
		NombreRol: nombreRol,
	}
	return crud.Insertar(`"Roles"`, datos)
}

func (s7 *ServicioUsuario) ActualizarRol(idRol int, nombreRol string) {
	//crud := database.NuevoCRUD(db)
	datos := RolVariable{
		NombreRol: nombreRol,
	}
	where := "id_rol = $2"
	//crud.Actualizar(`"Roles"`, datos, where, idRol)
}

func (s8 *ServicioUsuario) SeleccionarRoles(condicion string, args ...interface{}) ([]UserModels.Rol, error) {
	//crud := database.NuevoCRUD(db)
	var roles []UserModels.Rol

	columnas := []string{"id_rol", "nombrerol"}
	//var rows *sql.Rows
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
		var rol UserModels.Rol
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
