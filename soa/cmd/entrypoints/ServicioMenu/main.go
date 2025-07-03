package main

import (
	"log"
	"net/http"
	"soa/pkg/services/ServicioMenu/api/handlers"
	"soa/pkg/services/ServicioMenu/core/usecase/backBD"
	"soa/pkg/services/shared/crypto"
	"soa/pkg/services/shared/database"
)

func main() {
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
	cryptoCtx, err := crypto.New("alias/resy_master_key", "us-east-1", "prod/crypto_passphrase", 150000)
	if err != nil {
		log.Fatalf("Error al crear contexto de crypto %v", err)
	}
	dbManager, err := database.NuevoDBManager(databaseInformation, cryptoCtx)
	if err != nil {
		log.Fatalf("Error al conectar a la BD: %v", err)
	}
	defer dbManager.Cerrar()
	servicio := backBD.NuevoServicioMenu(dbManager.DB)
	server := handlers.NewServer(servicio)
	mux := http.NewServeMux()
	mux.HandleFunc("/platos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handlers.OpcionesHandler(w, r)
			return
		}
		switch r.Method {
		case http.MethodPost:
			server.InsertarPlato(w, r)
		case http.MethodGet:
			server.ListarPlatos(w, r)
		case http.MethodPut:
			server.ActualizarPlato(w, r)
		case http.MethodDelete:
			server.EliminarPlato(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// Rutas para categorías
	mux.HandleFunc("/categorias", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handlers.OpcionesHandler(w, r)
			return
		}
		switch r.Method {
		case http.MethodPost:
			server.InsertarCategoria(w, r)
		case http.MethodGet:
			server.ListarCategorias(w, r)
		case http.MethodPut:
			server.ActualizarCategoria(w, r)
		case http.MethodDelete:
			server.EliminarCategoria(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// Rutas para menú semanal
	mux.HandleFunc("/menu", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handlers.OpcionesHandler(w, r)
			return
		}
		switch r.Method {
		case http.MethodPost:
			server.InsertarMenuSemanal(w, r)
		case http.MethodGet:
			server.ListarMenusSemanales(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// Ruta para obtener menú semanal completo
	mux.HandleFunc("/menu/completo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handlers.OpcionesHandler(w, r)
			return
		}
		if r.Method == http.MethodGet {
			server.ObtenerMenuSemanalCompleto(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// Rutas para días del menú
	mux.HandleFunc("/menudia", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handlers.OpcionesHandler(w, r)
			return
		}
		switch r.Method {
		case http.MethodPost:
			server.InsertarMenudia(w, r)
		case http.MethodGet:
			server.ListarDiasDeMenu(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// Rutas para platos en menú diario
	mux.HandleFunc("/platosenmenudia", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handlers.OpcionesHandler(w, r)
			return
		}
		switch r.Method {
		case http.MethodPost:
			server.InsertarPlatoEnMenudia(w, r)
		case http.MethodGet:
			server.ListarPlatosEnMenudia(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	log.Println("ServicioMenu escuchando en :8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
