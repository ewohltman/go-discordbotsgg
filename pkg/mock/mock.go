// Package mock provides a mock implementation of the discord.bots.gg API for
// testing.
package mock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ewohltman/go-discordbotsgg/pkg/api"
)

type roundTripperFunc func(req *http.Request) (*http.Response, error)

func (rt roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return rt(req)
}

// NewHTTPClient returns a new *http.Client using a mock http.RoundTripper
// as its Transport.
func NewHTTPClient() *http.Client {
	return &http.Client{Transport: NewTransport()}
}

// NewTransport returns a new mock http.RoundTripper to be used as an
// *http.Client Transport.
func NewTransport() http.RoundTripper {
	return roundTripperFunc(
		func(req *http.Request) (*http.Response, error) {
			resp := &http.Response{
				StatusCode: http.StatusOK,
				Status:     http.StatusText(http.StatusOK),
				Header:     make(http.Header),
				Request:    req,
			}

			switch req.Method {
			case http.MethodGet:
				err := handleGetRequest(req, resp)
				if err != nil {
					return nil, err
				}

				return resp, nil
			case http.MethodPost:
				err := handlePostRequest(req, resp)
				if err != nil {
					return nil, err
				}

				return resp, nil
			}

			badRequestResponse(resp)

			return resp, nil
		},
	)
}

func handleGetRequest(req *http.Request, resp *http.Response) error {
	if strings.Contains(req.URL.Path, "/bots/") {
		return botResponse(req, resp)
	}

	if strings.Contains(req.URL.Path, "/bots") {
		return botsResponse(req, resp)
	}

	badRequestResponse(resp)

	return nil
}

func handlePostRequest(req *http.Request, resp *http.Response) error {
	if strings.Contains(req.URL.Path, "/stats") {
		return updateBotResponse(req, resp)
	}

	badRequestResponse(resp)

	return nil
}

func botResponse(req *http.Request, resp *http.Response) error {
	_, err := readRequestBody(req)
	if err != nil {
		return err
	}

	respBody := []byte(botResponseString)

	resp.ContentLength = int64(len(respBody))
	resp.Body = ioutil.NopCloser(bytes.NewReader(respBody))

	return nil
}

func botsResponse(req *http.Request, resp *http.Response) error {
	_, err := readRequestBody(req)
	if err != nil {
		return err
	}

	respBody := []byte(botsResponseString)

	resp.ContentLength = int64(len(respBody))
	resp.Body = ioutil.NopCloser(bytes.NewReader(respBody))

	return nil
}

func updateBotResponse(req *http.Request, resp *http.Response) error {
	reqBody, err := readRequestBody(req)
	if err != nil {
		return err
	}

	botStatsUpdate := &api.StatsUpdate{}

	err = json.Unmarshal(reqBody, botStatsUpdate)
	if err != nil {
		return err
	}

	botStatsResponse := &api.StatsResponse{
		Stats: botStatsUpdate.Stats,
	}

	respBody, err := json.Marshal(botStatsResponse)
	if err != nil {
		return err
	}

	resp.ContentLength = int64(len(respBody))
	resp.Body = ioutil.NopCloser(bytes.NewReader(respBody))

	return nil
}

func readRequestBody(req *http.Request) (reqBody []byte, err error) {
	if req.Body == nil {
		return nil, nil
	}

	defer func() {
		closeErr := req.Body.Close()
		if closeErr != nil {
			if err != nil {
				err = fmt.Errorf("%s: %w", closeErr, err)
				return
			}

			err = closeErr
		}
	}()

	return ioutil.ReadAll(req.Body)
}

func badRequestResponse(resp *http.Response) {
	resp.StatusCode = http.StatusBadRequest
	resp.Status = http.StatusText(http.StatusBadRequest)
	resp.Body = ioutil.NopCloser(bytes.NewReader([]byte{}))
}
