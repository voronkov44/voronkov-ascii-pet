package main

import (
	"log"
	"net/http"
	"voronkov-ascii-pet/internal/pet"
)

func main() {
	router := http.NewServeMux()
	pet.NewPetHandler(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server started at port 8080")
	server.ListenAndServe()
}
