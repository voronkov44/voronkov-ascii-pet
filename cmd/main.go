package main

import (
	"log"
	"net/http"
	"os"
	"voronkov-ascii-pet/internal/pet"
)

func main() {
	logFile, err := os.OpenFile("ascii-pet.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	router := http.NewServeMux()
	pet.NewPetHandler(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server started at port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("Failed to start server:", err)
	}
}
