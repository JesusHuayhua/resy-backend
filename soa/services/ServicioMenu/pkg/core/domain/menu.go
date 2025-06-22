package MenuModels

import (
	"time"
)

type DiaSemana string

const (
	Lunes     DiaSemana = "Lunes"
	Martes    DiaSemana = "Martes"
	Miercoles DiaSemana = "Miercoles"
	Jueves    DiaSemana = "Jueves"
	Viernes   DiaSemana = "Viernes"
	Sabado    DiaSemana = "Sabado"
	Domingo   DiaSemana = "Domingo"
)

// MenuSemanal representa un menú semanal.
type MenuSemanal struct {
	IDMenu string `db:"id_menu"`
	Semana MenuSemanalVariable
}

type MenuSemanalVariable struct {
	FechaDeInicio time.Time `db:"fechadeinicio"`
	FechaFin      time.Time `db:"fecha_fin"`
	MenusDiarios  []Menudia
}

// Menudia representa el menú de un día específico dentro del menú semanal.
type Menudia struct {
	IDDia int `db:"id_dia"`
	Info  MenudiaVariable
}

type MenudiaVariable struct {
	IDMenu    string `db:"id_menu"`
	DiaSemana string `db:"dia_semana"`
	Platos    []PlatosEnMenudia
}

// PlatosEnMenudia representa la relación entre platos y el menú de un día.
type PlatosEnMenudia struct {
	IDDia   int `db:"id_dia"` // FK a Menudia
	InfoDia PlatosEnMenudiaVariable
}

type PlatosEnMenudiaVariable struct {
	Plato                Plato
	CantidadDelPlato     int  `db:"cantidad_del_plato"`
	DisponibleParaVender bool `db:"disponible_para_vender"`
}
