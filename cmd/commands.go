package cmd

import "github.com/bwmarrin/discordgo"


func SetCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand {
		{
			Name: "ping",
			Description: "Send a request to bot",
		},
	}
}