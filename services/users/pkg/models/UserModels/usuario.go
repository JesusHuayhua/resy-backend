package UserModels

import (
	"time"
)

type Usuario struct {
	ID              int       `db:"id_usuario"`
	Nombres         string    `db:"nombres"`
	Apellidos       string    `db:"apellidos"`
	Correo          string    `db:"correo"`
	FechaNacimiento time.Time `db:"fechanacimiento"`
	Contrasenia     string    `db:"contrasenia"`
	RolID           int       `db:"rol"`
	EstadoAcceso    bool      `db:"estadoacceso"`
}
