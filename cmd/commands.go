package cmd

import "github.com/disgoorg/disgo/discord"

func RegisterCommands() []discord.ApplicationCommandCreate {
	return []discord.ApplicationCommandCreate{
		discord.SlashCommandCreate{
			Name:        "ping",
			Description: "ping bot",
			Options: nil,
		},
	}
}