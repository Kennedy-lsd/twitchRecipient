package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/smtp"
)

func getOAuthToken(clientID, clientSecret string) string {
	url := "https://id.twitch.tv/oauth2/token"
	data := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=client_credentials", clientID, clientSecret)

	resp, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Fatalf("Error getting OAuth token: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	return result["access_token"].(string)
}

func checkStreamStatus(username, clientID, token string) bool {

	url := fmt.Sprintf("https://api.twitch.tv/helix/streams?user_login=%s", username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Client-ID", clientID)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to fetch stream status: %s\n", resp.Status)
		return false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return false
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return false
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return false
	}

	return true
}

func sendEmail(streamerName, email, password string) {

	//sending emails via Gmail's SMTP servers
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", email, password, smtpHost)

	postContent := "https://www.twitch.tv/" + streamerName

	subject := "Subject: Hello from TwitchRecipient!"
	body := fmt.Sprintf("Good news! %s is now live on Twitch\n\nLink: %s", streamerName, postContent)

	message := []byte(subject + "\n" + body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email, []string{email}, message)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	fmt.Println("Post email notification sent successfully")
}
