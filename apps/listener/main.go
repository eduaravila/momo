package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/smollmegumin/level-TTS/listener/routes"
)

func init() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", routes.Routes())
	http.ListenAndServe(":8080", mux)

}
