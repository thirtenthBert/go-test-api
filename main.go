package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Define a struct to represent a resource
type Person struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// In-memory data store
var people []Person

// Handler to get all people
func getPeople(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(people)
}

// Handler to get a single person by ID
func getPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID == params["id"] {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(&Person{})
}

// Handler to create a new person
func createPerson(w http.ResponseWriter, r *http.Request) {
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    people = append(people, person)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(person)
}

// Handler to delete a person by ID
func deletePerson(w http.ResponseWriter, r *http.Request) {
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

func main() {
    // Initialize the router
    r := mux.NewRouter()

    // Mock data
    people = append(people, Person{ID: "1", Name: "John Doe", Age: 30})
    people = append(people, Person{ID: "2", Name: "Jane Doe", Age: 25})

    // Define routes
    r.HandleFunc("/people", getPeople).Methods("GET")
    r.HandleFunc("/people/{id}", getPerson).Methods("GET")
    r.HandleFunc("/people", createPerson).Methods("POST")
    r.HandleFunc("/people/{id}", deletePerson).Methods("DELETE")

    // Start the server
    http.ListenAndServe(":8000", r)
}