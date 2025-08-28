package UserModels

import "time"

type UsuarioVariable struct {
	Nombres         string    `db:"nombres"`
	Apellidos       string    `db:"apellidos"`
	Correo          string    `db:"correo"`
	Telefono        string    `db:"telefono"`
	Direccion       string    `db:"direccion"`
	FechaNacimiento time.Time `db:"fechanacimiento"`
	Contrasenia     string    `db:"contrasenia"`
	Rol             int       `db:"rol"`
	EstadoAcceso    bool      `db:"estadoacceso"`
}
type UsuarioBD struct {
	IdUsuario   int `db:"id_usuario"`
	DataUsuario UsuarioVariable
}
