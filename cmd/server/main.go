package main

import (
	"log"
	"net/http"

	"hmi-sonic/internal/render"
	"hmi-sonic/internal/router"
	"hmi-sonic/internal/sse"
	"hmi-sonic/internal/terminal"
)

func main() {
	hub := sse.NewHub()

	mux := http.NewServeMux()
	mux.Handle("/app.css", http.FileServer(http.Dir("public")))
	mux.Handle("/sprite.svg", http.FileServer(http.Dir("assets")))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	mux.Handle("/", router.New(render.Render))
	mux.HandleFunc("/events", hub.Handler)

	go hub.Run()
	go terminal.Run(hub)

	log.Println("http://localhost:8081")
	log.Fatal(http.ListenAndServe("localhost:8081", mux))
}
