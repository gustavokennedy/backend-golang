// /*
// ______      ________ _____            _      _        _____ _      ____  _    _ _____
// / __ \ \    / /  ____|  __ \     /\   | |    | |      / ____| |    / __ \| |  | |  __ \
// | |  | \ \  / /| |__  | |__) |   /  \  | |    | |     | |    | |   | |  | | |  | | |  | |
// | |  | |\ \/ / |  __| |  _  /   / /\ \ | |    | |     | |    | |   | |  | | |  | | |  | |
// | |__| | \  /  | |____| | \ \  / ____ \| |____| |____ | |____| |___| |__| | |__| | |__| |
// \____/   \/   |______|_|  \_\/_/    \_\______|______(_)_____|______\____/ \____/|_____/
// overall.cloud
// */

package middlewares

import (
	"log"
	"os"

	gomail "gopkg.in/mail.v2"
)

type EmailMiddleware struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
}

func NovoEmailMiddleware() *EmailMiddleware {
	// Obter as variáveis de ambiente
	host := os.Getenv("SMTP_HOST")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	return &EmailMiddleware{
		SMTPHost:     host,
		SMTPPort:     465,
		SMTPUsername: username,
		SMTPPassword: password,
	}
}

func (em *EmailMiddleware) SendEmail(to, subject, body string) error {
	// Configurar o objeto de mensagem de email
	msg := gomail.NewMessage()
	msg.SetHeader("From", em.SMTPUsername)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)

	// Configurar as credenciais SMTP
	d := gomail.NewDialer(em.SMTPHost, em.SMTPPort, em.SMTPUsername, em.SMTPPassword)

	// Enviar o email
	if err := d.DialAndSend(msg); err != nil {
		return err
	}
	// Colocar usuário que enviou no log
	log.Println("E-mail enviado!")
	return nil
}
