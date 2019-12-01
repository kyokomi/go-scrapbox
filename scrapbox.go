package scrapbox

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	tokenKey = "connect.sid"
	baseURL  = "https://scrapbox.io/api/"
)

// Client scrapbox client
type Client struct {
	token      string
	httpClient *http.Client

	PageService *PageService
}

// NewClient return a create Client
func NewClient(token string) *Client {
	c := &Client{token: token, httpClient: http.DefaultClient}
	c.PageService = &PageService{Client: c}
	return c
}

// WithHTTPClient setup http.Client
func (c *Client) WithHTTPClient(httpClient *http.Client) *Client {
	if httpClient != nil {
		c.httpClient = httpClient
	}
	return c
}

func (c *Client) createURL(action string, projectName string) (*url.URL, error) {
	return url.Parse(fmt.Sprintf(baseURL+"%s/%s", action, projectName))
}

func (c *Client) newRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: tokenKey, Value: c.token})

	return req, nil
}

func (c *Client) doAndJSONDecode(req *http.Request, v interface{}) error {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 400 {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errorResponse); err != nil {
			return err
		}
		return fmt.Errorf("status %d %s: %s", errorResponse.StatusCode, errorResponse.Name, errorResponse.Message)
	}

	return json.NewDecoder(res.Body).Decode(v)
}
