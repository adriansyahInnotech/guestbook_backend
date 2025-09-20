package controllers

import (
	"fmt"
	"guestbook_backend/helper/utils/hub"
	"log"
	"net/http"
	"time"
)

type Controllers struct {
	hubSSe *hub.HubSse
}

func NewControllers(hub *hub.HubSse) *Controllers {
	return &Controllers{
		hubSSe: hub,
	}
}

func (c *Controllers) ScanVisitorHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		// Register device
		msgChan := c.hubSSe.Register(deviceID)
		defer c.hubSSe.Unregister(deviceID)

		// kirim event awal agar koneksi terasa cepat
		fmt.Fprintf(w, "event: open\ndata: connected\n\n")
		flusher.Flush()

		ticker := time.NewTicker(15 * time.Second)
		defer ticker.Stop()

		notify := r.Context().Done()
		for {
			select {
			case <-notify:
				log.Printf("SSE closed for %s", deviceID)
				return
			case <-ticker.C:
				// heartbeat
				fmt.Fprintf(w, ": ping\n\n")
				flusher.Flush()
			case msg := <-msgChan:
				fmt.Fprintf(w, "data: %s\n\n", msg)
				flusher.Flush()
			}
		}
	}
}
