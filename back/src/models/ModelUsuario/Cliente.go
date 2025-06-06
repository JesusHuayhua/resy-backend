package Usuario

import "time"

// Cliente representa la tabla cliente que hereda de usuario
type Cliente struct {
	Usuario                   // Embedded struct para herencia
	FechaNacimiento      *time.Time `json:"fecha_nacimiento,omitempty"`
	EsFrecuente          bool       `json:"es_frecuente" gorm:"default:false"`
	PuntosFidelizacion   int        `json:"puntos_fidelizacion" gorm:"default:0"`
	RestriccionesAlimenticias string `json:"restricciones_alimenticias,omitempty" gorm:"size:255"`
}