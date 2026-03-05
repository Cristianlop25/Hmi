package terminal

import (
	"fmt"
	"log"

	"hmi-sonic/internal/sse"
)

func Run(hub *sse.Hub) {
	for {
		var input string
		fmt.Scanln(&input)

		switch input {
		case "connectors":
			hub.Broadcast("connectors")
			log.Println("Sending connectors")

		case "identification":
			hub.Broadcast("identification")
			log.Println("Sending identification")

		default:
			log.Println("Unknown command:", input)
		}
	}
}
