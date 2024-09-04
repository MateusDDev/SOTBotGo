package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"sot/src/config"
	"sot/src/events"

	"github.com/bwmarrin/discordgo"
)

func main() {
	Token := config.GetToken()

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session, ", err)
		return
	}

	dg.AddHandler(events.MessageCreate)

	dg.Identify.Intents = discordgo.IntentGuildMessages | discordgo.IntentGuildMembers

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
