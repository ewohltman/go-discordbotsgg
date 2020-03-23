package discordbotsgg

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

type QueryParameters struct {
	Q          string
	Page       int
	Limit      int
	AuthorID   int64
	AuthorName string
	Unverified bool
	Lib        string
	Sort       string
	Order      string
}

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

type Page struct {
	Count int    `json:"count"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Bots  []*Bot `json:"bots"`
}

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

type BotOwner struct {
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	UserID        string `json:"userId"`
}

type BotStats struct {
	GuildCount int `json:"guildCount"`
	ShardCount int `json:"shardCount"`
	ShardID    int `json:"shardID,omitempty"`
}
