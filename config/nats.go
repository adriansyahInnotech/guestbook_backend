package config

import (
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

type NatsBroker struct {
	Conn *nats.Conn
}

func NewNatsBroker() *NatsBroker {
	nc, err := nats.Connect(
		os.Getenv("URL_NATS"),
		nats.Timeout(5*time.Second),
		nats.ReconnectWait(2*time.Second),
		nats.MaxReconnects(-1), // coba reconnect terus
	)
	if err != nil {
		log.Fatalf("Error connect NATS: %v", err)
	}

	return &NatsBroker{Conn: nc}
}
