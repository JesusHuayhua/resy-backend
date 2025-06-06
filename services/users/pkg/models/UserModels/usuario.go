package UserModels

import (
	"time"
)

type Usuario struct {
	ID              int       `json:"id_usuario"`
	Nombres         string    `json:"nombres"`
	Apellidos       string    `json:"apellidos"`
	Correo          string    `json:"correo"`
	FechaNacimiento time.Time `json:"fecha_nacimiento"`
	Contrasenia     string    `json:"contrasenia"`
	RolID           int       `json:"rol_id"`
	EstadoAcceso    bool      `json:"estado_acceso"`
}
