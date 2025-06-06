package UserModels

import "time"

type Usuario struct {
	IdUsuario       int       `db:"id_usuario"`
	Nombres         string    `db:"nombres"`
	Apellidos       string    `db:"apellidos"`
	Correo          string    `db:"correo"`
	FechaNacimiento time.Time `db:"fechanacimiento"`
	Contrasenia     string    `db:"contrasenia"`
	Rol             int       `db:"rol"`
	EstadoAcceso    bool      `db:"estadoacceso"`
}
