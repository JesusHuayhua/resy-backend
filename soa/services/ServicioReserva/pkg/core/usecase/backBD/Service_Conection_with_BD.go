package backBD

import (
	ReservaModels "ServicioReserva/pkg/core/domain"
	"ServicioReserva/pkg/repository"
	repoInterface "ServicioReserva/pkg/repository/interfaces"
	"database/sql"
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

// Listar reservas (todas o por filtro)
func (s *ServicioReserva) ListarReservas(condicion string, args ...interface{}) ([]ReservaModels.Reserva, error) {
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

	rows, err := s.crud.Seleccionar(`"ResyDB"."Reserva"`, columnas, condicion, args...)
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
