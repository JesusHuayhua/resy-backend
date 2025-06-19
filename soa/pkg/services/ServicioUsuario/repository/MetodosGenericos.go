package repository

import (
	"ServicioUsuario/pkg/repository/database"
	"ServicioUsuario/pkg/repository/interfaces"
	"database/sql"
)

type UserRepositoryImpl struct {
	crud *database.CRUD
}

func NewUserRepository(db *sql.DB) interfaces.UserRepository {
	return &UserRepositoryImpl{
		crud: database.NuevoCRUD(db),
	}
}

func (r *UserRepositoryImpl) Insertar(tabla string, datos interface{}) error {
	return r.crud.Insertar(tabla, datos)
}

func (r *UserRepositoryImpl) Actualizar(tabla string, datos interface{}, whereClause string, whereArgs ...interface{}) error {
	return r.crud.Actualizar(tabla, datos, whereClause, whereArgs...)
}

func (r *UserRepositoryImpl) Eliminar(tabla string, id string) error {
	return r.crud.Eliminar(tabla, id)
}

func (r *UserRepositoryImpl) Seleccionar(tabla string, columnas []string, whereClause string, args ...interface{}) (*sql.Rows, error) {
	return r.crud.Seleccionar(tabla, columnas, whereClause, args...)
}

func (r *UserRepositoryImpl) Crud() *database.CRUD {
	return r.crud
}
