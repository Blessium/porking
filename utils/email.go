package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
	"log"
	"net/smtp"
	"os"
    "net/http"
	"mime/multipart"
)

type Message struct {
	To      string
	From    string
	Subject string
	Body    string
	QRCode  []byte
}

func (m *Message) ToBytes() []byte {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", m.To))
	buf.WriteString(fmt.Sprintf("From: %s\n", m.From))

	buf.WriteString("MIME-Version: 1.0\n")
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()
	buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
	buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	buf.WriteString(m.Body)
	buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
	buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(m.QRCode)))
	buf.WriteString("Content-Transfer-Encoding: base64\n")
	buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", "qr_code.png"))

	b := make([]byte, base64.StdEncoding.EncodedLen(len(m.QRCode)))
	base64.StdEncoding.Encode(b, m.QRCode)
	buf.Write(b)
	buf.WriteString(fmt.Sprintf("\n--%s", boundary))
	buf.WriteString("--")

	return buf.Bytes()
}

func SendEmail(to string, u *model.User, r *model.Reservation, qr_id string) error {
	from := os.Getenv("PORKING_EMAIL")
	pass := os.Getenv("PORKING_PWD")

	host := "smtp.gmail.com"

	msg, err := createResText(from, to, u, r, qr_id)
	if err != nil {
		return err
	}
	auth := smtp.PlainAuth("", from, pass, host)

	if err := smtp.SendMail(host+":587", auth, from, []string{to}, msg); err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func createResText(from string, to string, u *model.User, r *model.Reservation, qr_id string) ([]byte, error) {
	var m Message

	m.To = to
	m.From = from

	db, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	var qr model.QR

	result := db.Limit(1).Find(&qr, "id = ?", qr_id)
	if result.Error != nil {
		return nil, result.Error
	}

	qr_bytes, err := base64.StdEncoding.DecodeString(qr.Image)
	if err != nil {
		return nil, err
	}

	m.Body = "Hello " + u.Name + "\n" + "Confermato blah blah"
	m.Subject = "Your parking reservation"

	m.QRCode = qr_bytes
	return m.ToBytes(), nil
}
