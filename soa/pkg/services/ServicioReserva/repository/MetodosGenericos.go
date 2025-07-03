package repository

import (
	"database/sql"
	"soa/pkg/services/ServicioReserva/repository/database"
	"soa/pkg/services/ServicioReserva/repository/interfaces"
)

type ReservaRepositoryImpl struct {
	crud *database.CRUD
}

func NewReservaRepository(db *sql.DB) interfaces.ReservaRepository {
	return &ReservaRepositoryImpl{
		crud: database.NuevoCRUD(db),
	}
}

func (r *ReservaRepositoryImpl) Insertar(tabla string, datos interface{}) error {
	return r.crud.Insertar(tabla, datos)
}

func (r *ReservaRepositoryImpl) Actualizar(tabla string, datos interface{}, whereClause string, whereArgs ...interface{}) error {
	return r.crud.Actualizar(tabla, datos, whereClause, whereArgs...)
}

func (r *ReservaRepositoryImpl) Eliminar(tabla string, id string) error {
	return r.crud.Eliminar(tabla, id)
}

func (r *ReservaRepositoryImpl) Seleccionar(tabla string, columnas []string, whereClause string, args ...interface{}) (*sql.Rows, error) {
	return r.crud.Seleccionar(tabla, columnas, whereClause, args...)
}

func (r *ReservaRepositoryImpl) Crud() *database.CRUD {
	return r.crud
}
