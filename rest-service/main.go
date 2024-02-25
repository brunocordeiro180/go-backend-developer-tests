package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/handlers"
	"net/http"
)

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	// TODO: Add RESTful web service here
	r := mux.NewRouter()
	r.HandleFunc("/people", handlers.PeopleHandler).
		Methods("GET")
	r.HandleFunc("/people", handlers.PeopleHandler).
		Queries("first_name", "{first_name}").
		Queries("last_name", "{last_name}")
	r.HandleFunc("/people", handlers.PeopleHandler).
		Queries("phone_number", "{phone_number}")
	r.HandleFunc("/people/{id}", handlers.PeopleByIdHandler)

	http.ListenAndServe(":8010", r)
}
