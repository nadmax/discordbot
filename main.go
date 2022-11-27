package main

import (
	"discordbot/events"

	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New(os.Getenv("TOKEN"))
	
	if err != nil {
		log.Fatalf("error while creating app: %v", err)
	}

	dg.AddHandler(events.OnLogin)
	
	if err := dg.Open(); err != nil {
		log.Fatalf("Cannot open session: %v", err)
	}

	log.Println("Adding commands...")
	
}