// +// Paquete repository contiene implementaciones de repositorios para interactuar con la base de datos
package repository

// +// Este paquete define la interfaz UserRepository y su implementación UserRepositoryImpl.
import (
	"database/sql"
	"soa/pkg/services/ServicioUsuario/repository/database"
	"soa/pkg/services/ServicioUsuario/repository/interfaces"
)

// UserRepositoryImpl es la implementación de la interfaz UserRepository
type UserRepositoryImpl struct {
	crud *database.CRUD
}

// NewUserRepository crea una nueva instancia de UserRepositoryImpl
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
