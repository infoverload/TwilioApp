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

	words := []string{
		"sturmfrei (adjective): Literally »storm free«. Comes from Sturmfreiheit, describing that a castle or fort is protected against attackers. Sturmfrei describes a family’s apartment/house when the parents are gone overnight and their teenagers see opportunity for a party.",
		"Wegbier, das: Literally »Beer for the way«. Describes the taking a beer to drink on the way to go somewhere, usually to a party or event.",
		"Saftladen, der: Literally »Juice store«. Derogatory. Describes a really badly organised company with incapable employees/managers. The German equivalent to lemonade stand. Comes from the idea that a juice store is the easiest kind of producing company.",
		"Morgenmuffel, der: Literally »Morning bad/mouldy air personified noun«. Describes someone who is grumpy in the morning/does not like to get up early/quickly or needs a lot of time/coffee to get up to speed in the morning.",
		"Schnapsidee, die: Literally »hard liquor idea«. Describes an idea whose realisation is unrealistic. Comes from the phenomenon of drunk people coming up with the wildest ideas and considering them brilliant. Mostly used to colloquially label ideas of sober people.",
		"Papierkrieg, der: Literally »paper war«. Describes longish correspondences with bureaucratic offices through forms and letters. Includes a slight hint of perception as unnecessary.",
		"Kabelsalat, der: Literally »cable salad«. Describes the mess that a bunch of cables somehow naturally entangle to.",
		"Blümchenkaffee, der: Literally »little flower coffee«. Weak drip coffee. Named after the little flowers painted on the bottom of porcelain cups. The aforementioned coffee is so weak that these are still visible even when the cup is full.",
	}

	rand.Seed(time.Now().Unix())

	c := twilio.NewClient(accountSid, authToken, nil)

	params := twilio.MessageParams{
		Body: words[rand.Intn(len(words))],
	}
	s, resp, err := c.Messages.Send(sender, receiver, params)
	if err != nil {
		log.Fatal("Err:", err)
	}
	log.Printf("Message Sent: %v\n", s)
	log.Printf("Response: %v\n", resp)
}
