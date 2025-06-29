package ReservaModels

import (
	"database/sql"
	"time"
)

// Reserva principal (para usuario registrado o no registrado)
type Reserva struct {
	IDReserva   string `db:"id_reserva"`
	DataReserva ReservaData
}

// Para inserción de reservas (sin id_reserva, que es autogenerado)
type ReservaData struct {
	IDCliente        sql.NullInt64  `db:"id_cliente"`
	NombreCliente    sql.NullString `db:"nombre_cliente"`   // Añadir NullString
	TelefonoCliente  sql.NullString `db:"telefono_cliente"` // Añadir NullString
	CorreoCliente    sql.NullString `db:"correo_cliente"`   // Añadir NullString
	FechaReservada   time.Time      `db:"fecha_reservada"`
	NumPersonas      int            `db:"num_personas"`
	EstadoReserva    string         `db:"estado_reserva"`
	Especificaciones string         `db:"especificaciones"`
}

// Relación entre reserva y pago
type ReservaPago struct {
	IDPago    int    `db:"id_pago"`
	IDReserva string `db:"id_reserva"`
}

// Platos reservados en una reserva
type PlatoReservado struct {
	IDLinea int `db:"id_linea"`
	Data    PlatoData
}

type PlatoData struct {
	IDReserva string  `db:"id_reserva"`
	IDPlato   int     `db:"id_plato"`
	Cantidad  int     `db:"cantidad"`
	Subtotal  float64 `db:"subtotal"`
}
