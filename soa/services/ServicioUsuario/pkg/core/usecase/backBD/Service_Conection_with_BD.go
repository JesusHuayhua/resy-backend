package backBD

import (
	UserModels "ServicioUsuario/pkg/core/domain"
	"ServicioUsuario/pkg/core/internal"
	"ServicioUsuario/pkg/repository"
	"ServicioUsuario/pkg/repository/crypton"
	repoInterface "ServicioUsuario/pkg/repository/interfaces"
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	log "github.com/go-kit/log"
	"gopkg.in/gomail.v2"
)

type ServicioUsuario struct {
	logger      log.Logger
	crud        repoInterface.UserRepository
	cryptConfig crypton.Config
}

// Permite la Conexion de un nuevo usuario para hacer operaciones CRUD con la base de datos
func NuevoServicioUsuario(db *sql.DB, cryptConfig crypton.Config) *ServicioUsuario {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	crud := repository.NewUserRepository(db)
	return &ServicioUsuario{
		logger:      logger,
		crud:        crud,
		cryptConfig: cryptConfig,
	}
}

func (s *ServicioUsuario) Get(_ context.Context, filters ...internal.Filter) (internal.StatusCode, []UserModels.UsuarioBD, error) {
	// Implementar lógica de filtrado si es necesario
	// Por ahora retorna todos los usuarios
	return s.SeleccionarUsuarios("", nil)
}

func (s *ServicioUsuario) Status(_ context.Context, userId string) (internal.StatusCode, error) {
	// Implementar lógica de status si es necesario
	return internal.InProgress, nil
}

func (s *ServicioUsuario) ServiceStatus(_ context.Context) (int, error) {
	s.logger.Log("msg", "Checking status")
	return http.StatusOK, nil
}

// Inserta un nuevo usuario en la base de datos y encripta su contraseña
func (s *ServicioUsuario) InsertarNuevoUsuario(nombres, apellidos, correo, telefono, direccion string, fechaNacimiento time.Time, contrasenia string, rol int) (internal.StatusCode, error) {
	// Encriptar la contraseña antes de guardar
	contraseniaEncriptada, err := crypton.Encrypt(contrasenia, s.cryptConfig)
	if err != nil {
		s.logger.Log("err", fmt.Sprintf("error al encriptar contraseña: %v", err))
		return internal.Error, fmt.Errorf("error al encriptar contraseña: %w", err)
	}
	datos := UserModels.UsuarioVariable{
		Nombres:         nombres,
		Apellidos:       apellidos,
		Correo:          correo,
		Telefono:        telefono,
		Direccion:       direccion,
		FechaNacimiento: fechaNacimiento,
		Contrasenia:     contraseniaEncriptada,
		Rol:             rol,
		EstadoAcceso:    true,
	}
	if err := s.crud.Insertar(`"Usuario"`, datos); err != nil {
		s.logger.Log("err", fmt.Sprintf("error al insertar usuario: %v", err))
		return internal.Error, err
	}
	return internal.InProgress, nil
}

// Actualiza los datos de un usuario existente, encriptando la contraseña si se proporciona una nueva
func (s *ServicioUsuario) ActualizarUsuario(idUsuario int, nombres, apellidos, correo, telefono, direccion string, fechaNacimiento time.Time, contrasenia string, rol int, estado bool) (internal.StatusCode, error) {
	// Encriptar la contraseña si se proporciona una nueva
	contraseniaEncriptada := contrasenia
	if contrasenia != "" {
		var err error
		contraseniaEncriptada, err = crypton.Encrypt(contrasenia, s.cryptConfig)
		if err != nil {
			s.logger.Log("err", fmt.Sprintf("error al encriptar contraseña: %v", err))
			return internal.Error, fmt.Errorf("error al encriptar contraseña: %w", err)
		}
	}
	datos := UserModels.UsuarioVariable{
		Nombres:         nombres,
		Apellidos:       apellidos,
		Correo:          correo,
		Telefono:        telefono,
		Direccion:       direccion,
		FechaNacimiento: fechaNacimiento,
		Contrasenia:     contraseniaEncriptada,
		Rol:             rol,
		EstadoAcceso:    estado,
	}
	where := "id_usuario = $1"
	if err := s.crud.Actualizar(`"Usuario"`, datos, where, idUsuario); err != nil {
		s.logger.Log("err", fmt.Sprintf("error al actualizar usuario: %v", err))
		return internal.Error, err
	}
	return internal.InProgress, nil
}

