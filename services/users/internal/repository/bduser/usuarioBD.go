package bduser

import (
	"RESI-BACKEND/services/shared/utils"
	"RESI-BACKEND/services/users/pkg/models/UserModels"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
)

func InsertNewUser(db *sql.DB, nombres string, apellidos string, correo string, fechaNacimiento time.Time, contrasenia string, rolID int) error {
	crud := utils.NewCRUD(db)
	data := UserModels.Usuario{
		Nombres:         nombres,
		Apellidos:       apellidos,
		Correo:          correo,
		FechaNacimiento: fechaNacimiento,
		Contrasenia:     contrasenia,
		RolID:           rolID,
	}

	if err := crud.Insert(`"Usuario"`, data); err != nil {
		log.Printf("Error inserting user data: %v", err)
		return err
	}
	return nil
}

func UpdateUser(db *sql.DB, id int, nombres string, apellidos string, correo string, fechaNacimiento time.Time, contrasenia string, rolID int, nombreRol string, estadoAcceso bool) error {
	crud := utils.NewCRUD(db)

	data := UserModels.Usuario{
		ID:              id,
		Nombres:         nombres,
		Apellidos:       apellidos,
		Correo:          correo,
		FechaNacimiento: fechaNacimiento,
		Contrasenia:     contrasenia,
		RolID:           rolID,
		EstadoAcceso:    estadoAcceso,
	}

	if err := crud.Update("usuario", strconv.FormatInt(int64(id), 10), data); err != nil {
		log.Printf("Error updating user data: %v", err)
		return err
	}
	return nil
}

func SelectUsers(db *sql.DB, condicion string, args ...interface{}) ([]UserModels.Usuario, error) {
	crud := utils.NewCRUD(db)
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
			&usuario.ID,
			&usuario.Nombres,
			&usuario.Apellidos,
			&usuario.Correo,
			&usuario.FechaNacimiento,
			&usuario.Contrasenia,
			&usuario.RolID,
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
