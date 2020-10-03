package clubhouse

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Tag     string `json:"tag"`
}

const DefaultURL = "https://api.clubhouse.io/api/v3"

func New(token string) *Client {
	return &Client{
		Token:      token,
		HTTPClient: &http.Client{Timeout: 5 * time.Second},
		URL:        DefaultURL,
	}
}

type Client struct {
	Debug      bool
	HTTPClient *http.Client
	Token      string
	URL        string
}

func (c *Client) makeRequest(method, path string, body interface{}) (*http.Request, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s?token=%s", c.URL, path, c.Token)
	req, err := http.NewRequest(method, url, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	if c.Debug {
		dumpRequest(req)
	}

	return req, nil
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		dumpResponse(resp)
	}

	return resp, nil
}

func (c *Client) decode(resp *http.Response, data interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var er ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&er); err != nil {
			return err
		}
		return errors.New(er.Message)
	}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return err
	}

	return nil
}

func (c *Client) get(path string, data interface{}) error {
	req, err := c.makeRequest(http.MethodGet, path, nil)
	if err != nil {
		return err
	}

	resp, err := c.do(req)
	if err != nil {
		return err
	}

	if err := c.decode(resp, data); err != nil {
		return err
	}

	return nil
}

func (c *Client) post(path string, body, data interface{}) error {
	req, err := c.makeRequest(http.MethodPost, path, body)
	if err != nil {
		return err
	}

	resp, err := c.do(req)
	if err != nil {
		return err
	}

	if err := c.decode(resp, data); err != nil {
		return err
	}

	return nil
}

func (c *Client) put(path string, body, data interface{}) error {
	req, err := c.makeRequest(http.MethodPut, path, body)
	if err != nil {
		return err
	}

	resp, err := c.do(req)
	if err != nil {
		return err
	}

	if err := c.decode(resp, data); err != nil {
		return err
	}

	return nil
}

func dumpResponse(resp *http.Response) {
	b, _ := httputil.DumpResponse(resp, true)
	_, _ = fmt.Fprintln(os.Stderr, string(b))
}

func dumpRequest(req *http.Request) {
	b, _ := httputil.DumpRequest(req, true)
	_, _ = fmt.Fprintln(os.Stderr, string(b))
}
