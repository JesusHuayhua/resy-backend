package database

import (
	"database/sql"
	"fmt"
	"log"
	"soa/services/users/pkg/repository/shared/config"
	"sync"

	_ "github.com/lib/pq"
)

var (
	commonDBInstance *sql.DB
	commonOnce       sync.Once
)

func InitCommonDB() *sql.DB {
	commonOnce.Do(func() {
		dbURL := config.LoadConfig()
		var err error
		commonDBInstance, err = sql.Open("postgres", dbURL)
		if err != nil {
			log.Fatalf("Error al abrir la conexión a la base de datos compartida: %v", err)
		}

		if err = commonDBInstance.Ping(); err != nil {
			log.Fatalf("Error al conectar a la base de datos compartida: %v", err)
		}

		commonDBInstance.SetMaxOpenConns(50) // Puede requerir más conexiones para una DB compartida
		commonDBInstance.SetMaxIdleConns(25)
		commonDBInstance.SetConnMaxLifetime(5 * 60 * 1000)

		fmt.Println("Conexión a la base de datos compartida inicializada exitosamente.")
	})
	return commonDBInstance
}

func CloseCommonDB() {
	if commonDBInstance != nil {
		err := commonDBInstance.Close()
		if err != nil {
			log.Printf("Error al cerrar la conexión a la base de datos compartida: %v", err)
		} else {
			fmt.Println("Conexión a la base de datos compartida cerrada.")
		}
	}
}

// Requiere try catch (esto puede tirar un error.)
func GetCommonDB() *sql.DB {
	if commonDBInstance == nil {
		log.Printf("La conexión a la base de datos compartida no ha sido inicializada. Inicializando...")
		InitCommonDB()
	}
	return commonDBInstance
}
