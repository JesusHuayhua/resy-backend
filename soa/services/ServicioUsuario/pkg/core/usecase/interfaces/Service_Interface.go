package interfaces

import (
	UserModels "ServicioUsuario/pkg/core/domain"
	"ServicioUsuario/pkg/core/internal"
	"context"
	"time"
)

type Service interface {
	Get(ctx context.Context, filters ...internal.Filter) (internal.StatusCode, []UserModels.UsuarioBD, error)
	Status(ctx context.Context, userId string) (internal.StatusCode, error)
	ServiceStatus(ctx context.Context) (int, error)
	InsertarNuevoUsuario(nombres, apellidos, correo, telefono string, fechaNacimiento time.Time, contrasenia string, rol int) (internal.StatusCode, error)
	ActualizarUsuario(id int, nombres, apellidos, correo, telefono string, fechaNacimiento time.Time, contrasenia string, rol int, estado bool) (internal.StatusCode, error)
	EliminarUsuario(id int) (internal.StatusCode, error)
	InsertarNuevoRol(nombre string) (internal.StatusCode, error)
	ActualizarRol(id int, nombre string) (internal.StatusCode, error)
	EliminarRol(id int) (internal.StatusCode, error)
	SeleccionarRoles(filtro string, params []interface{}) (internal.StatusCode, []UserModels.RolDB, error)
	IniciarRecuperacionPassword(correo string) (string, error)
	RecuperarPassword(correo, token, nuevaContrasenia string) error
}
