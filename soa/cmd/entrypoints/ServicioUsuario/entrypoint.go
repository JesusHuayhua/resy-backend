package main

import (
	"log"
	"net/http"
	"soa/pkg/services/ServicioUsuario/api/handlers"
	"soa/pkg/services/ServicioUsuario/core/usecase"
	"soa/pkg/services/ServicioUsuario/repository/crypto"
	"soa/pkg/services/ServicioUsuario/repository/database"

	_ "github.com/lib/pq"
)

func main() {
	databaseInformation := database.Config{
		Driver:       "postgres",
		TipoDriver:   "PostgreSQL",
		DBName:       "postgres",
		Host:         "127.0.0.1",
		Port:         "5432",
		User:         "postgres",
		DatabaseName: "ResyDB",
		Password:     "eyJrbXNDaXBoZXJ0ZXh0IjoiQVFJREFIaW4wYXVRQnR4dXppdldKY1ZHVkRMTThIQllFTTVhbFRhWEV3ZlpqZk1XTFFITGlwRjIxQm96a0tvbHBlYTVwUFdYQUFBQWZqQjhCZ2txaGtpRzl3MEJCd2FnYnpCdEFnRUFNR2dHQ1NxR1NJYjNEUUVIQVRBZUJnbGdoa2dCWlFNRUFTNHdFUVFNVTdmQ2thUlc4WnltVzRRQkFnRVFnRHNLSlBwOGtEM0dDQkVEVTlxaEJwNkRpUEoyRTVib0t5V1EvYWo2NGpHU0xMaXMrcEFzbHlobmgvSW5iR0NTMy82OEpjZ1IzWWErblpHNE1RPT0iLCJ3cmFwTm9uY2UiOiJtL2o3a1AyMnNCSWg2RGJwIiwid3JhcHBlZEtleSI6IjRIRFZQUlRiazZnZWVQUFZvaGxVRGlaL3ZpcVRMYUdTbjI3MzhteG1uR1FhakZYVXZackZ6Sk5xVWRBTFVMcFYiLCJwYXlsb2FkTm9uY2UiOiJYWXRadnVMSzVZcVhuTFMzIiwiY2lwaGVydGV4dCI6ImEvaGw1N0g0WDRlYUJ5QUxsT3ZoYlJDQXBRPT0iLCJzYWx0IjoiZXBXSDFUOFRXQTFNRjRQeGxXMHVwUT09IiwiaXRlciI6MTUwMDAwfQ==",
	}
	cryptoCtx, err := crypto.New("alias/resy_master_key", "us-east-1", "prod/crypto_passphrase", 150000)
	if err != nil {
		log.Fatalf("Error al crear contexto de crypto %v", err)
	}
	dbManager, err := database.NuevoDBManager(databaseInformation, cryptoCtx)
	if err != nil {
		log.Fatalf("Error al conectar a la BD: %v", err)
	}
	defer dbManager.Cerrar()
	servicio := usecase.NuevoServicioUsuario(dbManager.DB, cryptoCtx)
	server := handlers.NewServer(servicio)
	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			server.InsertarUsuario(w, r)
		case http.MethodGet:
			server.ListarUsuarios(w, r)
		case http.MethodPut:
			server.ActualizarUsuario(w, r)
		case http.MethodDelete:
			server.EliminarUsuario(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/roles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			server.InsertarRol(w, r)
		case http.MethodGet:
			server.ListarRoles(w, r)
		case http.MethodPut:
			server.ActualizarRol(w, r)
		case http.MethodDelete:
			server.EliminarRol(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/recuperar", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			server.IniciarRecuperacionPassword(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/recuperar/confirmar", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			server.RecuperarPassword(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
	log.Println("Servidor escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
