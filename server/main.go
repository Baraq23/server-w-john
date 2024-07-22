package main

import (
	"fmt"
	"net/http"

	"server/handlers"
)

func main() {
	http.HandleFunc("/home", handlers.HomeHandler)
	fmt.Println("server http://localhost:8001")

	http.ListenAndServe(":8001", nil)
}
