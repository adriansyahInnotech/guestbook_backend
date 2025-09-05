package controllers

import (
	"fmt"
	natshelper "guestbook_backend/helper/utils/nats_helper"
	"log"
	"net/http"
	"time"

	"github.com/nats-io/nats.go"
)

type Controllers struct {
	natsHelper *natshelper.NatsHelper
}

func NewControllers(natsHelper *natshelper.NatsHelper) *Controllers {
	return &Controllers{natsHelper: natsHelper}
}

func (s *Controllers) ScanVisitorHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		// ------------------

		deviceID := r.URL.Query().Get("device_id")
		if deviceID == "" {
			http.Error(w, "device_id required", http.StatusBadRequest)
			return
		}

		// SSE headers
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}

		// Subscribe ke NATS
		subject := fmt.Sprintf("device.%s", deviceID)
		sub, err := s.natsHelper.Subscribe(subject, func(msg *nats.Msg) {
			fmt.Fprintf(w, "data: %s\n\n", string(msg.Data))
			flusher.Flush()
		})
		if err != nil {
			http.Error(w, "failed to subscribe", http.StatusInternalServerError)
			return
		}

		defer sub.Unsubscribe()

		// Heartbeat supaya koneksi tetap hidup
		ticker := time.NewTicker(15 * time.Second)
		defer ticker.Stop()

		notify := r.Context().Done()
		for {
			select {
			case <-notify:
				log.Printf("SSE closed for device %s\n", deviceID)
				return
			case <-ticker.C:
				fmt.Fprintf(w, ": ping\n\n")
				flusher.Flush()
			}
		}
	}
}
