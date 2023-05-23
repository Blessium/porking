package types

import (
    "github.com/gookit/event"
)

type EmailQREvent struct {
	event.BasicEvent
    Qr_data []byte
    Receiver string
}

func (e *EmailQREvent) GetQRInfo() []byte {
    return e.Qr_data
}

func (e *EmailQREvent) GetReceiver() string {
    return e.Receiver
}
