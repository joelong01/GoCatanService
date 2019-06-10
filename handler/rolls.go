package handler

import (
	"log"
	"net/http"
	"net/http/httputil"
	"fmt"
	"strconv"
	
)

var newRoll int

// HandleRollsRequest  route handler for /rolls/<roll>
func HandleRollsRequest(w http.ResponseWriter, req *http.Request) {
	roll, err := strconv.Atoi(req.Header.Get("Roll"))
	log.Println("Rolls called ", roll )
	dump, err := httputil.DumpRequest(req, true)
	log.Println(dump);
	if err != nil {
		fmt.Fprintf(w, "%q", dump)
	}

	fmt.Println(string(dump))
	
	if req.Method == "POST" {
		newRoll = roll
		return
	}

	if req.Method == "GET" {
		w.Write([]byte(strconv.Itoa(newRoll)))
		newRoll = -1
	}

	
	
}
