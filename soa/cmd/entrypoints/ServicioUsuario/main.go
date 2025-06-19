package entrypoints

import (
	"log"
	"net/http"
	"soa/pkg/services/ServicioUsuario/api/handlers"
	"soa/pkg/services/ServicioUsuario/core/usecase/backBD"
	"soa/pkg/services/ServicioUsuario/repository/crypton"
	"soa/pkg/services/ServicioUsuario/repository/database"

	_ "github.com/lib/pq" //Driver Para base de datos postgreSQL
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
	encriptacionKey := crypton.Config{
		EncryptionKey: "53WDFETRFQFC1?*OS!0LNSADJUER2YU8",
		Salt:          "RCumoV7j",
	}
	dbManager, err := database.NuevoDBManager(databaseInformation, encriptacionKey)
	if err != nil {
		log.Fatalf("Error al conectar a la BD: %v", err)
	}
	defer dbManager.Cerrar()
	servicio := backBD.NuevoServicioUsuario(dbManager.DB, encriptacionKey)
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
