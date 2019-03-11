package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/subosito/twilio"
)

var accountSid, authToken, receiver, sender string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accountSid = os.Getenv("SID")
	authToken = os.Getenv("TOKEN")
	receiver = os.Getenv("RECEIVER")
	sender = os.Getenv("SENDER")

	if accountSid == "" {
		log.Fatal("SID need to be set", err)
	}
	if authToken == "" {
		log.Fatal("TOKEN need to be set", err)
	}
	if receiver == "" {
		log.Fatal("RECEIVER need to be set", err)
	}
	if sender == "" {
		log.Fatal("SENDER need to be set", err)
	}

	quotes := []string{
		"I urge you to please notice when you are happy, and exclaim or murmur or think at some point, 'If this isn't nice, I don't know what is.'",
		"Peculiar travel suggestions are dancing lessons from God.",
		"There's only one rule that I know of, babies—God damn it, you've got to be kind.",
		"Many people need desperately to receive this message: 'I feel and think much as you do, care about many of the things you care about, although most people do not care about them. You are not alone.'",
		"That is my principal objection to life, I think: It's too easy, when alive, to make perfectly horrible mistakes.",
		"So it goes.",
		"We must be careful about what we pretend to be.",
	}

	rand.Seed(time.Now().Unix())

	c := twilio.NewClient(accountSid, authToken, nil)

	params := twilio.MessageParams{
		Body: quotes[rand.Intn(len(quotes))],
	}
	s, resp, err := c.Messages.Send(sender, receiver, params)
	if err != nil {
		log.Fatal("Err:", err)
	}
	log.Printf("Message Sent: %v\n", s)
	log.Printf("Response: %v\n", resp)

}
