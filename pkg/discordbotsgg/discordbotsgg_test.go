package discordbotsgg

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ewohltman/go-discordbotsgg/pkg/api"
	"github.com/ewohltman/go-discordbotsgg/pkg/mock"
)

const (
	testBotID = "12345"

	testParameterPage     = 1
	testParameterLimit    = 1
	testParameterAuthorID = 1

	testGuildCount = 100
	testShardCount = 5

	queryBotErrorMessage       = "Error querying bot: %s"
	queryBotsErrorMessage      = "Error querying bots: %s"
	updateBotStatsErrorMessage = "Error updating bot stats: %s"
)

func TestNewClient(t *testing.T) {
	client := NewClient(mock.NewHTTPClient(), "")

	if client == nil {
		t.Fatalf("Unexpected nil *Client")
	}
}

func ExampleNewClient() {
	httpClient := mock.NewHTTPClient() // Substitute a real *http.Client here.

	client := NewClient(httpClient, "apiToken")

	bot, err := client.QueryBot("botID", true)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("Bot: %s\n", bot)
	// Output: Bot: Test Bot 1
}

func TestClient_QueryBot(t *testing.T) {
	client := NewClient(mock.NewHTTPClient(), "")

	_, err := client.QueryBot(testBotID, false)
	if err != nil {
		t.Errorf(queryBotErrorMessage, err)
	}

	_, err = client.QueryBot(testBotID, true)
	if err != nil {
		t.Errorf(queryBotErrorMessage, err)
	}
}

func BenchmarkClient_QueryBot(b *testing.B) {
	client := NewClient(mock.NewHTTPClient(), "")
	client.queryLimiter.ReserveN(time.Now(), burstSize)

	start := time.Now()

	for i := 0; i < b.N; i++ {
		_, err := client.QueryBot(testBotID, false)
		if err != nil {
			b.Fatalf(queryBotErrorMessage, err)
		}
	}

	duration := time.Since(start).Seconds()
	actualRPS := float64(b.N) / duration
	maxRPS := float64(queryLimit) / queryTimeframe.Seconds()

	b.Logf(
		"Requests: %d, Seconds: %f, RPS: %f, Max RPS: %f",
		b.N,
		duration,
		actualRPS,
		maxRPS,
	)

	if actualRPS > maxRPS {
		b.Errorf("Failed to enforce rate limit")
	}
}

func ExampleClient_QueryBot() {
	httpClient := mock.NewHTTPClient() // Substitute a real *http.Client here.

	client := NewClient(httpClient, "apiToken")

	bot, err := client.QueryBot("botID", true)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("Bot: %s\n", bot)
	// Output: Bot: Test Bot 1
}

func TestClient_QueryBotWithContext(t *testing.T) {
	client := NewClient(mock.NewHTTPClient(), "")

	_, err := client.QueryBotWithContext(context.Background(), testBotID, false)
	if err != nil {
		t.Errorf(queryBotErrorMessage, err)
	}

	_, err = client.QueryBotWithContext(context.Background(), testBotID, true)
	if err != nil {
		t.Errorf(queryBotErrorMessage, err)
	}
}

func ExampleClient_QueryBotWithContext() {
	const contextTimeout = 30 * time.Second

	httpClient := mock.NewHTTPClient() // Substitute a real *http.Client here.

	client := NewClient(httpClient, "apiToken")

	ctx, cancelCtx := context.WithTimeout(context.Background(), contextTimeout)
	defer cancelCtx()

	bot, err := client.QueryBotWithContext(ctx, "botID", true)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("Bot: %s\n", bot)
	// Output: Bot: Test Bot 1
}

func TestClient_QueryBots(t *testing.T) {
	client := NewClient(mock.NewHTTPClient(), "")

	_, err := client.QueryBots(&api.QueryParameters{})
	if err != nil {
		t.Errorf(queryBotsErrorMessage, err)
	}

	queryParameters := &api.QueryParameters{
		Q:          "test",
		Page:       testParameterPage,
		Limit:      testParameterLimit,
		AuthorID:   testParameterAuthorID,
		AuthorName: "test",
		Unverified: true,
		Lib:        "discordgo",
		Sort:       "username",
		Order:      "DESC",
	}

	_, err = client.QueryBots(queryParameters)
	if err != nil {
		t.Errorf(queryBotsErrorMessage, err)
	}
}

