package discordbotsgg

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

// QueryParameters are parameters that can be set for querying bots.
type QueryParameters struct {
	Q          string // Searches for bots that contain the query in their username or short description.
	Page       int    // The page to look at. Default is 0.
	Limit      int    // The number of results to retrieve. Must be between 1 and 100. Default is 50.
	AuthorID   int64  // Retrieves bots by the specified author/co-owner's ID.
	AuthorName string // Retrieves bots by the specified author/co-owner’s username and discriminator. Must be url encoded. (e.g. User%231234)
	Unverified bool   // Retrieves unverified bots. Requires authentication. Default is false.
	Lib        string // Retrieves bots written in a specific library.
	Sort       string // Sorts the results by any of the following keys: username, id, guildcount, library, author.
	Order      string // Sorts the results in ASC or DESC order.
}

// String is the URL value-encoded representation of a *QueryParameters.
func (queryParameters *QueryParameters) String() string {
	values := make(url.Values)

	if queryParameters.Q != "" {
		values["q"] = []string{queryParameters.Q}
	}

	if queryParameters.Page > 0 {
		values["page"] = []string{strconv.Itoa(queryParameters.Page)}
	}

	if queryParameters.Limit > 0 {
		values["limit"] = []string{strconv.Itoa(queryParameters.Limit)}
	}

	if queryParameters.AuthorID > 0 {
		values["authorId"] = []string{strconv.FormatInt(queryParameters.AuthorID, 10)}
	}

	if queryParameters.AuthorName != "" {
		values["authorName"] = []string{queryParameters.AuthorName}
	}

	if queryParameters.Unverified {
		values["unverified"] = []string{"true"}
	}

	if queryParameters.Lib != "" {
		values["lib"] = []string{queryParameters.Lib}
	}

	if queryParameters.Sort != "" {
		sort := strings.ToLower(queryParameters.Sort)

		switch sort {
		case "username":
			fallthrough
		case "id":
			fallthrough
		case "guildcount":
			fallthrough
		case "library":
			fallthrough
		case "author":
			values["sort"] = []string{sort}
		}
	}

	if queryParameters.Order != "" {
		order := strings.ToUpper(queryParameters.Order)

		switch order {
		case "DESC":
			fallthrough
		case "ASC":
			values["order"] = []string{order}
		}
	}

	return values.Encode()
}

// Page is a response struct from the discord.bots.gg API.
type Page struct {
	Count int    `json:"count"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Bots  []*Bot `json:"bots"`
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

// BotOwner is a response struct from the discord.bots.gg API.
type BotOwner struct {
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	UserID        string `json:"userId"`
}

// BotStats is both a request and response struct for the discord.bots.gg API.
type BotStats struct {
	GuildCount int `json:"guildCount"`
	ShardCount int `json:"shardCount"`
	ShardID    int `json:"shardID,omitempty"`
}
