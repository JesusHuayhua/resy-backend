package database

import (
	"database/sql"
	"fmt"
	"soa/pkg/services/ServicioUsuario/repository/crypto"

	_ "github.com/lib/pq"
)

type Config struct {
	Driver       string `json:"driver"`
	TipoDriver   string `json:"tipo_de_driver"`
	DBName       string `json:"base_de_datos"`
	Host         string `json:"nombre_de_host"`
	Port         string `json:"puerto"`
	User         string `json:"usuario"`
	DatabaseName string `json:"esquemabd"`
	Password     string `json:"contrasenha"`
}

type DBManager struct {
	DB *sql.DB
}

func NuevoDBManager(config Config, cryptoCtx *crypto.EnvelopeCrypto) (*DBManager, error) {
	password, err := cryptoCtx.Decrypt(config.Password)
	if err != nil {
		return nil, fmt.Errorf("error al descifrar password: %v", err)
	}
	//Workaround temporal de sslmode desactivado (en AWS si puedes usar require)
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		password, //
		config.DBName,
	)
	db, err := sql.Open(config.Driver, connStr)
	if err != nil {
		return nil, fmt.Errorf("error al conectar a PostgreSQL: %v", err)
	}
	_, err = db.Exec(fmt.Sprintf("SET search_path TO '%s', public", config.DatabaseName))
	if err != nil {
		return nil, fmt.Errorf("error al configurar search_path: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error al hacer ping a la BD: %v", err)
	}
	return &DBManager{DB: db}, nil
}

// Cerrar cierra la conexión
func (m *DBManager) Cerrar() {
	m.DB.Close()
}
