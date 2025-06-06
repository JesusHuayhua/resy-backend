package main

import (
	"RESI-BACKEND/services/shared/database"
	"log"
)

func main() {
	dbManager, err := database.NewDBManager("conf2/key.json")
	if err != nil {
		log.Fatalf("Error al conectar a la BD: %v", err)
	}
	defer dbManager.Close()
}
