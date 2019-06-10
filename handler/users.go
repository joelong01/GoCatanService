package handler

import (
	"log"
	"net/http"	
	"net/http/httputil"
	"fmt"
	"encoding/json"

)

//User the model for users
type User struct {
	ID string `json:"id"`
	Name string `json:"Name"`
	Order int `json:"Order"`
}

// HandleUserRequest a function for when / is called
func HandleUserRequest(w http.ResponseWriter, req *http.Request) {	
	log.Println("/users/ called")	
	dump, err := httputil.DumpRequest(req, true)
	log.Println(dump);	
	if err != nil {
		fmt.Fprintf(w, "%q", dump)
	}

	if req.Method == "GET" {

		users := []User {
			User {
				ID: "0x1234",
				Name: "Joe",
				Order: 0,
			},
			User {
				ID: "0x1235",
				Name: "Dodgy",
				Order: 1,
			},
			User {
				ID: "0x1236",
				Name: "Doug",
				Order: 2,
			},
			User {
				ID: "0x1237",
				Name: "Robert",
				Order: 3,
			},
			User {
				ID: "0x1238",
				Name: "Chris",
				Order: 4,
			},

		}
		
		out, err := json.Marshal(users)
		if err != nil{
			w.Write([]byte((err.Error())))
		} else {
			w.Write([]byte(out))			
		}
		return
	}

	if req.Method != "POST" {
		log.Println("only POST and GET are accepted for users!")
		return
	}


	fmt.Println(string(dump))
}
