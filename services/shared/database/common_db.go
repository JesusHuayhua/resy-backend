package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq" // O el driver de tu BD compartida
	"mi-proyecto-soa/shared/config"
)

var (
	commonDBInstance *sql.DB // La instancia de la conexión a la base de datos común
	commonOnce       sync.Once // Para asegurar que la conexión se inicializa una sola vez
)

// InitCommonDB inicializa la conexión a la base de datos compartida utilizando el patrón Singleton.
func InitCommonDB() *sql.DB {
	commonOnce.Do(func() {
		cfg := config.LoadConfig() // Carga la configuración del servicio, que incluye la URL de la DB.

		var err error
		// Asumiendo que la URL para la DB compartida está en cfg.CommonDatabaseURL
		// O que es la única DatabaseURL en la configuración global
		dbURL := cfg.DatabaseURL // Usamos la misma variable de ejemplo, pero ten claro que es la global
		
		commonDBInstance, err = sql.Open("postgres", dbURL) // Reemplaza "postgres" y dbURL
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

// CloseCommonDB cierra la conexión a la base de datos compartida.
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

// GetCommonDB devuelve la instancia existente de la base de datos compartida.
func GetCommonDB() *sql.DB {
	if commonDBInstance == nil {
		log.Fatal("La conexión a la base de datos compartida no ha sido inicializada. Llama a InitCommonDB() primero.")
	}
	return commonDBInstance
}