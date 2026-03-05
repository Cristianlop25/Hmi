package sse

import "sync"

type Hub struct {
	clients   map[chan string]bool
	broadcast chan string
	mu        sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:   make(map[chan string]bool),
		broadcast: make(chan string),
	}
}

func (hub *Hub) Broadcast(msg string) {
	hub.broadcast <- msg
}

func (hub *Hub) Run() {
	for msg := range hub.broadcast {
		hub.mu.Lock()
		for client := range hub.clients {
			select {
			case client <- msg:
			default:
			}
		}
		hub.mu.Unlock()
	}
}
