package backBD

import (
	ReservaModels "ServicioReserva/pkg/core/domain"
	"ServicioReserva/pkg/repository"
	repoInterface "ServicioReserva/pkg/repository/interfaces"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type ServicioReserva struct {
	crud repoInterface.ReservaRepository
}

func NuevoServicioReserva(db *sql.DB) *ServicioReserva {
	crud := repository.NewReservaRepository(db)
	return &ServicioReserva{crud: crud}
}

// Insertar una nueva reserva (registrado o no registrado)
func (s *ServicioReserva) InsertarReserva(data ReservaModels.ReservaData) error {
	return s.crud.Insertar(`"ResyDB"."Reserva"`, data)
}

// Actualizar una reserva
func (s *ServicioReserva) ActualizarReserva(id string, data ReservaModels.ReservaData) error {
	where := "id_reserva = $1"
	return s.crud.Actualizar(`"ResyDB"."Reserva"`, data, where, id)
}

// Eliminar una reserva
func (s *ServicioReserva) EliminarReserva(id string) error {
	return s.crud.Eliminar(`"ResyDB"."Reserva"`, id)
}

// ListarReservas lista reservas aplicando filtros opcionales.
// Los filtros se pasan como un mapa donde la clave es el nombre de la columna y el valor es el valor a filtrar.
func (s *ServicioReserva) ListarReservas(filtros map[string]interface{}) ([]ReservaModels.Reserva, error) {
	columnas := []string{
		"id_reserva",
		"id_cliente",
		"nombre_cliente",
		"telefono_cliente",
		"correo_cliente",
		"fecha_reservada",
		"num_personas",
		"estado_reserva",
		"especificaciones",
	}

	tabla := `"ResyDB"."Reserva"`
	condiciones := []string{}
	args := []interface{}{}
	paramIndex := 1 // Para los placeholders de la consulta SQL ($1, $2, etc. para PostgreSQL)

	if idReserva, ok := filtros["id_reserva"]; ok && idReserva != "" {
		condiciones = append(condiciones, fmt.Sprintf("id_reserva = $%d", paramIndex))
		args = append(args, idReserva)
		paramIndex++
	}
	if idCliente, ok := filtros["id_cliente"]; ok && idCliente != 0 {
		condiciones = append(condiciones, fmt.Sprintf("id_cliente = $%d", paramIndex))
		args = append(args, idCliente)
		paramIndex++
	}
	if fechaReservada, ok := filtros["fecha_reservada"]; ok { // No verificar IsZero() aquí, porque time.Time{} es válido si no se parseó.
		if fecha, ok := fechaReservada.(time.Time); ok && !fecha.IsZero() {
			// Compara solo la parte de la fecha, ignorando la hora
			condiciones = append(condiciones, fmt.Sprintf("fecha_reservada::date = $%d::date", paramIndex))
			args = append(args, fecha)
			paramIndex++
		}
	}
	if estadoReserva, ok := filtros["estado_reserva"]; ok && estadoReserva != "" {
		condiciones = append(condiciones, fmt.Sprintf("estado_reserva = $%d", paramIndex))
		args = append(args, estadoReserva)
		paramIndex++
	}

	whereClause := ""
	if len(condiciones) > 0 {
		whereClause = strings.Join(condiciones, " AND ")
	}
	// Usar el método Seleccionar existente
	rows, err := s.crud.Seleccionar(tabla, columnas, whereClause, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservas []ReservaModels.Reserva
	for rows.Next() {
		var r ReservaModels.Reserva
		var data ReservaModels.ReservaData

		err := rows.Scan(
			&r.IDReserva,
			&data.IDCliente,
			&data.NombreCliente,
			&data.TelefonoCliente,
			&data.CorreoCliente,
			&data.FechaReservada,
			&data.NumPersonas,
			&data.EstadoReserva,
			&data.Especificaciones,
		)

		if err != nil {
			return nil, err
		}

		r.DataReserva = data
		reservas = append(reservas, r)
	}
	return reservas, nil
}

// ListarReservasParaRecordatorio devuelve las reservas que ocurrirán entre 1h58m y 2h02m desde ahora
func (s *ServicioReserva) ListarReservasParaRecordatorio() ([]ReservaModels.Reserva, error) {
	columnas := []string{
		"id_reserva",
		"id_cliente",
		"nombre_cliente",
		"telefono_cliente",
		"correo_cliente",
		"fecha_reservada",
		"num_personas",
		"estado_reserva",
		"especificaciones",
	}

	tabla := `"ResyDB"."Reserva"`
	desde := time.Now().Add(1*time.Hour + 58*time.Minute)
	hasta := time.Now().Add(2*time.Hour + 2*time.Minute)

	whereClause := "fecha_reservada BETWEEN $1 AND $2"
	args := []interface{}{desde, hasta}

	rows, err := s.crud.Seleccionar(tabla, columnas, whereClause, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservas []ReservaModels.Reserva
	for rows.Next() {
		var r ReservaModels.Reserva
		var data ReservaModels.ReservaData

		err := rows.Scan(
			&r.IDReserva,
			&data.IDCliente,
			&data.NombreCliente,
			&data.TelefonoCliente,
			&data.CorreoCliente,
			&data.FechaReservada,
			&data.NumPersonas,
			&data.EstadoReserva,
			&data.Especificaciones,
		)
		if err != nil {
			return nil, err
		}
		r.DataReserva = data
		reservas = append(reservas, r)
	}

	return reservas, nil
}
