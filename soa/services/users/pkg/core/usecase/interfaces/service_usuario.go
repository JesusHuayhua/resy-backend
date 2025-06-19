package interfaces

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	UserModels "soa/services/users/pkg/core/domain"
	"soa/services/users/pkg/core/response"
	"soa/services/users/pkg/core/svc_internal"
	"soa/services/users/pkg/core/usecase"
	"soa/services/users/pkg/repository"
	repoInterface "soa/services/users/pkg/repository/interfaces"
	"strconv"
	"time"

	log "github.com/go-kit/log"
)

type ServicioUsuario struct {
}

var (
	logger log.Logger
	crud   repoInterface.UserRepository
)

type UserService interface {
	usecase.Service
	Usuario(ctx context.Context, tipoOP int, args []svc_internal.Filter) response.GenericResponse
	Roles(ctx context.Context, tipoOP int, args []svc_internal.Filter) response.GenericResponse
}

func NuevoServicio() UserService {
	return &ServicioUsuario{}
}
func (s1 *ServicioUsuario) Get(_ context.Context, filters ...svc_internal.Filter) error {
	return nil
}

func (s1 *ServicioUsuario) Status(_ context.Context, userId string) (svc_internal.StatusCode, error) {
	return svc_internal.InProgress, fmt.Errorf("None")
}

func (s1 *ServicioUsuario) ServiceStatus(_ context.Context) (int, error) {
	logger.Log("Checking status")
	return http.StatusOK, nil
}

func valor(args []svc_internal.Filter, key string) (string, bool) {
	for _, f := range args {
		if f.Key == key {
			return f.Value, true
		}
	}
	return "", false
}

func InsertarUsuario(args []svc_internal.Filter) (int, string) {
	if len(args) != 7 {
		return http.StatusNotAcceptable, "[INSERTAR] Error insertando usuario con los parametros"
	}
	nombres, _ := valor(args, "nombres")
	apellidos, _ := valor(args, "apellidos")
	correo, _ := valor(args, "correo")
	telefono, _ := valor(args, "telefono")
	nacStr, _ := valor(args, "fechaNacimiento")
	pass, _ := valor(args, "contraseña")
	rolStr, _ := valor(args, "rol")

	d, _ := time.Parse("02/01/2006", nacStr)
	r, _ := strconv.Atoi(rolStr)
	datos := UserModels.UsuarioVariable{
		Nombres:         nombres,
		Apellidos:       apellidos,
		Correo:          correo,
		Telefono:        telefono,
		FechaNacimiento: d,
		Contrasenia:     pass,
		Rol:             r,
		EstadoAcceso:    true,
	}
	err := crud.Insertar(`"Usuario"`, datos)
	return http.StatusOK, err.Error()
}

func ActualizarUsuario(args []svc_internal.Filter) (int, string) {
	if len(args) != 7 {
		return http.StatusNotAcceptable, "[ACTUALIZAR] Error en los parametros pasados"
	}

	id, _ := valor(args, "idUsuario")
	nombres, _ := valor(args, "nombres")
	apellidos, _ := valor(args, "apellidos")
	correo, _ := valor(args, "correo")
	telefono, _ := valor(args, "telefono")
	fechaNac, _ := valor(args, "fechaNacimiento")
	pass, _ := valor(args, "contraseña")
	rolStr, _ := valor(args, "rol")
	stat, _ := valor(args, "estado")

	idUsuario, _ := strconv.Atoi(id)
	rol, _ := strconv.Atoi(rolStr)
	d, _ := time.Parse("02/01/2006", fechaNac)
	estado, _ := strconv.Atoi(stat)
	datos := UserModels.UsuarioVariable{
		Nombres:         nombres,
		Apellidos:       apellidos,
		Correo:          correo,
		Telefono:        telefono,
		FechaNacimiento: d,
		Contrasenia:     pass,
		Rol:             rol,
		EstadoAcceso:    estado != 0,
	}
	where := "id_usuario = $9"
	err := crud.Actualizar(`"Usuario"`, datos, where, idUsuario)
	return http.StatusOK, err.Error()
}

