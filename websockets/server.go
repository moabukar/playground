package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all origins (only temporary, dont use in prod!!)
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		// Read message from browser
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		// Print the message to the console
		log.Printf("%s sent: %s\n", ws.RemoteAddr(), string(p))

		// Write message back to browser
		if err := ws.WriteMessage(messageType, p); err != nil {
			return
		}
	}
}

func main() {
	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start the server on localhost port 8080 and log any errors
	log.Println("http server started on :8085")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
