package handler

import (
	"fmt"
	"net"
	"log"
	"net/http"
	"strconv"
	"encoding/json"
	"sync/atomic"
	
)
// CatanMessage the JSON formatted message passed in the header of a POST
type CatanMessage struct {
	ID string 
	Value string
}

// GlobalMessage the threadsafe storage of the last message sent to the relay
var GlobalMessage atomic.Value;

// HandleMessageRequest  route handler for /rolls/<roll>
func HandleMessageRequest(w http.ResponseWriter, req *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	
	if req.Method == "POST" {		
		message := CatanMessage{req.Header.Get("id"), req.Header.Get("value")}
		log.Printf("POST %v", message)
		GlobalMessage.Store(message);	
		return
	}

	if req.Method == "GET" {
		message := GlobalMessage.Load().(CatanMessage)
		log.Printf("GET %v", message)
		data, err := json.Marshal(message)
		if err != nil {
			log.Panicln("error marshaling struct: " + string(err.Error()))
		}
		w.Write(data)
		GlobalMessage.Store(CatanMessage{"-1", ""})
	}

}


func writeToServer(roll int, w http.ResponseWriter) {
	conn, err := net.Dial("tcp", "localhost:1337")
	if err != nil {
		fmt.Fprintf(conn, "Roll="+strconv.Itoa(roll))
	} else {
		fmt.Fprintf(w, "%q", err)
	}
}
