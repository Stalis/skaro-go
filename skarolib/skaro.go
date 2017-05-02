package skaro

import (
	"log"

	//"bitbucket.org/bodhisnarkva/cbor/go"
	//"github.com/gorilla/websocket"
)

func StartConnection(host string) {
	log.Println("Opening WebSocket connect...")
	var connection Connection
	connection.start(host)
}
