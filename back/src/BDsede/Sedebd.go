package bdsede

import (
	"back/src/crudi"
	"back/src/modelsede"
	"database/sql"
	"fmt"
	"log"
)

func InsertNewSede(db *sql.DB, sede_id string, nombre string, direccion string, ciudad string, telefono string, AFORO_TOTAL int, horario_apertura string, horario_cierre string) error {
	crud := crudi.NewCRUD(db)

	data := modelsede.Sede{
		SedeID:          sede_id,
		Nombre:          nombre,
		Direccion:       direccion,
		Ciudad:          ciudad,
		Telefono:        telefono,
		AforoTotal:      AFORO_TOTAL,
		HorarioApertura: horario_apertura,
		HorarioCierre:   horario_cierre,
	}

	if err := crud.Insert("sede", data); err != nil {
		log.Printf("Error inserting data: %v", err)
		return err
	}
	return nil
}

func UpdateSede(db *sql.DB, sede_id string, nombre string, direccion string, ciudad string, telefono string, AFORO_TOTAL int, horario_apertura string, horario_cierre string, activa bool) error {
	crud := crudi.NewCRUD(db)

	// Usar la estructura del modelo en lugar de un mapa
	data := modelsede.Sede{
		SedeID:          sede_id,
		Nombre:          nombre,
		Direccion:       direccion,
		Ciudad:          ciudad,
		Telefono:        telefono,
		AforoTotal:      AFORO_TOTAL,
		HorarioApertura: horario_apertura,
		HorarioCierre:   horario_cierre,
		Activa:          activa,
	}

	if err := crud.Update("sede", sede_id, data); err != nil {
		log.Printf("Error updating data: %v", err)
		return err
	}
	return nil
}

func SelectSedes(db *sql.DB, condicion string, args ...interface{}) ([]modelsede.Sede, error) {
	crud := crudi.NewCRUD(db)
	var sedes []modelsede.Sede

	// Usar nombres de columnas exactamente como están en la BD
	columnas := []string{
		"id",
		"nombre",
		"direccion",
		"ciudad",
		"telefono",
		"AFORO_TOTAL", // Asegúrate que coincide con el nombre en la BD
		"horario_apertura",
		"horario_cierre",
		"activa",
	}

	// Si no hay condición, pasar una cadena vacía y ningún argumento
	var rows *sql.Rows
	var err error

	if condicion == "" {
		rows, err = crud.Select("sede", columnas, "", args...)
	} else {
		rows, err = crud.Select("sede", columnas, condicion, args...)
	}

	if err != nil {
		return nil, fmt.Errorf("error en Select: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sede modelsede.Sede
		err := rows.Scan(
			&sede.SedeID,
			&sede.Nombre,
			&sede.Direccion,
			&sede.Ciudad,
			&sede.Telefono,
			&sede.AforoTotal,
			&sede.HorarioApertura,
			&sede.HorarioCierre,
			&sede.Activa,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		sedes = append(sedes, sede)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error después de iterar filas: %v", err)
	}

	return sedes, nil
}
