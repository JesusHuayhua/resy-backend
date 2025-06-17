package usecase

import (
	"ServicioUsuario/pkg/core/internal"
	"context"
)

type Service interface {
	Get(ctx context.Context, filters ...internal.Filter)
	Status(ctx context.Context, userId string)
	ServiceStatus(ctx context.Context) (int, error)
	InsertarNuevoUsuario(nombres, apellidos, correo, telefono string, fechaNacimiento internal.DateTime, contrasenia string, rol int) (internal.StatusCode, error)
	ActualizarUsuario(id int, nombres, apellidos, correo, telefono string, fechaNacimiento internal.DateTime, contrasenia string, rol int) (internal.StatusCode, error)
	EliminarUsuario(id int) (internal.StatusCode, error)
}
