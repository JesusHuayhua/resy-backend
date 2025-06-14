package UserModels

import "time"

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
type UsuarioBD struct {
	IdUsuario int `db:"id_usuario"`
	usuario   UsuarioVariable
}
