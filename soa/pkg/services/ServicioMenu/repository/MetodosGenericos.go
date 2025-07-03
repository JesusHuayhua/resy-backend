package repository

import (
	"database/sql"
	"soa/pkg/services/ServicioMenu/repository/database"
	"soa/pkg/services/ServicioMenu/repository/interfaces"
)

type PlatoRepositoryImpl struct {
	crud *database.CRUD
}

func NewPlatoRepository(db *sql.DB) interfaces.PlatoRepository {
	return &PlatoRepositoryImpl{
		crud: database.NuevoCRUD(db),
	}
}

func (r *PlatoRepositoryImpl) Insertar(tabla string, datos interface{}) error {
	return r.crud.Insertar(tabla, datos)
}

func (r *PlatoRepositoryImpl) Actualizar(tabla string, datos interface{}, whereClause string, whereArgs ...interface{}) error {
	return r.crud.Actualizar(tabla, datos, whereClause, whereArgs...)
}

func (r *PlatoRepositoryImpl) Eliminar(tabla string, id string) error {
	return r.crud.Eliminar(tabla, id)
}

func (r *PlatoRepositoryImpl) Seleccionar(tabla string, columnas []string, whereClause string, args ...interface{}) (*sql.Rows, error) {
	return r.crud.Seleccionar(tabla, columnas, whereClause, args...)
}

func (r *PlatoRepositoryImpl) Crud() *database.CRUD {
	return r.crud
}
