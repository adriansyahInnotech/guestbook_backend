package natshelper

import (
	"fmt"
	"guestbook_backend/config"

	"github.com/nats-io/nats.go"
)

type NatsHelper struct {
	broker *config.NatsBroker
}

func NewNatsHelper(b *config.NatsBroker) *NatsHelper {
	return &NatsHelper{broker: b}
}

// Publish pesan ke device tertentu
func (h *NatsHelper) Publish(deviceID, payload string) error {
	subject := fmt.Sprintf("device.%s", deviceID)
	return h.broker.Conn.Publish(subject, []byte(payload))
}

// Subscribe ke pattern subject
func (h *NatsHelper) Subscribe(subject string, handler nats.MsgHandler) (*nats.Subscription, error) {
	return h.broker.Conn.Subscribe(subject, handler)
}
