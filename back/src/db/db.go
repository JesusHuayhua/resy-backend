package db

import (
	"back/src/crypto"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Driver de PostgreSQL
)

// Config representa la estructura del JSON
type Config struct {
	Driver     string `json:"driver"`
	TipoDriver string `json:"tipo_de_driver"`
	DBName     string `json:"base_de_datos"`
	Host       string `json:"nombre_de_host"`
	Port       string `json:"puerto"`
	User       string `json:"usuario"`
	Password   string `json:"contrasenha"` // Asumimos que está cifrada
}

// DBManager maneja la conexión a la BD
type DBManager struct {
	DB *sql.DB
}

// NewDBManager carga la configuración desde JSON y conecta a PostgreSQL
func NewDBManager(configPath string) (*DBManager, error) {
	// 1. Leer el archivo JSON
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo JSON: %v", err)
	}

	// 2. Decodificar JSON en la estructura Config
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar JSON: %v", err)
	}
	password, err := crypto.Decrypt(config.Password)
	if err != nil {
		return nil, fmt.Errorf("error al descifrar password: %v", err)
	}
	// 4. Cadena de conexión
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		password, // Contraseña descifrada
		config.DBName,
	)

	// 5. Conectar a PostgreSQL
	db, err := sql.Open(config.Driver, connStr)
	// En NewDBManager, después de sql.Open
	_, err = db.Exec("SET search_path TO restaurante, public")
	if err != nil {
		return nil, fmt.Errorf("error al configurar search_path: %v", err)
	}
	if err != nil {
		return nil, fmt.Errorf("error al conectar a PostgreSQL: %v", err)
	}

	// 6. Verificar conexión
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error al hacer ping a la BD: %v", err)
	}

	return &DBManager{DB: db}, nil
}

// Close cierra la conexión
func (m *DBManager) Close() {
	m.DB.Close()
}
