package main

import (
	"encoding/json"
	"log"
	"net/http"

	"./handler"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

//Message the marshalled message from the clients
type Message struct {
	ID    string
	User  string
	Value string
}

func checkOrigin(r *http.Request) bool {
	return true
}

var upgrader = websocket.Upgrader{CheckOrigin: checkOrigin}
var userMap map[string]*websocket.Conn

func main() {
	userMap = make(map[string]*websocket.Conn)
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HandleHomeRequest)
	mux.HandleFunc("/message/", handler.HandleMessageRequest)
	mux.HandleFunc("/ws/", serveWs)

	handler := cors.Default().Handler(mux)
	var err = http.ListenAndServe(":8080", handler)

	if err != nil {
		log.Print("Server failed starting. Error: ", err)
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		mt, msgString, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("[mt=%d] [msg=%s]", mt, msgString)

		//
		//	unmarshal the message to get the user 

		var message Message;
		var jsonErr = json.Unmarshal(msgString, &message)
		if jsonErr != nil {
			writeAll([]byte("error"))
		}

		//
		//	 make sure this user is in the map of users
		var _, ok = userMap[message.User]
		if ok == false {
			userMap[message.User] = conn;
		}
		//
		//	send the message to all the clients (including the one that just sent the message)
		writeAll([]byte(msgString))
		
	}
}

func writeAll(msg []byte) {
	for user, conn := range userMap {
		log.Printf("sending message %s to user %s", msg, user)
		conn.WriteMessage(1, msg)
	}
	
}
