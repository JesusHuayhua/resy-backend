package backBD

import (
	"fmt"
	"net/smtp"
	"time"
)

// Credenciales de Salon Verde
const (
	salonVerdeCorreo   = "salonverde620@gmail.com"
	salonVerdePassword = "T7J4N3N44"
	smtpHost           = "smtp.gmail.com"
	smtpPort           = "587"
)

func MensajeRecordatorio(nombre string, fechaHora time.Time) string {
	return fmt.Sprintf(
		"Estimado(a) %s, le enviamos este mensaje para recordarle que su reservación inicia a las %s.",
		nombre,
		fechaHora.Format("02/01/2006 15:04"),
	)
}

func EnviarCorreo(destinatario, nombre string, fechaHora time.Time) error {
	asunto := "Recordatorio de Reserva"
	cuerpo := MensajeRecordatorio(nombre, fechaHora)

	msg := "From: " + salonVerdeCorreo + "\n" +
		"To: " + destinatario + "\n" +
		"Subject: " + asunto + "\n\n" +
		cuerpo

	auth := smtp.PlainAuth("", salonVerdeCorreo, salonVerdePassword, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, salonVerdeCorreo, []string{destinatario}, []byte(msg))
}

func EnviarWhatsApp(numero, nombre string, fechaHora time.Time) error {
	mensaje := MensajeRecordatorio(nombre, fechaHora)
	fmt.Printf("Enviando WhatsApp a %s: %s\n", numero, mensaje)
	return nil
}
