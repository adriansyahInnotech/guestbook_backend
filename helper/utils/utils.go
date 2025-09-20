package utils

import (
	"guestbook_backend/config"
	"guestbook_backend/helper/utils/hub"
	natshelper "guestbook_backend/helper/utils/nats_helper"
)

type Utils struct {
	JaegerTracer *JaegerTracer
	ApiKey       *ApiKey
	HubSse       *hub.HubSse
	Nats         *natshelper.NatsHelper
}

func NewUtils(b *config.NatsBroker) *Utils {
	return &Utils{
		JaegerTracer: NewJaegerTracer(),
		ApiKey:       NewApiKey(),
		HubSse:       hub.NewHubSse(),
		Nats:         natshelper.NewNatsHelper(b),
	}
}
