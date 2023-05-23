package service

import (
	"net/smtp"
	"os"
    "github.com/blessium/porking/model"
    "fmt"
)

type IEmailService interface {
	SendWelcomeEmail(to string) error
	SendQREmail(to string, qr []byte) error
}

type EmailService struct {
	from string
	pass string
	host string
}

func (e EmailService) initEmailClient() {
	e.from = os.Getenv("PORKING_EMAIL")
	e.pass = os.Getenv("PORKING_PWD")
    e.host ="smtp.gmail.com"
}

func (e EmailService) SendQREmail(to string, qr []byte) error {
	e.from = os.Getenv("PORKING_EMAIL")
	e.pass = os.Getenv("PORKING_PWD")
    e.host ="smtp.gmail.com"
    fmt.Println("Bruh")
    auth := smtp.PlainAuth("", e.from, e.pass, e.host)
    message := model.EmailQRMessage {
        To: to,
        From: e.from,
        Subject: "Registrazione effettuata",
        Body: "Registrazione effettuata",
        QRCode: qr,
    }

    if err := smtp.SendMail(e.host + ":587", auth, e.from, []string{to}, message.ToBytes()); err != nil {
        fmt.Println("Error fucking sengin")
        return err
    }

    return nil
}

func (e EmailService) SendWelcomeEmail(to string) error {
    return nil
}
