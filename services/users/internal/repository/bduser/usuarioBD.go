package bduser

import (
	"RESI-BACKEND/services/shared/utils"
	"RESI-BACKEND/services/users/pkg/models/UserModels"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type UsuarioInsertar struct {
	Nombres         string    `db:"nombres"`
	Apellidos       string    `db:"apellidos"`
	Correo          string    `db:"correo"`
	FechaNacimiento time.Time `db:"fechanacimiento"`
	Contrasenia     string    `db:"contrasenia"`
	Rol             int       `db:"rol"`
	EstadoAcceso    bool      `db:"estadoacceso"`
}

type UsuarioActualizar struct {
	Nombres         string    `db:"nombres"`
	Apellidos       string    `db:"apellidos"`
	Correo          string    `db:"correo"`
	FechaNacimiento time.Time `db:"fechanacimiento"`
	Contrasenia     string    `db:"contrasenia"`
	Rol             int       `db:"rol"`
	EstadoAcceso    bool      `db:"estadoacceso"`
}

func InsertarNuevoUsuario(db *sql.DB, nombres string, apellidos string, correo string, fechaNacimiento time.Time, contrasenia string, rol int) error {
	crud := utils.NuevoCRUD(db)
	datos := UsuarioInsertar{
		Nombres:         nombres,
		Apellidos:       apellidos,
		Correo:          correo,
		FechaNacimiento: fechaNacimiento,
		Contrasenia:     contrasenia,
		Rol:             rol,
		EstadoAcceso:    true,
	}
	return crud.Insertar(`"Usuario"`, datos)
}

func ActualizarUsuario(db *sql.DB, idUsuario int, nombres string, apellidos string, correo string, fechaNacimiento time.Time, contrasenia string, rol int, estado bool) error {
	crud := utils.NuevoCRUD(db)
	datos := UsuarioActualizar{
		Nombres:         nombres,
		Apellidos:       apellidos,
		Correo:          correo,
		FechaNacimiento: fechaNacimiento,
		Contrasenia:     contrasenia,
		Rol:             rol,
		EstadoAcceso:    estado,
	}
	where := "id_usuario = $8"
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
		"fechanacimiento",
		"contrasenia",
		"rol",
		"estadoacceso",
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

type CRUD struct {
	db *sql.DB
}

func (crud *CRUD) Actualizar(tabla string, datos interface{}, condicion string, whereArgs ...interface{}) error {
	v := reflect.ValueOf(datos)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()

	setClauses := make([]string, 0)
	values := make([]interface{}, 0)

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if field.Tag.Get("db") != "" {
			setClauses = append(setClauses, fmt.Sprintf("%s = ?", field.Tag.Get("db")))
			values = append(values, value)
		}
	}
	values = append(values, whereArgs...)

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE %s`, tabla, strings.Join(setClauses, ", "), condicion)
	_, err := crud.db.Exec(query, values...)
	return err
}
