package main

import (
	"fmt"
	"net/http"

	"backend/internal/config"
	"backend/internal/controller"
)

func main() {
	err := config.SetupDB()
	if err != nil {
		fmt.Println(err)
	}

	controller.SetupRoute()
	http.ListenAndServe(":8080", nil)

	fmt.Println("Listening on port 8080")
}
