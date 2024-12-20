package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Kennedy-lsd/twitchRecipien/config"
)

func main() {
	env := config.NewConfig()

	// get clientId and clientSecret from your twitch developer account
	clientID := env.ClientId
	clientSecret := env.ClientSecret
	email := env.Email
	password := env.Password

	streamerName := os.Args[1]

	token := getOAuthToken(clientID, clientSecret)

	for {
		if checkStreamStatus(streamerName, clientID, token) {
			fmt.Printf("%s is live!\n", streamerName)
			sendEmail(streamerName, email, password)
			break
		}
		fmt.Println("Streamer is offline. Checking again in 5 minutes...")
		time.Sleep(5 * time.Minute)
	}
}
