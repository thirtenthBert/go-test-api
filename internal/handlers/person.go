package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thirtenthBert/go-test-api/internal/models"
)

var people []models.Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID == params["id"] {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(&models.Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
    var person models.Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    people = append(people, person)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(people)
}