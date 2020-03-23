package discordbotsgg

import (
	"context"
	"net/http"
	"testing"
)

const (
	testBotID = "241930933962407936"

	queryBotErrorMessage  = "Error querying bot: %s"
	queryBotsErrorMessage = "Error querying bots: %s"
)

func TestNewClient(t *testing.T) {
	client := NewClient(&http.Client{}, "")

	if client == nil {
		t.Fatalf("Unexpected nil *Client")
	}
}

func TestClient_QueryBot(t *testing.T) {
	var err error

	client := NewClient(&http.Client{}, "")

	_, err = client.QueryBot(testBotID, false)
	if err != nil {
		t.Errorf(queryBotErrorMessage, err)
	}

	_, err = client.QueryBot(testBotID, true)
	if err != nil {
		t.Errorf(queryBotErrorMessage, err)
	}
}

func TestClient_QueryBotWithContext(t *testing.T) {
	var err error

	client := NewClient(&http.Client{}, "")

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
	_, err := NewClient(&http.Client{}, "").QueryBots(&QueryParameters{})
	if err != nil {
		t.Errorf(queryBotsErrorMessage, err)
	}
}

func TestClient_QueryBotsWithContext(t *testing.T) {
	_, err := NewClient(&http.Client{}, "").QueryBotsWithContext(context.Background(), &QueryParameters{})
	if err != nil {
		t.Errorf(queryBotsErrorMessage, err)
	}
}
