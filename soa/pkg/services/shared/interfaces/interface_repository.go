// interfaces para las operacione con BD
package interfaces

import "database/sql"

type Repository interface {
	Insertar(tabla string, datos interface{}) error
	Actualizar(tabla string, datos interface{}, whereClause string, whereArgs ...interface{}) error
	Eliminar(tabla string, columna string, id string) error
	Seleccionar(tabla string, columnas []string, whereClause string, args ...interface{}) (*sql.Rows, error)
}