func SeleccionarUsuarios(args []svc_internal.Filter) ([]UserModels.UsuarioVariable, int) {
	var usuarios []UserModels.UsuarioVariable
	columnas := []string{"condicion", "id_usuario", "nombres", "apellidos", "correo", "telefono",
		"fechanacimiento", "contrasenia", "rol", "estadoacceso",
	}
	var rows *sql.Rows
	var err error
	args = reorderFilters(args, columnas)
	condicion := args[0].Value
	if condicion == "" {
		rows, err = crud.Seleccionar(`"Usuario"`, columnas, "", args[1:])
	} else {
		rows, err = crud.Seleccionar(`"Usuario"`, columnas, condicion, args[1:])
	}
	if err != nil {
		return nil, http.StatusNotAcceptable
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
			return nil, http.StatusNotAcceptable
		}
		usuarios = append(usuarios, usuario.U)
	}
	if err = rows.Err(); err != nil {
		return nil, http.StatusNotAcceptable
	}
	return usuarios, http.StatusNotAcceptable
}

func InsertarNuevoRol(args []svc_internal.Filter) (int, string) {
	if len(args) != 1 {
		return http.StatusNotAcceptable, "[ROL_INSERTAR] No se pudo insertar rol por argumentos invalidos"
	}
	datos := UserModels.Rol{
		NombreRol: args[0].Value,
	}
	err := crud.Insertar(`"Roles"`, datos)
	return http.StatusOK, err.Error()
}

func reorderFilters(args []svc_internal.Filter, keyOrder []string) []svc_internal.Filter {
	grouped := make(map[string][]svc_internal.Filter, len(args))
	for _, f := range args {
		grouped[f.Key] = append(grouped[f.Key], f)
	}
	ordered := make([]svc_internal.Filter, 0, len(args))
	for _, key := range keyOrder {
		if fs, ok := grouped[key]; ok {
			ordered = append(ordered, fs...)
			delete(grouped, key)
		}
	}
	for _, f := range args {
		if fs, ok := grouped[f.Key]; ok && len(fs) > 0 {
			ordered = append(ordered, fs[0])
			grouped[f.Key] = fs[1:]
		}
	}
	return ordered
}

func SeleccionarRoles(args []svc_internal.Filter) ([]UserModels.RolDB, int) {
	var roles []UserModels.RolDB
	columnas := []string{"condicion", "id_rol", "nombrerol"}
	var rows *sql.Rows
	var err error
	args = reorderFilters(args, columnas)
	condicion := args[0].Value
	if condicion == "" {
		rows, err = crud.Seleccionar(`"Roles"`, columnas, "", args[1:])
	} else {
		rows, err = crud.Seleccionar(`"Roles"`, columnas, condicion, args[1:])
	}
	if err != nil {
		return nil, http.StatusNotAcceptable
	}
	defer rows.Close()

	for rows.Next() {
		var rol UserModels.RolDB
		err := rows.Scan(
			&rol.ID,
			&rol.NombreRol,
		)
		if err != nil {
			return nil, http.StatusNotAcceptable
		}
		roles = append(roles, rol)
	}
	if err = rows.Err(); err != nil {
		return nil, http.StatusNotAcceptable
	}
	return roles, http.StatusOK
}

func (s1 *ServicioUsuario) Roles(ctx context.Context, tipoOP int, args []svc_internal.Filter) response.GenericResponse {
	logger.Log("[User] Parseando informacion")
	switch tipo := tipoOP; tipo {
	case 1:
		{
			int_code, status := InsertarNuevoRol(args)
			return response.GenericResponse{Code: int_code, Data: status}
		}
	case 2:
		{
			roles, status := SeleccionarRoles(args)
			return response.GenericResponse{Code: status, Data: roles}
		}
	}
	logger.Log("[User] Insertado")
	return response.GenericResponse{Code: int(svc_internal.Error), Data: "[ERROR] Invalid service"}
}

func (s1 *ServicioUsuario) Usuario(_ context.Context, tipoOP int, args []svc_internal.Filter) response.GenericResponse {
	logger.Log("[User] Parseando informacion")
	switch tipo := tipoOP; tipo {
	case 1:
		{
			int_code, status := InsertarUsuario(args)
			return response.GenericResponse{Code: int_code, Data: status}
		}
	case 2:
		{
			int_code, status := ActualizarUsuario(args)
			return response.GenericResponse{Code: int_code, Data: status}
		}
	case 3:
		{
			usuarios, status := SeleccionarUsuarios(args)
			return response.GenericResponse{Code: status, Data: usuarios}
		}
	}
	return response.GenericResponse{Code: int(svc_internal.Error), Err: "[ERROR] Invalid service"}
}

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	crud = repository.NuevoCRUD()
}
