package discordbotsgg

import (
	"context"
	"testing"
	"time"

	"github.com/ewohltman/go-discordbotsgg/pkg/mock"
)

const (
	testBotID = "12345"

	testParameterPage     = 1
	testParameterLimit    = 1
	testParameterAuthorID = 1

	benchmarkRequests = 10

	queryBotErrorMessage  = "Error querying bot: %s"
	queryBotsErrorMessage = "Error querying bots: %s"
)

func TestNewClient(t *testing.T) {
	client := NewClient(mock.NewHTTPClient(), "")

	if client == nil {
		t.Fatalf("Unexpected nil *Client")
	}
}

func TestClient_QueryBot(t *testing.T) {
	var err error

	client := NewClient(mock.NewHTTPClient(), "")

	_, err = client.QueryBot(testBotID, false)
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

	for i := 0; i < benchmarkRequests; i++ {
		_, err := client.QueryBot(testBotID, false)
		if err != nil {
			b.Fatalf(queryBotErrorMessage, err)
		}
	}

	duration := time.Since(start).Seconds()
	actualRPS := float64(benchmarkRequests) / duration
	maxRPS := float64(queryLimit) / queryTimeframe.Seconds()

	b.Logf(
		"Requests: %d, Seconds: %f, RPS: %f, Max RPS: %f",
		benchmarkRequests,
		duration,
		actualRPS,
		maxRPS,
	)

	if actualRPS > maxRPS {
		b.Errorf("Failed to enforce rate limit")
	}
}

func TestClient_QueryBotWithContext(t *testing.T) {
	var err error

	client := NewClient(mock.NewHTTPClient(), "")

	_, err = client.QueryBotWithContext(context.Background(), testBotID, false)
	if err != nil {
		t.Errorf(queryBotErrorMessage, err)
	}

	_, err = client.QueryBotWithContext(context.Background(), testBotID, true)
	if err != nil {
		t.Errorf(queryBotErrorMessage, err)
	}
}

func TestClient_QueryBots(t *testing.T) {
	var err error

	_, err = NewClient(mock.NewHTTPClient(), "").QueryBots(&QueryParameters{})
	if err != nil {
		t.Errorf(queryBotsErrorMessage, err)
	}

	queryParameters := &QueryParameters{
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

	_, err = NewClient(mock.NewHTTPClient(), "").QueryBots(queryParameters)
	if err != nil {
		t.Errorf(queryBotsErrorMessage, err)
	}
}

func BenchmarkClient_QueryBots(b *testing.B) {
	queryParameters := &QueryParameters{
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

	for i := 0; i < benchmarkRequests; i++ {
		_, err := client.QueryBots(queryParameters)
		if err != nil {
			b.Fatalf(queryBotsErrorMessage, err)
		}
	}

	duration := time.Since(start).Seconds()
	actualRPS := float64(benchmarkRequests) / duration
	maxRPS := float64(queryLimit) / queryTimeframe.Seconds()

	b.Logf(
		"Requests: %d, Seconds: %f, RPS: %f, Max RPS: %f",
		benchmarkRequests,
		duration,
		actualRPS,
		maxRPS,
	)

	if actualRPS > maxRPS {
		b.Errorf("Failed to enforce rate limit")
	}
}

func TestClient_QueryBotsWithContext(t *testing.T) {
	var err error

	_, err = NewClient(mock.NewHTTPClient(), "").QueryBotsWithContext(context.Background(), &QueryParameters{})
	if err != nil {
		t.Errorf(queryBotsErrorMessage, err)
	}

	queryParameters := &QueryParameters{
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

	_, err = NewClient(mock.NewHTTPClient(), "").QueryBotsWithContext(context.Background(), queryParameters)
	if err != nil {
		t.Errorf(queryBotsErrorMessage, err)
	}
}
