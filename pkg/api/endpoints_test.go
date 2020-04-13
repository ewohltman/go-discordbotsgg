package api

import (
	"fmt"
	"testing"
)

const (
	testBotID = "testBotID"
	testQuery = "testQuery"
)

func TestBotEndpoint(t *testing.T) {
	got := BotEndpoint(testBotID, false)
	expected := fmt.Sprintf(botEndpoint, testBotID, false)

	if got != expected {
		t.Errorf("Unexpected result. Got: %s. Expected: %s", got, expected)
	}

	got = BotEndpoint(testBotID, true)
	expected = fmt.Sprintf(botEndpoint, testBotID, true)

	if got != expected {
		t.Errorf("Unexpected result. Got: %s. Expected: %s", got, expected)
	}
}

func TestBotsEndpoint(t *testing.T) {
	got := BotsEndpoint(nil)
	expected := botsEndpoint

	if got != expected {
		t.Errorf("Unexpected result. Got: %s. Expected: %s", got, expected)
	}

	queryParameters := &QueryParameters{
		Q: testQuery,
	}

	got = BotsEndpoint(queryParameters)
	expected = fmt.Sprintf("%s?%s", botsEndpoint, queryParameters)

	if got != expected {
		t.Errorf("Unexpected result. Got: %s. Expected: %s", got, expected)
	}
}

func TestStatsEndpoint(t *testing.T) {
	got := StatsEndpoint(testBotID)
	expected := fmt.Sprintf(statsEndpoint, testBotID)

	if got != expected {
		t.Errorf("Unexpected result. Got: %s. Expected: %s", got, expected)
	}
}
