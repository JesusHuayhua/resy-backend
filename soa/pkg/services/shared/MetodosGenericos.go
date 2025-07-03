// +// Paquete repository contiene implementaciones de repositorios para interactuar con la base de datos
package repository

// +// Este paquete define la interfaz UserRepository y su implementación UserRepositoryImpl.
import (
	"database/sql"
	"soa/pkg/services/shared/database"
	"soa/pkg/services/shared/interfaces"
)

type RepositoryImpl struct {
	crud *database.CRUD
}

// NewUserRepository crea una nueva instancia de UserRepositoryImpl
func NewRepository(db *sql.DB) interfaces.Repository {
	return &RepositoryImpl{
		crud: database.NuevoCRUD(db),
	}
}

func (r *RepositoryImpl) Insertar(tabla string, datos interface{}) error {
	return r.crud.Insertar(tabla, datos)
}

func (r *RepositoryImpl) Actualizar(tabla string, datos interface{}, whereClause string, whereArgs ...interface{}) error {
	return r.crud.Actualizar(tabla, datos, whereClause, whereArgs...)
}

func (r *RepositoryImpl) Eliminar(tabla string, id string) error {
	return r.crud.Eliminar(tabla, id)
}

func (r *RepositoryImpl) Seleccionar(tabla string, columnas []string, whereClause string, args ...interface{}) (*sql.Rows, error) {
	return r.crud.Seleccionar(tabla, columnas, whereClause, args...)
}

func (r *RepositoryImpl) Crud() *database.CRUD {
	return r.crud
}
