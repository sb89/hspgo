package hsp

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

// BaseURL sets the base URL used by Client to make requests to the HSP API
func BaseURL(path string) func(*Client) {
	return func(c *Client) {
		c.basePath = path
	}
}

// HTTPClient sets the *http.Client used by Client to make requests to the HSP API
func HTTPClient(client *http.Client) func(*Client) {
	return func(c *Client) {
		c.httpClient = client
	}
}

// Client used to communicate with HSP API
type Client struct {
	basePath   string
	httpClient *http.Client
	auth       string

	*serviceMetricsService
	*serviceDetailsService
}

// NewClient creates a new Client
func NewClient(email string, password string, options ...func(*Client)) *Client {
	creds := fmt.Sprintf("%s:%s", email, password)
	sEncCreds := base64.StdEncoding.EncodeToString([]byte(creds))

	c := &Client{
		auth:     sEncCreds,
		basePath: "https://hsp-prod.rockshore.net/api/v1",
	}

	c.serviceMetricsService = &serviceMetricsService{client: c}
	c.serviceDetailsService = &serviceDetailsService{client: c}

	for _, option := range options {
		option(c)
	}

	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}

	return c
}

func (c *Client) newRequest(path string, body interface{}) (*http.Request, *Error) {
	url := c.basePath + path

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, getError(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return nil, getError(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic "+c.auth)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, *Error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, getError(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, getHTTPError(resp)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return nil, getError(err)
	}

	return resp, nil
}
