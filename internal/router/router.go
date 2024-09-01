package router

import (
	"github.com/gorilla/mux"
	"github.com/thirtenthBert/go-test-api/internal/handlers"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    // Mock data
    handlers.People = append(handlers.People, handlers.Person{ID: "1", Name: "John Doe", Age: 30})
    handlers.People = append(handlers.People, handlers.Person{ID: "2", Name: "Jane Doe", Age: 25})

    r.HandleFunc("/people", handlers.GetPeople).Methods("GET")
    r.HandleFunc("/people/{id}", handlers.GetPerson).Methods("GET")
    r.HandleFunc("/people", handlers.CreatePerson).Methods("POST")
    r.HandleFunc("/people/{id}", handlers.DeletePerson).Methods("DELETE")

    return r
}