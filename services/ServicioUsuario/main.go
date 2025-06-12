package main

import (
	bduser "ServicioUsuario/persistencia"
	"fmt"
	"log"
	"time"

	"github.com/Shauanth/Singleton_Encription_ServiceGolang/crypton"
	"github.com/Shauanth/Singleton_Encription_ServiceGolang/database"
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
	// Insertar usuario de prueba
	_ = bduser.InsertarNuevoUsuario(dbManager.DB, "Nombre", "Apellido", "correo@ejemplo.com", "+51981923932", time.Date(1993, time.June, 15, 0, 0, 0, 0, time.UTC), "contrasena", 1)
	// Actualizar usuario de prueba (por ejemplo, el usuario con id_usuario = 1)
	err = bduser.ActualizarUsuario(dbManager.DB, 1, "NombreActualizado", "ApellidoActualizado", "correo@ejemplo.com", "+51981923932", time.Date(1993, time.June, 15, 0, 0, 0, 0, time.UTC), "nueva_contrasena", 1, true)
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
