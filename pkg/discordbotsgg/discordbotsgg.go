// Package discordbotsgg provides a client implementation for interacting with
// the discord.bots.gg API.
package discordbotsgg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/time/rate"

	"github.com/ewohltman/go-discordbotsgg/pkg/api"
)

const (
	burstSize = 1

	queryLimit     = 10
	queryTimeframe = 5 * time.Second

	updateLimit     = 20
	updateTimeframe = time.Second
)

// HTTPClient is an interface to abstract HTTP client implementations.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client is a discord.bots.gg client.
type Client struct {
	HTTPClient    HTTPClient
	BotToken      string
	queryLimiter  *rate.Limiter
	updateLimiter *rate.Limiter
}

// NewClient returns a new *Client with configured rate limiters.
func NewClient(httpClient HTTPClient, botToken string) *Client {
	return &Client{
		HTTPClient:    httpClient,
		BotToken:      botToken,
		queryLimiter:  rate.NewLimiter(rate.Every(queryTimeframe/queryLimit), burstSize),
		updateLimiter: rate.NewLimiter(rate.Every(updateTimeframe/updateLimit), burstSize),
	}
}

// QueryBot returns information about the given botID.
func (client *Client) QueryBot(botID string, sanitize bool) (*api.Bot, error) {
	return client.QueryBotWithContext(context.TODO(), botID, sanitize)
}

// QueryBotWithContext returns information about the given botID using the
// provided context.
func (client *Client) QueryBotWithContext(ctx context.Context, botID string, sanitize bool) (*api.Bot, error) {
	bot := &api.Bot{}

	err := client.queryLimiter.Wait(ctx)
	if err != nil {
		return nil, err
	}

	err = client.doGetRequest(ctx, api.BotEndpoint(botID, sanitize), bot)
	if err != nil {
		return nil, err
	}

	return bot, nil
}

// QueryBots returns results using the provided parameters.
func (client *Client) QueryBots(queryParameters fmt.Stringer) (*api.Page, error) {
	return client.QueryBotsWithContext(context.TODO(), queryParameters)
}

// QueryBotsWithContext returns results using the provided parameters and context.
func (client *Client) QueryBotsWithContext(ctx context.Context, queryParameters fmt.Stringer) (*api.Page, error) {
	page := &api.Page{}

	err := client.queryLimiter.Wait(ctx)
	if err != nil {
		return nil, err
	}

	err = client.doGetRequest(ctx, api.BotsEndpoint(queryParameters), page)
	if err != nil {
		return nil, err
	}

	return page, err
}

// Update updates the given botID with the provided botStats.
func (client *Client) Update(botID string, statsUpdate *api.StatsUpdate) (*api.StatsResponse, error) {
	return client.UpdateWithContext(context.TODO(), botID, statsUpdate)
}

// UpdateWithContext updates the given botID with the provided botStats and context.
func (client *Client) UpdateWithContext(ctx context.Context, botID string, statsUpdate *api.StatsUpdate) (*api.StatsResponse, error) {
	err := client.updateLimiter.Wait(ctx)
	if err != nil {
		return nil, err
	}

	statsResponse := &api.StatsResponse{}

	err = client.doPostRequest(ctx, api.StatsEndpoint(botID), statsUpdate, statsResponse)
	if err != nil {
		return nil, err
	}

	return statsResponse, nil
}

func (client *Client) doGetRequest(ctx context.Context, queryURL string, responseObject interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, queryURL, nil)
	if err != nil {
		return err
	}

	return client.doRequest(req, responseObject)
}

func (client *Client) doPostRequest(ctx context.Context, queryURL string, requestObject, responseObject interface{}) error {
	requestObjectBytes, err := json.Marshal(requestObject)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, queryURL, bytes.NewReader(requestObjectBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", client.BotToken)
	req.Header.Set("Content-Type", "application/json")
	req.ContentLength = int64(len(requestObjectBytes))

	return client.doRequest(req, responseObject)
}

func (client *Client) doRequest(req *http.Request, responseObject interface{}) (err error) {
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

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	return json.Unmarshal(respBody, responseObject)
}
