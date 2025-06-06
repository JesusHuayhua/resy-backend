package bduser

import (
	"RESI-BACKEND/services/shared/utils"
	"RESI-BACKEND/services/users/pkg/models/UserModels"
	"database/sql"
	"fmt"
	"time"
)

func InsertarNuevoUsuario(db *sql.DB, nombres string, apellidos string, correo string, fechaNacimiento time.Time, contrasenia string, rol int) error {
	crud := utils.NuevoCRUD(db)
	datos := UserModels.Usuario{
		Nombres:         nombres,
		Apellidos:       apellidos,
		Correo:          correo,
		FechaNacimiento: fechaNacimiento,
		Contrasenia:     contrasenia,
		Rol:             rol,
	}
	return crud.Insertar(`"Usuario"`, datos)
}

func ActualizarUsuario(db *sql.DB, idUsuario int, nombres string, apellidos string) error {
	crud := utils.NuevoCRUD(db)
	datos := UserModels.Usuario{
		Nombres:   nombres,
		Apellidos: apellidos,
	}
	where := "id_usuario = $3"
	return crud.Actualizar(`"Usuario"`, datos, where, idUsuario)
}

func SeleccionarUsuarios(db *sql.DB, condicion string, args ...interface{}) ([]UserModels.Usuario, error) {
	crud := utils.NuevoCRUD(db)
	var usuarios []UserModels.Usuario

	columnas := []string{
		"id_usuario",
		"nombres",
		"apellidos",
		"correo",
		"fecha_nacimiento",
		"contrasenia",
		"id_rol",
		"nombre_rol",
		"estado_acceso",
	}

	var rows *sql.Rows
	var err error

	if condicion == "" {
		rows, err = crud.Select("usuario", columnas, "", args...)
	} else {
		rows, err = crud.Select("usuario", columnas, condicion, args...)
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
