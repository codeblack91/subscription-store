package main

import (
	"fmt"
	"log"
	"net/http"
	"subscription-store/internal/api/routes"
)

// @title Example API
// @version 1.0
// @description This is a simple API documentation example
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := routes.NewRouter()
	fmt.Println("Сервер запущен на порте :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
