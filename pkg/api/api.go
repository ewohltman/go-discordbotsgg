// Package api contains structs for querying the discord.bots.gg API.
package api

import (
	"encoding/json"
	"strings"
	"time"
)

// Page is a response struct from the discord.bots.gg API.
type Page struct {
	Count int    `json:"count"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Bots  []*Bot `json:"bots"`
}

func (page *Page) String() string {
	botNames := make([]string, len(page.Bots))

	for i, bot := range page.Bots {
		botNames[i] = bot.Username
	}

	return strings.Join(botNames, ", ")
}

// Bot is a response struct from the discord.bots.gg API.
type Bot struct {
	UserID           string      `json:"userId"`
	ClientID         string      `json:"clientId"`
	Username         string      `json:"username"`
	Discriminator    string      `json:"discriminator"`
	AvatarURL        string      `json:"avatarURL"`
	CoOwners         []*BotOwner `json:"coOwners"`
	Prefix           string      `json:"prefix"`
	HelpCommand      string      `json:"helpCommand"`
	LibraryName      string      `json:"libraryName"`
	Website          string      `json:"website"`
	SupportInvite    string      `json:"supportInvite"`
	BotInvite        string      `json:"botInvite"`
	ShortDescription string      `json:"shortDescription"`
	LongDescription  string      `json:"longDescription"`
	OpenSource       string      `json:"openSource"`
	ShardCount       int         `json:"shardCount"`
	GuildCount       int         `json:"guildCount"`
	Verified         bool        `json:"verified"`
	Online           bool        `json:"online"`
	InGuild          bool        `json:"inGuild"`
	Owner            *BotOwner   `json:"owner"`
	AddedDate        time.Time   `json:"addedDate"`
	Status           string      `json:"status"`
}

// String satisfies the fmt.Stringer interface and returns the bot's name.
func (bot *Bot) String() string {
	return bot.Username
}

// BotOwner is a response struct from the discord.bots.gg API.
type BotOwner struct {
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	UserID        string `json:"userId"`
}

// StatsUpdate is a request struct for the discord.bots.gg API.
type StatsUpdate struct {
	Stats
	ShardID int `json:"shardID"`
}

// StatsResponse is a response struct from the discord.bots.gg API.
type StatsResponse struct {
	Stats
}

// String satisfies the fmt.Stringer interface and returns the JSON
// representation of the *StatsResponse.
func (statsResponse *StatsResponse) String() string {
	statsResponseBytes, err := json.Marshal(statsResponse)
	if err != nil {
		return err.Error()
	}

	return string(statsResponseBytes)
}

// Stats is a struct containing metrics for a bot on discord.bots.gg.
type Stats struct {
	GuildCount int `json:"guildCount"`
	ShardCount int `json:"shardCount"`
}
