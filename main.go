package main

import (
	handler "example/vercelgo/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
