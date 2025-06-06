package db

import (
	"back/src/crypto"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	_ "github.com/lib/pq"
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

func NewDBManager(configPath string) (*DBManager, error) {
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo JSON: %v", err)
	}
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar JSON: %v", err)
	}
	password, err := crypto.Decrypt(config.Password)
	if err != nil {
		return nil, fmt.Errorf("error al descifrar password: %v", err)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		password, // Contraseña descifrada
		config.DBName,
	)

	db, err := sql.Open(config.Driver, connStr)
	_, err = db.Exec("SET search_path TO restaurante, public")
	if err != nil {
		return nil, fmt.Errorf("error al configurar search_path: %v", err)
	}
	if err != nil {
		return nil, fmt.Errorf("error al conectar a PostgreSQL: %v", err)
	}

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