func BenchmarkClient_QueryBots(b *testing.B) {
	queryParameters := &api.QueryParameters{
		Q:          "test",
		Page:       testParameterPage,
		Limit:      testParameterLimit,
		AuthorID:   testParameterAuthorID,
		AuthorName: "test",
		Unverified: true,
		Lib:        "discordgo",
		Sort:       "username",
		Order:      "DESC",
	}

	client := NewClient(mock.NewHTTPClient(), "")
	client.queryLimiter.ReserveN(time.Now(), burstSize)

	start := time.Now()

	for i := 0; i < b.N; i++ {
		_, err := client.QueryBots(queryParameters)
		if err != nil {
			b.Fatalf(queryBotsErrorMessage, err)
		}
	}

	duration := time.Since(start).Seconds()
	actualRPS := float64(b.N) / duration
	maxRPS := float64(queryLimit) / queryTimeframe.Seconds()

	b.Logf(
		"Requests: %d, Seconds: %f, RPS: %f, Max RPS: %f",
		b.N,
		duration,
		actualRPS,
		maxRPS,
	)

	if actualRPS > maxRPS {
		b.Errorf("Failed to enforce rate limit")
	}
}

func ExampleClient_QueryBots() {
	const (
		pageBotLimit = 100
		authorID     = 123456789
	)

	httpClient := mock.NewHTTPClient() // Substitute a real *http.Client here.

	client := NewClient(httpClient, "apiToken")

	queryParameters := &api.QueryParameters{
		Q:          "query",
		Page:       0,
		Limit:      pageBotLimit,
		AuthorID:   authorID,
		AuthorName: "authorName",
		Unverified: false,
		Lib:        "discordgo",
		Sort:       "guildcount",
		Order:      "DESC",
	}

	bots, err := client.QueryBots(queryParameters)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("Bots: %s\n", bots)
	// Output: Bots: Test Bot 1, Test Bot 2
}

func TestClient_QueryBotsWithContext(t *testing.T) {
	client := NewClient(mock.NewHTTPClient(), "")

	_, err := client.QueryBotsWithContext(context.Background(), &api.QueryParameters{})
	if err != nil {
		t.Errorf(queryBotsErrorMessage, err)
	}

	queryParameters := &api.QueryParameters{
		Q:          "test",
		Page:       testParameterPage,
		Limit:      testParameterLimit,
		AuthorID:   testParameterAuthorID,
		AuthorName: "test",
		Unverified: true,
		Lib:        "discordgo",
		Sort:       "username",
		Order:      "DESC",
	}

	_, err = client.QueryBotsWithContext(context.Background(), queryParameters)
	if err != nil {
		t.Errorf(queryBotsErrorMessage, err)
	}
}

func ExampleClient_QueryBotsWithContext() {
	const contextTimeout = 30 * time.Second

	httpClient := mock.NewHTTPClient() // Substitute a real *http.Client here.

	client := NewClient(httpClient, "apiToken")

	queryParameters := &api.QueryParameters{
		Q:          "test",
		Page:       testParameterPage,
		Limit:      testParameterLimit,
		AuthorID:   testParameterAuthorID,
		AuthorName: "test",
		Unverified: true,
		Lib:        "discordgo",
		Sort:       "username",
		Order:      "DESC",
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), contextTimeout)
	defer cancelCtx()

	bots, err := client.QueryBotsWithContext(ctx, queryParameters)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("Bots: %s\n", bots)
	// Output: Bots: Test Bot 1, Test Bot 2
}

