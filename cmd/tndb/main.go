package main

import (
	"os"
	"os/signal"
	"syscall"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/bwmarrin/discordgo"
	. "github.com/fajarnugroho/twitter-news-discord-bot/bot"
	"github.com/sirupsen/logrus"
)

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	log := logrus.New()
	log.SetFormatter(&nested.Formatter{
		HideKeys:      true,
		NoColors:      true,
		CallerFirst:   true,
		ShowFullLevel: true,
	})
	log.SetReportCaller(true)
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(MessageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
