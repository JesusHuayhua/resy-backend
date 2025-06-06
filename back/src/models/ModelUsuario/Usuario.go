package Usuario

import "time"

// Usuario representa la tabla usuario
type Usuario struct {
	UsuarioID       string    `json:"id" gorm:"primaryKey;size:10"`
	Nombres         string    `json:"nombres" gorm:"size:150"`
	Apellidos       string    `json:"apellidos" gorm:"size:150"`
	TipoDocumento   string    `json:"tipo_documento" gorm:"size:20;check:tipo_documento IN ('DNI', 'CE', 'PASAPORTE')"`
	NumeroDocumento string    `json:"numero_documento" gorm:"size:20;unique"`
	Telefono        string    `json:"telefono" gorm:"size:20;not null"`
	Email           string    `json:"email" gorm:"size:100;unique"`
	FechaRegistro   time.Time `json:"fecha_registro" gorm:"not null"`
	Activo          bool      `json:"activo" gorm:"default:true"`
	Contrasena      string    `json:"-" gorm:"size:255;not null;default:'*'"`
}