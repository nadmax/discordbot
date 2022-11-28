package main

import (
	"discordbot/cmd"
	"syscall"

	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func main() {

	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	
	if err != nil {
		log.Fatalf("error while creating app: %v", err)
	}

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := cmd.HandleCommands()[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	
	if err := dg.Open(); err != nil {
		log.Fatalf("cannot open session: %v", err)
	}
	
	defer dg.Close()
	
	log.Println("Adding commands...")
	commands := cmd.SetCommands()

	for _, v := range commands {
		if _, err := dg.ApplicationCommandCreate(dg.State.User.ID, os.Getenv("GUILD_ID"), v); err != nil {
			log.Panicf("cannot create '%v' command: %v", v.Name, err)
		}
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	log.Println("Application is running. Press Ctrl+C to exit.")
	<-s

	log.Println("Application has been successfully stopped.")
}