package utils

import (
	"log"
)

type HubSse struct {
	clients map[string]chan string
}

func NewHubSse() *HubSse {
	return &HubSse{
		clients: make(map[string]chan string),
	}
}

// Register device, return buffered channel
func (h *HubSse) Register(deviceID string) chan string {
	if oldCh, ok := h.clients[deviceID]; ok {
		close(oldCh)
	}
	ch := make(chan string, 10)
	h.clients[deviceID] = ch
	log.Printf("HubSse: Registered device %s", deviceID)
	return ch
}

// Unregister device
func (h *HubSse) Unregister(deviceID string) {
	if ch, ok := h.clients[deviceID]; ok {
		close(ch)
		delete(h.clients, deviceID)
		log.Printf("HubSse: Unregistered device %s", deviceID)
	}
}

// Send message to device
func (h *HubSse) Send(deviceID, msg string) {
	if ch, ok := h.clients[deviceID]; ok {
		select {
		case ch <- msg:
		default:
			log.Printf("HubSse: Channel full for device %s", deviceID)
		}
	}
}
