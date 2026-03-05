package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"

	"hmi-sonic/internal/router"
)

var (
	clients   = make(map[chan string]bool)
	broadcast = make(chan string)
	mu        sync.Mutex
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app.css", http.FileServer(http.Dir("public")))
	mux.Handle("/sprite.svg", http.FileServer(http.Dir("assets")))
	mux.Handle("/", router.New(render))
	mux.HandleFunc("/events", sseHandler)

	go broadcastLoop()
	go terminalLoop()

	log.Println("http://localhost:8081")
	log.Fatal(http.ListenAndServe("localhost:8081", mux))
}

func render(w http.ResponseWriter, name string, data any) {
	tpl := template.Must(template.ParseFiles(
		"views/layouts/base.html",
		"views/pages/"+name+".html",
	))

	if err := tpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	messageChan := make(chan string, 1)

	mu.Lock()
	clients[messageChan] = true
	mu.Unlock()

	defer func() {
		mu.Lock()
		delete(clients, messageChan)
		mu.Unlock()
		close(messageChan)
	}()

	for msg := range messageChan {
		fmt.Fprintf(w, "event: message\ndata: {url:%s, message: redirecting to %s...}\n\n", msg, msg)
		flusher.Flush()
	}
}

func broadcastLoop() {
	for msg := range broadcast {
		mu.Lock()
		for client := range clients {
			select {
			case client <- msg:
			default:
			}
		}
		mu.Unlock()
	}
}

func terminalLoop() {
	for {
		var input string
		fmt.Scanln(&input)
		switch input {
		case "connectors":
			broadcast <- "/connectors"
			log.Println("Sending connectors")
		case "identification":
			broadcast <- "/identification"
			log.Println("Sending identification")
		default:
			log.Println("Unknown command:", input)
		}
	}
}
