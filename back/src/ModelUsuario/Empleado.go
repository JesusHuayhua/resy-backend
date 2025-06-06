package Usuario

import (
	"time"
)

// Empleado representa la tabla empleado que hereda de usuario
type Empleado struct {
	Usuario                     // Embedded struct para herencia
	SedeID            string    `json:"sede_id" gorm:"size:36;not null"`
	Cargo             string    `json:"cargo" gorm:"size:50;not null;check:cargo IN ('Mesero', 'Chef', 'Gerente', 'Delivery')"`
	FechaContratacion time.Time `json:"fecha_contratacion" gorm:"not null"`
	Salario           float64   `json:"salario" gorm:"type:decimal(10,2);not null"`
	Sede              Sede      `json:"sede,omitempty" gorm:"foreignKey:SedeID"` // Relación opcional
}
