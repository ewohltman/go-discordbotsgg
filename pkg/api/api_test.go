package api

import (
	"fmt"
	"testing"
)

const (
	testBotUsername1 = "testBotUsername1"
	testBotUsername2 = "testBotUsername2"
	testGuildCount   = 100
	testShardCount   = 5
)

func TestPage_String(t *testing.T) {
	page := &Page{
		Bots: []*Bot{
			{Username: testBotUsername1},
			{Username: testBotUsername2},
		},
	}

	got := page.String()
	expected := fmt.Sprintf("%s, %s", testBotUsername1, testBotUsername2)

	if got != expected {
		t.Errorf("Unexpected result. Got: %s. Expected: %s.", got, expected)
	}
}

func TestBot_String(t *testing.T) {
	bot := &Bot{Username: testBotUsername1}

	got := bot.String()
	expected := testBotUsername1

	if got != expected {
		t.Errorf("Unexpected result. Got: %s. Expected: %s.", got, expected)
	}
}

func TestStatsResponse_String(t *testing.T) {
	statsResponse := StatsResponse{
		Stats{
			GuildCount: testGuildCount,
			ShardCount: testShardCount,
		},
	}

	got := statsResponse.String()
	expected := `{"guildCount":100,"shardCount":5}`

	if got != expected {
		t.Errorf("Unexpected result. Got: %s. Expected: %s.", got, expected)
	}
}
