package BDoperators

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

// CRUD proporciona operaciones genéricas para cualquier tabla
type CRUD struct {
	DB *sql.DB
}

// NewCRUD crea una nueva instancia de CRUD
func NewCRUD(db *sql.DB) *CRUD {
	return &CRUD{DB: db}
}

// Insert inserta un registro en la tabla especificada
func (c *CRUD) Insert(table string, data interface{}) error {
	// Obtener el tipo y valor de la estructura
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()

	// Preparar los nombres de las columnas y los valores
	var columns []string
	var placeholders []string
	var values []interface{}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		// Obtener el nombre de la columna (usando el tag `json` o el nombre del campo)
		column := strings.ToLower(field.Name)
		if jsonTag := field.Tag.Get("json"); jsonTag != "" {
			column = strings.Split(jsonTag, ",")[0]
		}

		columns = append(columns, column)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
		values = append(values, value)
	}

	// Construir la consulta SQL
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	// Ejecutar la consulta
	_, err := c.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error al insertar en la tabla %s: %v", table, err)
	}

	return nil
}

// Update actualiza un registro en la tabla especificada
func (c *CRUD) Update(table string, id string, data interface{}) error {
	// Obtener el tipo y valor de la estructura
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()

	// Preparar las columnas y los valores
	var updates []string
	var values []interface{}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		// Obtener el nombre de la columna (usando el tag `json` o el nombre del campo)
		column := strings.ToLower(field.Name)
		if jsonTag := field.Tag.Get("json"); jsonTag != "" {
			column = strings.Split(jsonTag, ",")[0]
		}

		updates = append(updates, fmt.Sprintf("%s = $%d", column, i+1))
		values = append(values, value)
	}

	// Agregar el ID al final de los valores
	values = append(values, id)

	// Construir la consulta SQL
	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = $%d",
		table,
		strings.Join(updates, ", "),
		len(values),
	)

	// Ejecutar la consulta
	_, err := c.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error al actualizar en la tabla %s: %v", table, err)
	}

	return nil
}

// Delete elimina un registro de la tabla especificada
func (c *CRUD) Delete(table string, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table)
	_, err := c.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar de la tabla %s: %v", table, err)
	}
	return nil
}

func (c *CRUD) Select(table string, columns []string, whereClause string, args ...interface{}) (*sql.Rows, error) {
	// Construir la lista de columnas
	cols := "*"
	if len(columns) > 0 {
		cols = strings.Join(columns, ", ")
	}
	// Construir la consulta SQL
	query := fmt.Sprintf("SELECT %s FROM %s", cols, table)

	// Agregar WHERE si se especificó
	if whereClause != "" {
		query += " WHERE " + whereClause
	}
	// Ejecutar la consulta
	rows, err := c.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error al realizar select en tabla %s: %v", table, err)
	}
	return rows, nil
}
