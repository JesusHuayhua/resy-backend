package main

import (
	"log"
	"net/http"
	"soa/pkg/services/ServicioUsuario/api/handlers"
	"soa/pkg/services/ServicioUsuario/core/usecase/backBD"
	"soa/pkg/services/shared/crypto"
	"soa/pkg/services/shared/database"
	"time"

	_ "github.com/lib/pq" //Driver Para base de datos postgreSQL
)

// main es el punto de entrada de la aplicación.
// Aquí se configura la conexión a la base de datos, se inicializa el servicio de usuario
// y se configuran las rutas del servidor HTTP.
// También maneja las solicitudes HTTP para las operaciones CRUD de usuarios y roles,
// así como las operaciones de recuperación de contraseña y autenticación.
// El servidor escucha en el puerto 8080 y maneja las solicitudes de manera concurrente.
// Se utiliza un enrutador HTTP simple para manejar las diferentes rutas y métodos HTTP.
// Además, se implementa un manejador para las solicitudes OPTIONS, lo que permite
// que el servidor responda a las solicitudes de preflight CORS, lo cual es útil
// para aplicaciones web que interactúan con el servidor desde diferentes dominios.
// La configuración de la base de datos y la clave de encriptación se obtienen de una estructura
// de configuración, que se inicializa con los valores necesarios para conectarse a la base de datos
// y para realizar operaciones de encriptación y desencriptación de datos sensibles.
func main() {
	//PONER EN UN .env ambos
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
	servicio := backBD.NuevoServicioUsuario(dbManager.DB, cryptoCtx)
	server := handlers.NewServer(servicio)
	mux := http.NewServeMux()
	// Tus rutas normales
	mux.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handlers.OpcionesHandler(w, r)
			return
		}
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
	mux.HandleFunc("/roles", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handlers.OpcionesHandler(w, r)
			return
		}
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
	mux.HandleFunc("/recuperar", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handlers.OpcionesHandler(w, r)
			return
		}
		if r.Method == http.MethodPost {
			server.IniciarRecuperacionPassword(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/recuperar/confirmar", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handlers.OpcionesHandler(w, r)
			return
		}
		if r.Method == http.MethodPost {
			server.RecuperarPassword(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handlers.OpcionesHandler(w, r)
			return
		}
		if r.Method == http.MethodPost {
			server.Login(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/recuperar/verificar", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handlers.OpcionesHandler(w, r)
			return
		}
		if r.Method == http.MethodPost {
			server.VerificarTokenRecuperacion(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/recuperar/actualizar", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handlers.OpcionesHandler(w, r)
			return
		}
		if r.Method == http.MethodPost {
			server.ActualizarPasswordRecuperacion(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
	// Ruta global para OPTIONS (esto acepta cualquier ruta)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handlers.OpcionesHandler(w, r)
			return
		}
		// ...aquí puedes delegar a otros handlers si lo deseas...
		http.NotFound(w, r)
	})

	// Ejemplo: mostrar a qué hora apunta un time.Time
	now := time.Now()
	log.Printf("Hora local: %v", now)
	log.Printf("Hora UTC: %v", now.UTC())

	log.Println("Servidor escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
