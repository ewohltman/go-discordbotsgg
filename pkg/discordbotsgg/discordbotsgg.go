// Package discordbotsgg provides a client implementation for interacting with
// the discord.bots.gg API.
package discordbotsgg

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

const (
	apiURL = "https://discord.bots.gg/api/v1/bots"

	queryLimit     = 10
	queryBurstSize = queryLimit
	queryTimeframe = 5 * time.Second

	updateLimit     = 20
	updateBurstSize = updateLimit
	updateTimeframe = time.Second
)

// HTTPClient is an interface to abstract HTTP client implementations.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client is a discord.bots.gg client.
type Client struct {
	HTTPClient    HTTPClient
	APIToken      string
	queryLimiter  *rate.Limiter
	updateLimiter *rate.Limiter
}

// NewClient returns a new *Client with configured rate limiters.
func NewClient(httpClient HTTPClient, apiToken string) *Client {
	return &Client{
		HTTPClient:    httpClient,
		APIToken:      apiToken,
		queryLimiter:  rate.NewLimiter(rate.Every(queryTimeframe/queryLimit), queryBurstSize),
		updateLimiter: rate.NewLimiter(rate.Every(updateTimeframe/updateLimit), updateBurstSize),
	}
}

// QueryBot returns information about the given botID.
func (client *Client) QueryBot(botID string, sanitize bool) (*Bot, error) {
	return client.QueryBotWithContext(context.Background(), botID, sanitize)
}

// QueryBotWithContext returns information about the given botID using the
// provided context.Context.
func (client *Client) QueryBotWithContext(ctx context.Context, botID string, sanitize bool) (*Bot, error) {
	return client.queryBot(ctx, botID, sanitize)
}

func (client *Client) queryBot(ctx context.Context, botID string, sanitize bool) (*Bot, error) {
	queryURL := fmt.Sprintf("%s/%s?sanitize=%t", apiURL, botID, sanitize)

	bot := &Bot{}

	err := client.queryLimiter.Wait(ctx)
	if err != nil {
		return nil, err
	}

	err = client.doRequest(ctx, queryURL, bot)
	if err != nil {
		return nil, err
	}

	return bot, nil
}

// QueryBots returns results using the provided parameters.
func (client *Client) QueryBots(queryParameters fmt.Stringer) ([]*Bot, error) {
	return client.QueryBotsWithContext(context.Background(), queryParameters)
}

// QueryBotsWithContext returns results using the provided parameters and context.Context.
func (client *Client) QueryBotsWithContext(ctx context.Context, queryParameters fmt.Stringer) ([]*Bot, error) {
	return client.queryBots(ctx, queryParameters)
}

func (client *Client) queryBots(ctx context.Context, queryParameters fmt.Stringer) ([]*Bot, error) {
	parametersValues := queryParameters.String()

	var queryURL string

	if parametersValues == "" {
		queryURL = apiURL
	} else {
		queryURL = fmt.Sprintf("%s?%s", apiURL, parametersValues)
	}

	page := &Page{}

	err := client.queryLimiter.Wait(ctx)
	if err != nil {
		return nil, err
	}

	err = client.doRequest(ctx, queryURL, page)
	if err != nil {
		return nil, err
	}

	return page.Bots, err
}

// Update updates the given botID with the provided botStats.
func (client *Client) Update(botID string, botStats *BotStats) error {
	err := client.updateLimiter.Wait(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) doRequest(ctx context.Context, queryURL string, object interface{}) (err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, queryURL, nil)
	if err != nil {
		return err
	}

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			if err != nil {
				err = fmt.Errorf("%s: %w", closeErr, err)
				return
			}

			err = closeErr
		}
	}()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(respBody, object)
}
