package main

//Pa los que no saben hacer un ruleset
//NO CAMBIEN EL MODULE
// USEN RESI-BACKEND COMO CARPETA RAIZ
import (
	"RESI-BACKEND/services/shared/database"
	"RESI-BACKEND/services/users/internal/repository/bduser"
	"fmt"
	"log"
	"time"
)

func main() {
	dbManager, err := database.NuevoDBManager("conf2/key.json")
	if err != nil {
		log.Fatalf("Error al conectar a la BD: %v", err)
	}
	defer dbManager.Cerrar()
	// Insertar usuario de prueba
	_ = bduser.InsertarNuevoUsuario(dbManager.DB, "Nombre", "Apellido", "correo@ejemplo.com", time.Date(1993, time.June, 15, 0, 0, 0, 0, time.UTC), "contrasena", 1)
	// Actualizar usuario de prueba (por ejemplo, el usuario con id_usuario = 1)
	err = bduser.ActualizarUsuario(dbManager.DB, 1, "NombreActualizado", "ApellidoActualizado", "correo@ejemplo.com", time.Date(1993, time.June, 15, 0, 0, 0, 0, time.UTC), "nueva_contrasena", 1, true)
	if err != nil {
		log.Fatalf("Error al actualizar usuario: %v", err)
	}
	// Seleccionar usuarios (todos)
	usuarios, err := bduser.SeleccionarUsuarios(dbManager.DB, "id_usuario=1")
	if err != nil {
		log.Fatalf("Error al seleccionar usuarios: %v", err)
	}
	for _, usuario := range usuarios {
		fmt.Printf("Usuario: %+v\n", usuario)
	}
}
