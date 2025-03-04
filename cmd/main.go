package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ahmadammarm/go-rest-api/config"
	"github.com/ahmadammarm/go-rest-api/routes"
)


func main() {

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Tidak bisa connect ke database: %v", err)
	}

    routes := routes.SetupRoutes(db)
    fmt.Println("Server dijalankan pada port 8080")
    http.ListenAndServe(":8080", routes)
}