package main

import (
	"log"

	ServicioUsuario "ServicioUsuario/pkg/core/usecase/BDoperators"

	"github.com/Shauanth/Singleton_Encription_ServiceGolang/crypton"
	"github.com/Shauanth/Singleton_Encription_ServiceGolang/database"
	_ "github.com/lib/pq" //Driver Para base de datos postgreSQL
)

func main() {
	databaseInformation := database.Config{
		Driver:       "postgres",
		TipoDriver:   "PostgreSQL",
		DBName:       "postgres",
		Host:         "localhost",
		Port:         "5432",
		User:         "postgres",
		DatabaseName: "ResyDB",
		Password:     "a5i3aJtCcU0P56OTDmXSGb/kfkZY1/lEGdh5eVsbomGgL6ss7Q==",
	}
	encriptacionKey := crypton.Config{
		EncryptionKey: "53WDFETRFQFC1?*OS!0LNSADJUER2YU8",
		Salt:          "RCumoV7j",
	}
	dbManager, err := database.NuevoDBManager(databaseInformation, encriptacionKey)
	if err != nil {
		log.Fatalf("Error al conectar a la BD: %v", err)
	}
	defer dbManager.Cerrar()
	// Inicializa el servicio de usuario
	_ = ServicioUsuario.NuevoServicioUsuario(dbManager.DB, encriptacionKey)
	// Ejemplo de uso: insertar un usuario
	/*
		status, err := servicioUsuario.InsertarNuevoUsuario(
			"Juan", "Pérez", "juan.perez@email.com", "123456789",
			time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), "mi_contraseña_segura", 1,
		)
		if err != nil {
			fmt.Printf("Error al insertar usuario: %v\n", err)
		} else {
			fmt.Printf("Usuario insertado, status: %v\n", status)
		}
	*/
}