func TestClient_Update(t *testing.T) {
	client := NewClient(mock.NewHTTPClient(), "")

	botStatsUpdate := &api.StatsUpdate{
		Stats: api.Stats{
			GuildCount: testGuildCount,
			ShardCount: testShardCount,
		},
		ShardID: 0,
	}

	botStatsResponse, err := client.Update(testBotID, botStatsUpdate)
	if err != nil {
		t.Errorf(updateBotStatsErrorMessage, err)
	}

	if botStatsResponse.Stats.GuildCount != testGuildCount {
		t.Errorf("Unexpected guild count stat: %d", botStatsResponse.Stats.GuildCount)
	}

	if botStatsResponse.Stats.ShardCount != testShardCount {
		t.Errorf("Unexpected shard count stat: %d", botStatsResponse.Stats.ShardCount)
	}
}

func BenchmarkClient_Update(b *testing.B) {
	client := NewClient(mock.NewHTTPClient(), "")
	client.updateLimiter.ReserveN(time.Now(), burstSize)

	start := time.Now()

	statsUpdate := &api.StatsUpdate{
		Stats: api.Stats{
			GuildCount: testGuildCount,
			ShardCount: testShardCount,
		},
		ShardID: 0,
	}

	for i := 0; i < b.N; i++ {
		_, err := client.Update(testBotID, statsUpdate)
		if err != nil {
			b.Fatalf(updateBotStatsErrorMessage, err)
		}
	}

	duration := time.Since(start).Seconds()
	actualRPS := float64(b.N) / duration
	maxRPS := float64(updateLimit) / updateTimeframe.Seconds()

	b.Logf(
		"Requests: %d, Seconds: %f, RPS: %f, Max RPS: %f",
		b.N,
		duration,
		actualRPS,
		maxRPS,
	)

	if actualRPS > maxRPS {
		b.Errorf("Failed to enforce rate limit")
	}
}

func ExampleClient_Update() {
	const (
		exampleGuildCount = 100
		exampleShardCount = 5
		exampleShardID    = 0
	)

	httpClient := mock.NewHTTPClient() // Substitute a real *http.Client here.

	client := NewClient(httpClient, "apiToken")

	botStatsUpdate := &api.StatsUpdate{
		Stats: api.Stats{
			GuildCount: exampleGuildCount,
			ShardCount: exampleShardCount,
		},
		ShardID: exampleShardID,
	}

	botStatsResponse, err := client.Update("botID", botStatsUpdate)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("%s", botStatsResponse)
	// Output: {"guildCount":100,"shardCount":5}
}

func TestClient_UpdateWithContext(t *testing.T) {
	client := NewClient(mock.NewHTTPClient(), "")

	botStatsUpdate := &api.StatsUpdate{
		Stats: api.Stats{
			GuildCount: testGuildCount,
			ShardCount: testShardCount,
		},
		ShardID: 0,
	}

	botStatsResponse, err := client.UpdateWithContext(context.Background(), testBotID, botStatsUpdate)
	if err != nil {
		t.Errorf(updateBotStatsErrorMessage, err)
	}

	if botStatsResponse.Stats.GuildCount != testGuildCount {
		t.Errorf("Unexpected guild count stat: %d", botStatsResponse.Stats.GuildCount)
	}

	if botStatsResponse.Stats.ShardCount != testShardCount {
		t.Errorf("Unexpected shard count stat: %d", botStatsResponse.Stats.ShardCount)
	}
}

func ExampleClient_UpdateWithContext() {
	const (
		contextTimeout = 30 * time.Second

		exampleGuildCount = 100
		exampleShardCount = 5
		exampleShardID    = 0
	)

	httpClient := mock.NewHTTPClient() // Substitute a real *http.Client here.

	client := NewClient(httpClient, "apiToken")

	botStatsUpdate := &api.StatsUpdate{
		Stats: api.Stats{
			GuildCount: exampleGuildCount,
			ShardCount: exampleShardCount,
		},
		ShardID: exampleShardID,
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), contextTimeout)
	defer cancelCtx()

	botStatsResponse, err := client.UpdateWithContext(ctx, "botID", botStatsUpdate)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("%s", botStatsResponse)
	// Output: {"guildCount":100,"shardCount":5}
}
