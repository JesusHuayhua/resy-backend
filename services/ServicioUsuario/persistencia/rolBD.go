package bduser

import (
	UserModels "ServicioUsuario/dominio"
	"database/sql"
	"fmt"

	"github.com/Shauanth/Singleton_Encription_ServiceGolang/database"
)

type RolVariable struct {
	NombreRol string `db:"nombrerol"`
}

func InsertarNuevoRol(db *sql.DB, nombreRol string) error {
	crud := database.NuevoCRUD(db)
	datos := UserModels.Rol{
		NombreRol: nombreRol,
	}
	return crud.Insertar(`"Roles"`, datos)
}

func ActualizarRol(db *sql.DB, idRol int, nombreRol string) error {
	crud := database.NuevoCRUD(db)
	datos := RolVariable{
		NombreRol: nombreRol,
	}
	where := "id_rol = $2"
	return crud.Actualizar(`"Roles"`, datos, where, idRol)
}

func SeleccionarRoles(db *sql.DB, condicion string, args ...interface{}) ([]UserModels.Rol, error) {
	crud := database.NuevoCRUD(db)
	var roles []UserModels.Rol

	columnas := []string{"id_rol", "nombrerol"}
	var rows *sql.Rows
	var err error
	if condicion == "" {
		rows, err = crud.Seleccionar(`"Roles"`, columnas, "", args...)
	} else {
		rows, err = crud.Seleccionar(`"Roles"`, columnas, condicion, args...)
	}
	if err != nil {
		return nil, fmt.Errorf("error en Select: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var rol UserModels.Rol
		err := rows.Scan(
			&rol.ID,
			&rol.NombreRol,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		roles = append(roles, rol)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error después de iterar filas: %v", err)
	}
	return roles, nil
}
