package mangadexv5

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/dyninc/qstring"
	"github.com/pkg/errors"
	"go.uber.org/ratelimit"
)

type Client struct {
	token      *LoginToken
	httpClient *http.Client
	limiter    ratelimit.Limiter
}

type Config func(*Client) *Client

func NewClient(configs ...Config) *Client {
	c := &Client{
		httpClient: http.DefaultClient,
		limiter:    ratelimit.New(5),
	}
	for _, config := range configs {
		c = config(c)
	}
	return c
}

func RateLimit(duration time.Duration) Config {
	return func(c *Client) *Client {
		c.limiter = ratelimit.New(int(duration))
		return c
	}
}

func HttpClient(client *http.Client) Config {
	return func(c *Client) *Client {
		c.httpClient = client
		return c
	}
}

func (c *Client) request(method, url string, body io.Reader) (*http.Response, error) {
	c.limiter.Take()
	r, err := http.NewRequest(method, "https://api.mangadex.org"+url, body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	if c.token != nil {
		r.Header.Add("Authorization", "Bearer "+c.token.Session)
	}
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, NewAPIResponseError(resp)
	}

	return resp, nil
}

func (c *Client) post(url string, body, result interface{}) error {
	bodyReader := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(bodyReader).Encode(body)
	if err != nil {
		return errors.Wrap(err, "failed to encode request body")
	}

	resp, err := c.request("POST", url, bodyReader)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return errors.Wrap(err, "failed to decode respose body")
	}

	return nil
}

func (c *Client) get(url string, params, result interface{}) error {
	var err error
	var q string

	if params != nil {
		q, err = qstring.MarshalString(params)
		if err != nil {
			return err
		}
		q = "?" + q
	}

	resp, err := c.request("GET", url+q, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return errors.Wrap(err, "failed to decode respose body")
	}

	return nil
}
