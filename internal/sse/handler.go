package sse

import (
	"fmt"
	"net/http"
)

func (hub *Hub) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	messageChan := make(chan string, 1)

	hub.mu.Lock()
	hub.clients[messageChan] = true
	hub.mu.Unlock()

	defer func() {
		hub.mu.Lock()
		delete(hub.clients, messageChan)
		hub.mu.Unlock()
		close(messageChan)
	}()

	for msg := range messageChan {
		fmt.Fprintf(w, "event: %s\ndata: %s\n\n", msg, msg)
		flusher.Flush()
	}
}
