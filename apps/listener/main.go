package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/smollmegumin/level-TTS/listener/routes"
	"github.com/smollmegumin/level-TTS/listener/routes/twitch"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	http.Handle("/", routes.NewRoute(twitch.NewTwitchRouter()))
	http.ListenAndServe(":8080", nil)

}
