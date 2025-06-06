package modelsede

// Sede representa la tabla sede
type Sede struct {
	SedeID          string    `json:"id" gorm:"primaryKey;size:36"`
	Nombre          string    `json:"nombre" gorm:"size:100;not null"`
	Direccion       string    `json:"direccion" gorm:"size:200;unique;not null"`
	Ciudad          string    `json:"ciudad" gorm:"size:50;not null"`
	Telefono        string    `json:"telefono" gorm:"size:20;unique;not null"`
	AforoTotal      int       `json:"aforo_total" gorm:"not null"`
	HorarioApertura string    `json:"horario_apertura" gorm:"type:time;not null"` // Usar string o time.Time dependiendo de tu necesidad
	HorarioCierre   string    `json:"horario_cierre" gorm:"type:time;not null"`   // Usar string o time.Time dependiendo de tu necesidad
	Activa          bool      `json:"activa" gorm:"default:true"`
}