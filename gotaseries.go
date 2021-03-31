package gotaseries

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

type service struct {
	client *Client
}

// Client manage communication with the Betaseries API.
type Client struct {
	baseURL   url.URL
	userAgent string
	ApiKey    string
	Token     *string

	httpClient *http.Client

	common service
	Show   *ShowService
}

// NewClient returns a new Betaseries API client. To use API methods which require
// authentication, use token parameter otherwise set it to nil
func NewClient(apiKey string, token *string) *Client {
	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	u, err := url.Parse("https://api.betaseries.com")
	if err != nil {
		log.Fatal(err)
	}

	c := &Client{
		baseURL:    *u,
		userAgent:  "gotaseries",
		httpClient: httpClient,
		ApiKey: apiKey,
		Token: token,
	}
	c.common.client = c
	c.Show = (*ShowService)(&c.common)

	return c
}

func (c *Client) newRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("X-BetaSeries-Version", "3.0")
	req.Header.Set("X-BetaSeries-Key", c.ApiKey)
	if c.Token != nil {
		req.Header.Set("X-BetaSeries-Token", *c.Token)
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)

	return resp, err
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface {}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)

	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()

	return u.String(), nil
}