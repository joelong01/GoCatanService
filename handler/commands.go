package handler

import (
	"log"
	"net/http"
	"net/http/httputil"
	"fmt"
	)

// HandleCommandsRequest for route "/command/<command>"
func HandleCommandsRequest(w http.ResponseWriter, req *http.Request) {
	log.Println("commands called")	
	dump, err := httputil.DumpRequest(req, true)
	log.Println(dump);
	if err != nil {
		fmt.Fprintf(w, "%q", dump)
	}

	fmt.Println(string(dump))
}
