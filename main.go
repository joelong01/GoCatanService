package main

import (
	"log"
	"net/http"
	"./handler"
)

func main() {
	http.HandleFunc("/", handler.HandleHomeRequest)
	http.HandleFunc("/roll/", handler.HandleRollsRequest)
	http.HandleFunc("/commands/", handler.HandleCommandsRequest)
	http.HandleFunc("/users/", handler.HandleUserRequest)

	var err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln("Server failed starting. Error: ", err)
	}
}