package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func PeopleHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	firstName := query.Get("first_name")
	lastName := query.Get("last_name")
	phoneNumber := query.Get("phone_number")

	var people []*models.Person

	if firstName != "" && lastName != "" {
		people = models.FindPeopleByName(firstName, lastName)
	} else if phoneNumber != "" {
		people = models.FindPeopleByPhoneNumber(phoneNumber)
	} else {
		people = models.AllPeople()
	}

	jsonResponse, jsonError := json.Marshal(people)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func PeopleByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	person, err := models.FindPersonByID(uuid.Must(uuid.FromString(id)))

	if err != nil {
		jsonResponse, _ := json.Marshal(ErrorResponse{Error: err.Error()})
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonResponse)
		return
	}

	jsonResponse, _ := json.Marshal(person)
	w.Write(jsonResponse)
}
