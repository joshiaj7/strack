package controller

import (
	"log"
	"net/http"
)

// SetupRoute handles routing for API request
func SetupRoute() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	log.Println("This is root")
}
