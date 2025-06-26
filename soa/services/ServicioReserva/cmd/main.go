package main

import (
	"ServicioReserva/pkg/api/handlers"
	"ServicioReserva/pkg/core/usecase/backBD"
	"ServicioReserva/pkg/repository/crypton"
	"ServicioReserva/pkg/repository/database"
	"log"
	"net/http"
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
		Password:     "WwF3OBYuf8Tx1opemwPSc4LrAMv2NDQLZ/mYh4HPwcVZymIShg==",
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
			// Buscar reservas que ocurren en 2 horas (rango de 4 minutos)
			condicion := `fecha_reservada BETWEEN NOW() + INTERVAL '1 hour 58 minutes' AND NOW() + INTERVAL '2 hour 2 minutes'`
			reservas, err := servicio.ListarReservas(condicion)
			if err == nil {
				for _, r := range reservas {
					data := r.DataReserva
					if data.CorreoCliente != "" {
						_ = backBD.EnviarCorreo(data.CorreoCliente, data.NombreCliente, data.FechaReservada)
					}
					if data.TelefonoCliente != "" {
						_ = backBD.EnviarWhatsApp(data.TelefonoCliente, data.NombreCliente, data.FechaReservada)
					}
				}
			}
			time.Sleep(5 * time.Minute)
		}
	}()
	log.Println("ServicioReserva escuchando en :8082")
	log.Fatal(http.ListenAndServe(":8082", mux))
}
