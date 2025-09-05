package utils

import (
	"guestbook_backend/config"
	natshelper "guestbook_backend/helper/utils/nats_helper"
)

type Utils struct {
	JaegerTracer *JaegerTracer
	ApiKey       *ApiKey
	HubSse       *HubSse
	Nats         *natshelper.NatsHelper
}

func NewUtils(b *config.NatsBroker) *Utils {
	return &Utils{
		JaegerTracer: NewJaegerTracer(),
		ApiKey:       NewApiKey(),
		HubSse:       NewHubSse(),
		Nats:         natshelper.NewNatsHelper(b),
	}
}
