package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// DEFAULT_URL is the default base URL for
// the OpenData API endpoint
const DEFAULT_URL string = "https://us.api.insight.rapid7.com/opendata"

// Client is an API client for Rapid7's
// Sonar dataset
type Client struct {
	apiToken   string
	baseURL    *url.URL
	httpClient http.Client
}

// NewClient constructs a new api client
func NewClient(apiToken string, baseURL string) (*Client, error) {
	httpClient := http.Client{}
	apiURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		apiToken:   apiToken,
		baseURL:    apiURL,
		httpClient: httpClient,
	}, nil
}

func isErrStatus(code int) bool {
	return code < http.StatusOK || code >= http.StatusBadRequest
}

// decodeErr will extract and return the api's error message
// as an error, reporting its own error if unable to do so
func decodeErr(resp *http.Response) error {
	return fmt.Errorf("Error sending request:\n%s", resp.Status)
}

func (c *Client) sendRequest(req *http.Request, outData interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("X-Api-Key", c.apiToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if isErrStatus(resp.StatusCode) {
		return decodeErr(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(&outData)
	if err != nil {
		return err
	}

	return nil
}
