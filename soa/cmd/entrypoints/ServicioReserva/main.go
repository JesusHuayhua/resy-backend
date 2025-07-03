package main

import (
	"log"
	"net/http"
	"soa/pkg/services/ServicioReserva/api/handlers"
	"soa/pkg/services/ServicioReserva/core/usecase/backBD"
	"soa/pkg/services/shared/crypto"
	"soa/pkg/services/shared/database"
	"time"
)

func main() {
	// Configuración de la BD (ajusta según tu entorno)
	databaseInformation := database.Config{
		Driver:       "postgres",
		TipoDriver:   "PostgreSQL",
		DBName:       "ingesoft1",
		Host:         "ingesoft1.cyofngbo9tfh.us-east-1.rds.amazonaws.com",
		Port:         "5432",
		User:         "ingesoft1",
		DatabaseName: "ResyDB",
		Password:     "eyJrbXNDaXBoZXJ0ZXh0IjoiQVFJREFIaW4wYXVRQnR4dXppdldKY1ZHVkRMTThIQllFTTVhbFRhWEV3ZlpqZk1XTFFGbml5cWdnZmRTQ1RsUEE4YkdvZm9IQUFBQWZqQjhCZ2txaGtpRzl3MEJCd2FnYnpCdEFnRUFNR2dHQ1NxR1NJYjNEUUVIQVRBZUJnbGdoa2dCWlFNRUFTNHdFUVFNUkRMdk1OVlp3UnJNcDNHL0FnRVFnRHY4RlFuOW0zM2RLdlRRTzN0YzlFQTFBKzVSRXFPZ1BJWDdRRThCS3F5YzJwak41K1NPd2x4elhuNU5yVTRKa0JLREtzaTZ0N1RwZlJ4d3pnPT0iLCJ3cmFwTm9uY2UiOiI1YkRnSElPWElta3dXMUVEIiwid3JhcHBlZEtleSI6IjBqandHblhDZXNKWmxVNytWUFMxRHAvK2hCeDBOK1lKT2pYUUVuNzdJcjdIdjhWcEowQ052YlFJSW5nM0pwUTciLCJwYXlsb2FkTm9uY2UiOiI1M1NBWWIzUnRCSkU0SjI0IiwiY2lwaGVydGV4dCI6InlNaVRzZVBiSzgyM1dCaER2Rkx5Qzd4Z2duOUNaOVVwaFE9PSIsInNhbHQiOiJzN1RyNHRHUEJOZWpraERRRW4rNW53PT0iLCJpdGVyIjoxNTAwMDB9",
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

	servicio := backBD.NuevoServicioReserva(dbManager.DB)
	server := handlers.NewServer(servicio)
	mux := http.NewServeMux()
	mux.HandleFunc("/reservas", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handlers.OpcionesHandler(w, r)
			return
		}
		switch r.Method {
		case http.MethodPost:
			server.InsertarReserva(w, r)
		case http.MethodGet:
			server.ListarReservas(w, r)
		case http.MethodPut:
			server.ActualizarReserva(w, r)
		case http.MethodDelete:
			server.EliminarReserva(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
	// Concurrencia necesaria para recordatorios
	go func() {
		for {
			reservas, err := servicio.ListarReservasParaRecordatorio()
			if err == nil {
				for _, r := range reservas {
					data := r.DataReserva
					// Usar .Valid para verificar si el valor no es NULL
					if data.CorreoCliente.Valid && data.CorreoCliente.String != "" {
						// Acceder al valor real con .String
						_ = backBD.EnviarCorreo(
							data.CorreoCliente.String,
							data.NombreCliente.String, // También debe ser NullString
							data.FechaReservada,
						)
					}
					if data.TelefonoCliente.Valid && data.TelefonoCliente.String != "" {
						_ = backBD.EnviarWhatsApp(
							data.TelefonoCliente.String,
							data.NombreCliente.String, // También debe ser NullString
							data.FechaReservada,
						)
					}
				}
			}
			time.Sleep(5 * time.Minute)
		}
	}()

	log.Println("ServicioReserva escuchando en :8082")
	log.Fatal(http.ListenAndServe(":8082", mux))
}
