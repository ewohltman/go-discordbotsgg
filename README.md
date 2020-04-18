# go-discordbotsgg
Go Client for https://discord.bots.gg

The `go-discordbotsgg` library provides a client with built-in rate limiting
for sending requests to the `discord.bots.gg` API.

The API requests and rate limits follow the specs defined in the API
documentation. You must be logged in to view the documentation:
https://discord.bots.gg/docs

## Examples

### Query a specific bot
Note: An API token is not required to query the API. If you do not have an API
token, pass an empty string as the second parameter to `NewClient`.

```go
httpClient := &http.Client{}

client := discordbotsgg.NewClient(httpClient, "apiToken")
defer client.Close()

sanitize := true

bot, _ = client.QueryBotWithContext(context.TODO(), "botID", sanitize)

fmt.Printf("Bot: %+v\n", bot)
```

### Query bots with search parameters
Note: An API token is not required to query the API. If you do not have an API
token, pass an empty string as the second parameter to `NewClient`.

```go
httpClient := &http.Client{}

client := discordbotsgg.NewClient(httpClient, "apiToken")
defer client.Close()

queryParameters := &api.QueryParameters{
    Q:          "botNameOrDescription",
}

bots, _ := client.QueryBotsWithContext(context.TODO(), queryParameters)

fmt.Printf("Bots: %+v\n", bot)
```

### Update a bot's stats
Note: An API token is required to send updates to the API.

```go
httpClient := &http.Client{}

client := discordbotsgg.NewClient(httpClient, "apiToken")
defer client.Close()

botStatsUpdate := &api.StatsUpdate{
    Stats: &api.Stats{
        GuildCount: totalGuildCount,
        ShardCount: totalShardCount,
    },
}

botStatsResponse, _ := client.UpdateWithContext(context.TODO(), "botID", botStatsUpdate)

fmt.Printf("Update bot response: %s\n", botStatsResponse)
```
