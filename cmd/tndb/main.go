package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	. "github.com/fajarnugroho/twitter-news-discord-bot/logger"
	"github.com/fajarnugroho/twitter-news-discord-bot/watcher"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		err := godotenv.Load()
		if err != nil {
			Log.Fatal("Error loading .env file")
		}
	}
	env := strings.ToLower(os.Getenv("ENV"))
	if env == "dev" || env == "development" {
		Log.SetLevel(logrus.DebugLevel)
	}

}

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		Log.Fatal("error creating Discord session,", err)
		return
	}

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	err = dg.Open()
	if err != nil {
		Log.Fatal("error opening connection,", err)
		return
	}
	accounts := strings.Split(os.Getenv("ACCOUNTS"), ",")
	tna := watcher.Init(dg, accounts)
	stream, err := tna.SetupStreamer()
	if err != nil {
		Log.Fatal("error setup streamer", err)
	}
	go tna.Demux.HandleChan(stream.Messages)

	Log.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	stream.Stop()
	dg.Close()
}
