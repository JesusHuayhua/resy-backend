package bduser

import (
	UserModels "ServicioUsuario/dominio"
	"database/sql"
	"fmt"
	"time"

	"github.com/Shauanth/Singleton_Encription_ServiceGolang/database"
)

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

func InsertarNuevoUsuario(db *sql.DB, nombres string, apellidos string, correo string, telefono string, fechaNacimiento time.Time, contrasenia string, rol int) error {
	crud := database.NuevoCRUD(db)
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
	return crud.Insertar(`"Usuario"`, datos)
}

func ActualizarUsuario(db *sql.DB, idUsuario int, nombres string, apellidos string, correo string, telefono string, fechaNacimiento time.Time, contrasenia string, rol int, estado bool) error {
	crud := database.NuevoCRUD(db)
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
	where := "id_usuario = $9"
	return crud.Actualizar(`"Usuario"`, datos, where, idUsuario)
}

func SeleccionarUsuarios(db *sql.DB, condicion string, args ...interface{}) ([]UserModels.Usuario, error) {
	crud := database.NuevoCRUD(db)
	var usuarios []UserModels.Usuario

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
		return nil, fmt.Errorf("error en Select: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var usuario UserModels.Usuario
		err := rows.Scan(
			&usuario.IdUsuario,
			&usuario.Nombres,
			&usuario.Apellidos,
			&usuario.Correo,
			&usuario.Telefono,
			&usuario.FechaNacimiento,
			&usuario.Contrasenia,
			&usuario.Rol,
			&usuario.EstadoAcceso,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		usuarios = append(usuarios, usuario)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error después de iterar filas: %v", err)
	}
	return usuarios, nil
}
