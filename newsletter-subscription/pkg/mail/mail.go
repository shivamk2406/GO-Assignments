package mail

import (
	"crypto/tls"
	"log"

	gomail "gopkg.in/gomail.v2"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
}

type NewMail struct {
	To          string
	From        string
	Subject     string
	Description string
}

type Mailer struct {
	Dialer *gomail.Dialer
}

func NewMailConn(cfg Config) *Mailer {
	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &Mailer{Dialer: d}
}

func (m Mailer) SendMail(mail NewMail) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", mail.From)
	msg.SetHeader("To", mail.To)
	msg.SetHeader("Subject", mail.Subject)
	msg.SetBody("text/plain", mail.Description)
	if err := m.Dialer.DialAndSend(msg); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
