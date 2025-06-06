package main

//Pa los que no saben hacer un ruleset
//NO CAMBIEN EL MODULE
// USEN RESI-BACKEND COMO CARPETA RAIZ
import (
	"RESI-BACKEND/services/shared/database"
	"RESI-BACKEND/services/users/internal/repository/bduser"
	"log"
	"time"
)

func main() {
	dbManager, err := database.NuevoDBManager("conf2/key.json")
	if err != nil {
		log.Fatalf("Error al conectar a la BD: %v", err)
	}
	defer dbManager.Cerrar()
	err = bduser.InsertarNuevoUsuario(dbManager.DB, "Nombre", "Apellido", "correo@ejemplo.com", time.Date(1993, time.June, 15, 0, 0, 0, 0, time.UTC), "contrasena", 1)
	if err != nil {
		log.Fatalf("Error al insertar usuario: %v", err)
	}
}
