# Configuration properties

Check out `.env.example` for example configuration

## Global configuration properties

Property                    | Type      | Default                       | Description
--------------------------- | --------- | ----------------------------- | -------------
TWITTER_ACCESS_TOKEN        | string    | twitter_access_token          | Twitter developer access token
TWITTER_ACCESS_TOKEN_SECRET |string     | twitter_access_token_secret   | Twitter developer access token secret
TWITTER_CONSUMER_KEY        |string     | twitter_consumer_key          | twitter developer consumer key
TWITTER_CONSUMER_SECRET     |string     | twitter_consumer_secret       | Twitter Developer consumer secret
DISCORD_BOT_TOKEN           |string     | discord_bot_token             | Discord bot token
DISCORD_CHANNEL_ID          |number     | 1234567890                    | Discord channel ID. Can obtained from discord application with developer mode enabled
ACCOUNTS                    |string     | TirtoID,asumsico              | Twitter accounts to follow
ENV                         |string     | production                    | Deployment environment (production,development)
TIMEZONE                    |string     | Asia/Jakarta                  | Timezone used in discord channel time