// Elimina un usuario de la base de datos por su ID de manera LOGICA
func (s *ServicioUsuario) EliminarUsuario(id int) (internal.StatusCode, error) {
	if err := s.crud.Eliminar(`"Usuario"`, fmt.Sprintf("%d", id)); err != nil {
		s.logger.Log("err", fmt.Sprintf("error al eliminar usuario: %v", err))
		return internal.Error, err
	}
	return internal.InProgress, nil
}

// Devuelve una lista de usuarios de la base de datos según una condición opcional
func (s *ServicioUsuario) SeleccionarUsuarios(condicion string, args []interface{}) (internal.StatusCode, []UserModels.UsuarioBD, error) {
	var usuarios []UserModels.UsuarioBD
	columnas := []string{
		"id_usuario", "nombres", "apellidos", "correo",
		"telefono", "direccion", "fechanacimiento", "contrasenia", "rol", "estadoacceso",
	}

	// Asegurar que la condición tenga el formato correcto
	if condicion != "" {
		// Reemplazar todos los ? por $n
		for i := 1; strings.Contains(condicion, "?"); i++ {
			condicion = strings.Replace(condicion, "?", fmt.Sprintf("$%d", i), 1)
		}
	}

	var rows *sql.Rows
	var err error
	if condicion == "" {
		rows, err = s.crud.Seleccionar(`"Usuario"`, columnas, "", args...)
	} else {
		rows, err = s.crud.Seleccionar(`"Usuario"`, columnas, condicion, args...)
	}
	if err != nil {
		s.logger.Log("err", fmt.Sprintf("error en Select: %v", err))
		return internal.Error, nil, fmt.Errorf("error en Select: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var usuario UserModels.UsuarioBD
		var contraseniaEncriptada string
		err := rows.Scan(
			&usuario.IdUsuario,
			&usuario.DataUsuario.Nombres,
			&usuario.DataUsuario.Apellidos,
			&usuario.DataUsuario.Correo,
			&usuario.DataUsuario.Telefono,
			&usuario.DataUsuario.Direccion,
			&usuario.DataUsuario.FechaNacimiento,
			&contraseniaEncriptada,
			&usuario.DataUsuario.Rol,
			&usuario.DataUsuario.EstadoAcceso,
		)
		if err != nil {
			s.logger.Log("err", fmt.Sprintf("error al escanear fila: %v", err))
			return internal.Error, nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		// Desencriptar la contraseña
		contraseniaDescifrada, err := crypton.Decrypt(contraseniaEncriptada, s.cryptConfig)
		if err != nil {
			s.logger.Log("err", fmt.Sprintf("error al descifrar contraseña: %v", err))
			return internal.Error, nil, fmt.Errorf("error al descifrar contraseña: %v", err)
		}
		usuario.DataUsuario.Contrasenia = contraseniaDescifrada
		usuarios = append(usuarios, usuario)
	}
	if err = rows.Err(); err != nil {
		s.logger.Log("err", fmt.Sprintf("error al escanear fila: %v", err))
		return internal.Error, nil, fmt.Errorf("error al escanear fila: %v", err)
	}
	return internal.InProgress, usuarios, nil
}

func (s *ServicioUsuario) InsertarNuevoRol(nombre string) (internal.StatusCode, error) {
	datos := UserModels.Rol{
		NombreRol: nombre,
	}
	if err := s.crud.Insertar(`"Roles"`, datos); err != nil {
		return internal.Error, err
	}
	return internal.InProgress, nil
}

func (s *ServicioUsuario) ActualizarRol(id int, nombre string) (internal.StatusCode, error) {
	datos := UserModels.Rol{
		NombreRol: nombre,
	}
	where := "id_rol = $1"
	if err := s.crud.Actualizar(`"Roles"`, datos, where, id); err != nil {
		return internal.Error, err
	}
	return internal.InProgress, nil
}

func (s *ServicioUsuario) EliminarRol(id int) (internal.StatusCode, error) {
	if err := s.crud.Eliminar(`"Roles"`, fmt.Sprintf("%d", id)); err != nil {
		return internal.Error, err
	}
	return internal.InProgress, nil
}

func (s *ServicioUsuario) SeleccionarRoles(filtro string, params []interface{}) (internal.StatusCode, []UserModels.RolDB, error) {
	var roles []UserModels.RolDB
	columnas := []string{"id_rol", "nombrerol"}
	var rows *sql.Rows
	var err error
	if filtro == "" {
		rows, err = s.crud.Seleccionar(`"Roles"`, columnas, "", params...)
	} else {
		rows, err = s.crud.Seleccionar(`"Roles"`, columnas, filtro, params...)
	}
	if err != nil {
		return internal.Error, nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rol UserModels.RolDB
		err := rows.Scan(
			&rol.ID,
			&rol.DataRol.NombreRol,
		)
		if err != nil {
			return internal.Error, nil, err
		}
		roles = append(roles, rol)
	}
	if err = rows.Err(); err != nil {
		return internal.Error, nil, err
	}
	return internal.InProgress, roles, nil
}

// Inicia el proceso de recuperación de contraseña
func (s *ServicioUsuario) IniciarRecuperacionPassword(correo string) (string, error) {
	// Verifica si el correo existe
	var existe int
	err := s.crud.(*repository.UserRepositoryImpl).Crud().DB.QueryRow(`SELECT COUNT(*) FROM "Usuario" WHERE correo=$1`, correo).Scan(&existe)
	if err != nil || existe == 0 {
		return "", errors.New("correo no registrado")
	}

	// Generar token
	token, err := generarToken()
	if err != nil {
		return "", err
	}

	// Guardar token en la base de datos
	expira := time.Now().Add(15 * time.Minute)
	db := s.crud.(*repository.UserRepositoryImpl).Crud().DB
	if err := guardarTokenRecuperacion(db, correo, token, expira); err != nil {
		return "", err
	}

	// Enviar token por email
	if err := enviarTokenPorEmail(correo, token); err != nil {
		return "", err
	}

	return token, nil
}

// Cambia la contraseña si el token es válido
func (s *ServicioUsuario) RecuperarPassword(correo, token, nuevaContrasenia string) error {
	db := s.crud.(*repository.UserRepositoryImpl).Crud().DB
	var expiraEn time.Time
	var tokenBD string
	err := db.QueryRow(`SELECT token, expira_en FROM "RecuperacionPassword" WHERE correo=$1`, correo).Scan(&tokenBD, &expiraEn)
	if err != nil {
		_ = eliminarToken(db, correo)
		return errors.New("token no encontrado o ya utilizado")
	}
	if tokenBD != token || time.Now().After(expiraEn) {
		_ = eliminarToken(db, correo)
		return errors.New("token inválido o expirado")
	}
	// Encriptar la nueva contraseña
	contraseniaEncriptada, err := crypton.Encrypt(nuevaContrasenia, s.cryptConfig)
	if err != nil {
		return errors.New("no se pudo encriptar la contraseña")
	}
	_, err = db.Exec(`UPDATE "Usuario" SET contrasenia=$1 WHERE correo=$2`, contraseniaEncriptada, correo)
	if err != nil {
		return errors.New("no se pudo actualizar la contraseña")
	}
	_ = eliminarToken(db, correo)
	return nil
}

// Verifica si el token es válido para el correo dado
func (s *ServicioUsuario) VerificarTokenRecuperacion(correo, token string) error {
	db := s.crud.(*repository.UserRepositoryImpl).Crud().DB
	var expiraEn time.Time
	s.logger.Log("debug", fmt.Sprintf("Verificando token para correo: %s", correo))
	s.logger.Log("debug", fmt.Sprintf("Token recibido: %s", token))
	err := db.QueryRow(`SELECT expira_en FROM "RecuperacionPassword" WHERE correo=$1 AND token=$2`, correo, token).Scan(&expiraEn)
	if err != nil {
		s.logger.Log("debug", fmt.Sprintf("Error en SELECT: %v", err))
		_ = eliminarToken(db, correo)
		return errors.New("token no encontrado o ya utilizado")
	}
	s.logger.Log("debug", fmt.Sprintf("Expira en (UTC): %v", expiraEn.UTC()))
	nowUTC := time.Now().UTC()
	s.logger.Log("debug", fmt.Sprintf("Hora actual (UTC): %v", nowUTC))
	if nowUTC.After(expiraEn.UTC()) {
		s.logger.Log("debug", "Token expirado")
		_ = eliminarToken(db, correo)
		return errors.New("token expirado")
	}
	s.logger.Log("debug", "Token válido")
	return nil
}

// Actualiza la contraseña (requiere que el token ya haya sido validado)
func (s *ServicioUsuario) ActualizarPasswordRecuperacion(correo, nuevaContrasenia string) error {
	db := s.crud.(*repository.UserRepositoryImpl).Crud().DB
	contraseniaEncriptada, err := crypton.Encrypt(nuevaContrasenia, s.cryptConfig)
	if err != nil {
		return errors.New("no se pudo encriptar la contraseña")
	}
	_, err = db.Exec(`UPDATE "Usuario" SET contrasenia=$1 WHERE correo=$2`, contraseniaEncriptada, correo)
	if err != nil {
		return errors.New("no se pudo actualizar la contraseña")
	}
	_ = eliminarToken(db, correo)
	return nil
}

func generarToken() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func guardarTokenRecuperacion(db *sql.DB, correo, token string, expira time.Time) error {
	// Guardar la fecha de expiración en UTC para evitar problemas de desfase horario
	_, err := db.Exec(`INSERT INTO "RecuperacionPassword" (correo, token, expira_en) VALUES ($1, $2, $3)`, correo, token, expira.UTC())
	return err
}

func enviarTokenPorEmail(correo, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "salonverde620@gmail.com")
	m.SetHeader("To", correo)
	m.SetHeader("Subject", "Recuperación de contraseña")
	htmlBody := fmt.Sprintf(`
        <html>
        <body>
            <h2 style="color: #2563eb;">Recuperación de contraseña</h2>
            <p>Hemos recibido una solicitud para restablecer tu contraseña.</p>
            <p>Tu código de verificación es: <strong style="font-size: 1.2em;">%s</strong></p>
            <p>Este código expirará en 15 minutos.</p>
            <p style="color: #6b7280; font-size: 0.9em;">
                Si no solicitaste este cambio, por favor ignora este mensaje.<br>
                <strong>Nota:</strong> Este correo fue enviado desde una cuenta común (Gmail), no corporativa.
            </p>
        </body>
        </html>
    `, token)
	m.SetBody("text/html", htmlBody)

	// Usa tu contraseña real o una variable de entorno
	d := gomail.NewDialer("smtp.gmail.com", 587, "salonverde620@gmail.com", "wnvv lqlr niqv hwwt")

	return d.DialAndSend(m)
}

func eliminarToken(db *sql.DB, correo string) error {
	_, err := db.Exec(`DELETE FROM "RecuperacionPassword" WHERE correo=$1`, correo)
	return err
}

// Login verifica las credenciales y retorna la info del usuario (sin contraseña) y estado de acceso
func (s *ServicioUsuario) Login(correo, contrasenia string) (bool, UserModels.UsuarioBD, error) {
	columnas := []string{
		"id_usuario", "nombres", "apellidos", "correo",
		"telefono", "direccion", "fechanacimiento", "contrasenia", "rol", "estadoacceso",
	}
	rows, err := s.crud.Seleccionar(`"Usuario"`, columnas, "correo = $1", correo)
	if err != nil {
		return false, UserModels.UsuarioBD{}, fmt.Errorf("error en select: %v", err)
	}
	defer rows.Close()
	if rows.Next() {
		var usuario UserModels.UsuarioBD
		var contraseniaEncriptada string
		err := rows.Scan(
			&usuario.IdUsuario,
			&usuario.DataUsuario.Nombres,
			&usuario.DataUsuario.Apellidos,
			&usuario.DataUsuario.Correo,
			&usuario.DataUsuario.Telefono,
			&usuario.DataUsuario.Direccion,
			&usuario.DataUsuario.FechaNacimiento,
			&contraseniaEncriptada,
			&usuario.DataUsuario.Rol,
			&usuario.DataUsuario.EstadoAcceso,
		)
		if err != nil {
			return false, UserModels.UsuarioBD{}, fmt.Errorf("error al escanear fila: %v", err)
		}
		// Desencriptar la contraseña
		contraseniaDescifrada, err := crypton.Decrypt(contraseniaEncriptada, s.cryptConfig)
		if err != nil {
			return false, UserModels.UsuarioBD{}, fmt.Errorf("error al descifrar contraseña: %v", err)
		}
		// Comparar contraseñas
		if contraseniaDescifrada == contrasenia {
			usuario.DataUsuario.Contrasenia = "" // No devolver la contraseña
			return true, usuario, nil
		}
	}
	return false, UserModels.UsuarioBD{}, nil
}
