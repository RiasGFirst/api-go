package api

import (
	"api-go/databases"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/", home)
	http.HandleFunc("/user/login", home)
	http.HandleFunc("/user/register", createUser)

	fmt.Printf("[API] Server listen at %s\n", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// Get the JSON body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Make sure the body is not empty and is valid JSON
	if len(body) == 0 {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	// Unmarshal the JSON body into a struct
	
	var temp databases.TempUser
	err = json.Unmarshal(body, &temp)
	if err != nil {
		http.Error(w, "Failed to unmarshal JSON body", http.StatusBadRequest)
		return
	}

	// Create the user
	user, err := databases.CreateUser(&temp)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "User created successfully, ", user)
}
