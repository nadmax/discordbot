package cmd

import (
	"discordbot/events"

	"github.com/bwmarrin/discordgo"
)

func HandleCommands() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i*discordgo.InteractionCreate) {
		"ping": events.Ping,
	}
}