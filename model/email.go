package model

import (
    "fmt"
    "bytes"
	"encoding/base64"
    "net/http"
	"mime/multipart"
)

type EmailQRMessage struct {
	To      string
	From    string
	Subject string
	Body    string
	QRCode  []byte
}

func (m *EmailQRMessage) ToBytes() []byte {
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
