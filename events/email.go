package events

import (
	"github.com/gookit/event"
    "github.com/blessium/porking/service"
    "github.com/blessium/porking/events/types"
)

type EmailSubscriber struct {
    emailService service.IEmailService `di.inject:"emailService"`
}


func (e *EmailSubscriber) SubscribedEvents() map[string]interface{} {
	return map[string]interface{}{
		"email.QR": event.ListenerFunc(e.SendQREmail),
        "email.Welcome": event.ListenerFunc(e.SendWelcomeEmail),
	}
}

func (e *EmailSubscriber) SendQREmail(ev event.Event) error {
    qr_data := ev.(*types.EmailQREvent).GetQRInfo() 
    to := ev.(*types.EmailQREvent).GetReceiver()

    if err := e.emailService.SendQREmail(to, qr_data); err != nil {
        return err
    }
	return nil
}

func (e *EmailSubscriber) SendWelcomeEmail(ev event.Event) error {
    return nil
}
