package scrapbox

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/xerrors"
)

const (
	tokenKey = "connect.sid"
	baseURL  = "https://scrapbox.io/api/"
)

var (
	errNotFound = errors.New("not found")
)

// Client scrapbox client
type Client struct {
	token      string
	httpClient *http.Client

	Page *PageService
}

// NewClient return a create Client
func NewClient(token string) *Client {
	c := &Client{token: token, httpClient: http.DefaultClient}
	c.Page = &PageService{Client: c}
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
	res, err := c.doWithErrorHandling(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(v)
}

func (c *Client) doAndGetText(req *http.Request) (string, error) {
	res, err := c.doWithErrorHandling(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (c *Client) doAndGetRedirectURL(req *http.Request) (bool, *url.URL, error) {
	res, err := c.doWithErrorHandling(req)
	if err != nil {
		if errors.Is(errNotFound, err) {
			return false, nil, nil
		}
		return false, nil, err
	}
	defer res.Body.Close()

	return true, res.Request.URL, nil
}

// required close
func (c *Client) doWithErrorHandling(req *http.Request) (*http.Response, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httpClient do request error: %w", err)
	}

	if res.StatusCode > 400 {
		defer res.Body.Close()

		if res.StatusCode == http.StatusNotFound {
			return nil, errNotFound
		}

		var errorResponse ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errorResponse); err != nil {
			data, _ := ioutil.ReadAll(res.Body)
			return nil, xerrors.Errorf("status %d: %s", res.StatusCode, string(data))
		}
		return nil, xerrors.Errorf("status %d %s: %s", errorResponse.StatusCode, errorResponse.Name, errorResponse.Message)
	}

	return res, nil
}
