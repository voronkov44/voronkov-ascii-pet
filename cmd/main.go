// @title        Pet API
// @version      1.0
// @description  REST API для работы с pet ascii картинками
// @host         localhost:8080
// @BasePath     /v1
// @schemes      http
// @contact.name   Andrew
// @contact.url    https://t.me/voronkov44
// @contact.email  voronkovworkemail@gmail.com

package main

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"

	_ "voronkov-ascii-pet/docs"

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
	router.Handle("/swagger/", httpSwagger.WrapHandler)
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
