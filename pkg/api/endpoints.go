package api

import "fmt"

const (
	baseURL       = "https://discord.bots.gg"
	botEndpoint   = baseURL + "/api/v1/bots/%s?sanitize=%t"
	botsEndpoint  = baseURL + "/api/v1/bots"
	statsEndpoint = baseURL + "/api/v1/bots/%s/stats"
)

// BotEndpoint returns an API URL string for querying the given botID.
func BotEndpoint(botID string, sanitize bool) string {
	return fmt.Sprintf(botEndpoint, botID, sanitize)
}

// BotsEndpoint returns an API URL string for querying bots.
func BotsEndpoint(queryParameters fmt.Stringer) string {
	if queryParameters == nil {
		return botsEndpoint
	}

	return fmt.Sprintf("%s?%s", botsEndpoint, queryParameters)
}

// StatsEndpoint returns an API URL string for updating stats for the given
// botID.
func StatsEndpoint(botID string) string {
	return fmt.Sprintf(statsEndpoint, botID)
}
