package main

import (
	
	"log"
	"net/http"
	"flag"	
	"github.com/gorilla/websocket"
	
)

//Message the marshalled message from the clients
type Message struct {
	ID    string
	User  string
	Value string
}


var addr = flag.String("addr", ":8080", "http service address")

var userMap map[string]*websocket.Conn

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

