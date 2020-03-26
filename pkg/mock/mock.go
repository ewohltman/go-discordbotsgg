package mock

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type roundTripperFunc func(req *http.Request) (*http.Response, error)

func (rt roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return rt(req)
}

func NewHTTPClient() *http.Client {
	return &http.Client{Transport: NewTransport()}
}

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
				return handleGetRequest(req, resp)
			case http.MethodPost:
				return handlePostRequest(req, resp)
			}

			return badRequestResponse(resp)
		},
	)
}

func handleGetRequest(req *http.Request, resp *http.Response) (*http.Response, error) {
	if req.URL.Path == "/bots" {
		return botResponse(req, resp)
	}

	if strings.Contains(req.URL.Path, "/bots") {
		return botsResponse(req, resp)
	}

	return badRequestResponse(resp)
}

func handlePostRequest(req *http.Request, resp *http.Response) (*http.Response, error) {
	if strings.Contains(req.URL.Path, "/bots") {
		return updateBotResponse(req, resp)
	}

	return badRequestResponse(resp)
}

func botResponse(req *http.Request, resp *http.Response) (*http.Response, error) {
	_, err := readRequestBody(req)
	if err != nil {
		return nil, err
	}

	respBody := []byte(botResponseString)

	resp.ContentLength = int64(len(respBody))
	resp.Body = ioutil.NopCloser(bytes.NewReader(respBody))

	return resp, nil
}

func botsResponse(req *http.Request, resp *http.Response) (*http.Response, error) {
	_, err := readRequestBody(req)
	if err != nil {
		return nil, err
	}

	respBody := []byte(botsResponseString)

	resp.ContentLength = int64(len(respBody))
	resp.Body = ioutil.NopCloser(bytes.NewReader(respBody))

	return resp, nil
}

func updateBotResponse(req *http.Request, resp *http.Response) (*http.Response, error) {
	_, err := readRequestBody(req)
	if err != nil {
		return nil, err
	}

	// TODO
	return resp, nil
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

func badRequestResponse(resp *http.Response) (*http.Response, error) {
	resp.StatusCode = http.StatusBadRequest
	resp.Status = http.StatusText(http.StatusBadRequest)
	resp.Body = ioutil.NopCloser(bytes.NewReader([]byte{}))

	return resp, nil
}
