package events

import (
	"github.com/gookit/event"
)

func RegisterListeners() {
    event.AddSubscriber(&EmailSubscriber{})
}
