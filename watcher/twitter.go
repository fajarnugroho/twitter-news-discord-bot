package watcher

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	. "github.com/fajarnugroho/twitter-news-discord-bot/logger"
	"github.com/fajarnugroho/twitter-news-discord-bot/util"
)

const (
	TWITTER_STATUS_URL = "https://twitter.com/%s/status/%s"
	RFC1123            = "Mon, 02 Jan 2006 15:04:05 MST"
)

type TwitterNewsAggregator struct {
	HTTPClient *http.Client
	Demux      twitter.SwitchDemux
	Client     *twitter.Client
	Ds         *discordgo.Session
	Users      map[string]bool
	AccountIDs []string
}

func Init(dg *discordgo.Session, userIDs []string) *TwitterNewsAggregator {
	config := oauth1.NewConfig(os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("TWITTER_ACCESS_TOKEN"),
		os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter Client
	tna := &TwitterNewsAggregator{}
	tna.Client = twitter.NewClient(httpClient)
	tna.Demux = twitter.NewSwitchDemux()
	tna.Ds = dg
	tna.AccountIDs = tna.GetAccountIDs(userIDs)
	tna.Users = map[string]bool{}
	for _, accountID := range tna.AccountIDs {
		tna.Users[accountID] = true
	}
	tna.Demux.Tweet = func(tweet *twitter.Tweet) {
		Log.Debug("new tweet ", tweet.Text)
		Log.Debugf("tweeted by %s", tweet.User.IDStr)
		userID := tweet.User.IDStr
		_, ok := tna.Users[userID]
		if !ok {
			return

		}
		tna.PostNews(tweet)
	}
	return tna
}

func (t *TwitterNewsAggregator) GetAccountIDs(usernames []string) []string {
	accountIDs := []string{}
	for _, username := range usernames {
		user, _, err := t.Client.Users.Show(&twitter.UserShowParams{
			ScreenName: username,
		})
		if err != nil {
			Log.Error("error getting user id for username %s: %v", username, err)
		}
		accountIDs = append(accountIDs, strconv.FormatInt(user.ID, 10))
	}
	return accountIDs
}

func (t *TwitterNewsAggregator) PostNews(tweet *twitter.Tweet) {
	statusUrl := fmt.Sprintf(TWITTER_STATUS_URL, tweet.User.ScreenName, tweet.IDStr)
	tct, err := tweet.CreatedAtTime()
	if err != nil {
		Log.Error("error parsing tweet created_at time", err)
	}
	tctLocal, err := util.FormatTime(tct, os.Getenv("TIMEZONE"))
	if err != nil {
		Log.Error("error get local time ", err)
	}
	message := fmt.Sprintf(
		"@%s tweeted this on %s : %s",
		tweet.User.ScreenName,
		tctLocal,
		statusUrl)
	t.Ds.ChannelMessageSend(os.Getenv("DISCORD_CHANNEL_ID"), util.EscapeString(message))
}

func (t *TwitterNewsAggregator) SetupStreamer() (*twitter.Stream, error) {
	// FILTER
	filterParams := &twitter.StreamFilterParams{
		Follow:        t.AccountIDs,
		StallWarnings: twitter.Bool(true),
	}
	return t.Client.Streams.Filter(filterParams)
}
