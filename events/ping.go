package events

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func Ping(event *events.ApplicationCommandInteractionCreate) {
	data := event.SlashCommandInteractionData()

	if data.CommandName() == "ping" {
		if err := event.CreateMessage(discord.NewMessageCreateBuilder().
			SetContent("pong").
			Build(),
		); err != nil {
			event.Client().Logger().Error("error on sending response: ", err)
		}
	}
}