package repository

import (
	"database/sql"
	"fmt"
	"reflect"
	commonDB "soa/services/users/pkg/core/shared/database"
	"soa/services/users/pkg/repository/interfaces"
	"strings"
)

type CRUD struct {
	DB *sql.DB
}

func NuevoCRUD() interfaces.UserRepository {
	db := commonDB.GetCommonDB()
	return &CRUD{
		DB: db,
	}
}

func (c *CRUD) Insertar(tabla string, datos interface{}) error {
	v := reflect.ValueOf(datos)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	var columns []string
	var placeholders []string
	var values []interface{}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		column := strings.ToLower(field.Name)
		if dbTag := field.Tag.Get("db"); dbTag != "" {
			column = strings.Split(dbTag, ",")[0]
		}

		columns = append(columns, column)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
		values = append(values, value)
	}
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		tabla,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)
	_, err := c.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error al insertar en la tabla %s: %v", tabla, err)
	}

	return nil
}

func (c *CRUD) Actualizar(tabla string, datos interface{}, whereClause string, whereArgs ...interface{}) error {
	v := reflect.ValueOf(datos)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	var updates []string
	var values []interface{}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		column := field.Name
		if dbTag := field.Tag.Get("db"); dbTag != "" {
			column = strings.Split(dbTag, ",")[0]
		}
		updates = append(updates, fmt.Sprintf("%s = $%d", column, i+1))
		values = append(values, value)
	}
	values = append(values, whereArgs...)
	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s",
		tabla,
		strings.Join(updates, ", "),
		whereClause,
	)
	_, err := c.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error al actualizar en la tabla %s: %v", tabla, err)
	}

	return nil
}

func (c *CRUD) Eliminar(tabla string, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tabla)
	_, err := c.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar de la tabla %s: %v", tabla, err)
	}
	return nil
}

func (c *CRUD) Seleccionar(tabla string, columnas []string, whereClause string, args ...interface{}) (*sql.Rows, error) {
	cols := "*"
	if len(columnas) > 0 {
		cols = strings.Join(columnas, ", ")
	}
	query := fmt.Sprintf("SELECT %s FROM %s", cols, tabla)
	if whereClause != "" {
		query += " WHERE " + whereClause
	}
	rows, err := c.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error al realizar select en tabla %s: %v", tabla, err)
	}
	return rows, nil
}
