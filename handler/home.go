package handler

import (
	"log"
	"net/http"	
	"net/http/httputil"
	"fmt"
)
// HandleHomeRequest a function for when / is called
func HandleHomeRequest(w http.ResponseWriter, req *http.Request) {
	log.Println("Home called")	
	dump, err := httputil.DumpRequest(req, true)
	log.Println(dump);
	if err != nil {
		fmt.Fprintf(w, "%q", dump)
	}

	fmt.Println(string(dump))
}
