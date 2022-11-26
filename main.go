package main

import (
	"discordbot/cmd"
	"discordbot/events"
	
	"strconv"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
)

func main() {
	log.SetLevel(log.LevelInfo)

	id, _ := strconv.Atoi(os.Getenv("GUILD_ID"))
	guildID := snowflake.ID(id)

	client, err := disgo.New(
		os.Getenv("TOKEN"),
		bot.WithDefaultGateway(),
		bot.WithEventListenerFunc(events.Ping),
	)
	if err != nil {
		log.Fatal("error while building disgo instance: ", err)
		return
	}

	defer client.Close(context.TODO())

	if _, err = client.Rest().SetGuildCommands(client.ApplicationID(), guildID, cmd.RegisterCommands()); err != nil {
		log.Fatal("error while registering commands: ", err)
	}

	if err = client.OpenGateway(context.TODO()); err != nil {
		log.Fatal("error while connecting to gateway: ", err)
	}

	log.Infof("Application is now running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}