package main

import (
	"back/src/bdsede"
	"back/src/db"
	"fmt"
	"log"
)

func main() {
	dbManager, err := db.NewDBManager("config/key.json")
	if err != nil {
		log.Fatalf("Error al conectar a la BD: %v", err)
	}
	err = bdsede.UpdateSede(dbManager.DB, "550e8400-e29b-41d4-a716-446655440000", "Sede Principal", "Av. Principal 123", "Lima", "99936478", 100, "09:00:00", "22:00:00", true)
	if err != nil {
		log.Printf("Error al actualizar sede: %v", err)
	}
	sedesActivas, err := bdsede.SelectSedes(dbManager.DB, "activa = $1", true)
	if err != nil {
		log.Fatalf("Error al obtener sedes activas: %v", err)
	}
	fmt.Println("Sedes activas:", sedesActivas)
	err = bdsede.InsertNewSede(dbManager.DB, "DEV444", "Sede Principal", "Av. Principal 777", "Ayacucho", "3435353", 100, "09:00:00", "22:00:00")
	if err != nil {
		log.Printf("Error al actualizar sede: %v", err)
	}
	dbManager.Close()
}
