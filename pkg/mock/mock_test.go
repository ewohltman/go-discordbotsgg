package mock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ewohltman/go-discordbotsgg/pkg/api"
)

const (
	testGuildCount = 100
	testShardCount = 5
)

func TestNewHTTPClient(t *testing.T) {
	mockClient := NewHTTPClient()

	if mockClient == nil {
		t.Error("Unexpected nil mock *http.Client")
	}

	err := doTestRequests(mockClient)
	if err != nil {
		t.Errorf("Unexpected error performing test requests: %s", err)
	}
}

func TestRoundTripperFunc_RoundTrip(t *testing.T) {
	transport := NewTransport()

	if transport == nil {
		t.Error("Unexpected nil mock http.RoundTripper")
	}

	mockClient := &http.Client{Transport: transport}

	err := doTestRequests(mockClient)
	if err != nil {
		t.Errorf("Unexpected error performing test requests: %s", err)
	}
}

func doTestRequests(client *http.Client) error {
	err := doTestRequest(client, http.MethodGet, "http://localhost/badEndpoint", nil)
	if err != nil {
		return err
	}

	err = doTestRequest(client, http.MethodGet, api.BotEndpoint("botID", true), nil)
	if err != nil {
		return err
	}

	err = doTestRequest(client, http.MethodGet, api.BotEndpoint("botID", true), nil)
	if err != nil {
		return err
	}

	err = doTestRequest(client, http.MethodGet, api.BotsEndpoint(nil), nil)
	if err != nil {
		return err
	}

	statsUpdate := &api.StatsUpdate{
		Stats: api.Stats{
			GuildCount: testGuildCount,
			ShardCount: testShardCount,
		},
		ShardID: 0,
	}

	statsUpdateBytes, err := json.Marshal(statsUpdate)
	if err != nil {
		return err
	}

	err = doTestRequest(client, http.MethodPost, "http://localhost/badEndpoint", bytes.NewReader(statsUpdateBytes))
	if err != nil {
		return err
	}

	err = doTestRequest(client, http.MethodPost, api.StatsEndpoint("botID"), bytes.NewReader(statsUpdateBytes))
	if err != nil {
		return err
	}

	err = doTestRequest(client, http.MethodPut, "http://localhost/badMethod", nil)
	if err != nil {
		return err
	}

	return nil
}

func doTestRequest(client *http.Client, method, url string, body io.Reader) (err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			if err != nil {
				err = fmt.Errorf("%s: %w", closeErr, err)
			} else {
				err = closeErr
			}
		}
	}()

	_, err = io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
