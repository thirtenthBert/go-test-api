package main

import (
	"log"
	"net/http"

	"github.com/thirtenthBert/go-test-api/internal/router"
)

func main() {
    r := router.SetupRouter()
    log.Println("Server running on port 8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}