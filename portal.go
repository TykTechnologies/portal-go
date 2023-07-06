package portal

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

const (
	headerAuthorization      = "Authorization"
	headerAccept             = "Accept"
	headerContentType        = "Content-Type"
	defaultBaseURL           = "http://localhost:3001"
	defaultConnectionTimeout = 30 * time.Second
)

type Config struct {
	Token             string
	Debug             bool
	Insecure          bool
	BaseURL           string
	HTTPClient        *http.Client
	ConnectionTimeout time.Duration
}

type Client struct {
	config    *Config
	providers ProvidersService
}

func (c Client) Providers() ProvidersService {
	return c.providers
}

func (c *Client) SetProviders(providers ProvidersService) {
	c.providers = providers
}

func New(config *Config) (*Client, error) {
	return newClient(config)
}

func newClient(config *Config) (*Client, error) {
	if config == nil {
		return nil, errors.New("config should not be empty")
	}

	if config.BaseURL == "" {
		config.BaseURL = defaultBaseURL
	}

	if config.ConnectionTimeout == 0 {
		config.ConnectionTimeout = defaultConnectionTimeout
	}

	client := &Client{
		config: config,
	}

	client.providers = &providersService{client: client}

	return client, nil
}

func (c Client) newRequest(method string, path string, body io.Reader, params url.Values) (*http.Request, error) {
	fullURL, err := url.JoinPath(c.config.BaseURL, path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(headerAuthorization, c.config.Token)
	req.Header.Add(headerAccept, "application/json")
	req.Header.Add(headerContentType, "application/json")

	req.URL.RawQuery = params.Encode()

	return req, nil
}

func (c Client) newGetRequest(path string, params url.Values) (*http.Request, error) {
	return c.newRequest(http.MethodGet, path, nil, params)
}

func (c Client) newPostRequest(path string, body io.Reader, params url.Values) (*http.Request, error) {
	return c.newRequest(http.MethodPost, path, body, params)
}

func (c Client) newPutRequest(path string, body io.Reader, params url.Values) (*http.Request, error) {
	return c.newRequest(http.MethodPut, path, body, params)
}

func (c Client) newDeleteRequest(path string, body io.Reader, params url.Values) (*http.Request, error) {
	return c.newRequest(http.MethodDelete, path, body, params)
}

func (c Client) doGet(path string, params url.Values) (*apiResponse, error) {
	req, err := c.newGetRequest(path, params)
	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) doPost(path string, body io.Reader, params url.Values) (*apiResponse, error) {
	req, err := c.newPostRequest(path, body, params)

	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) doDelete(path string, body io.Reader, params url.Values) (*apiResponse, error) {
	req, err := c.newDeleteRequest(path, body, params)

	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) doPut(path string, body io.Reader, params url.Values) (*apiResponse, error) {
	req, err := c.newPutRequest(path, body, params)

	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) performRequest(req *http.Request) (*apiResponse, error) {
	httpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: c.config.Insecure,
			},
			DialContext: (&net.Dialer{
				Timeout: c.config.ConnectionTimeout,
			}).DialContext,
		},
	}

	if c.config.HTTPClient != nil {
		httpClient = *c.config.HTTPClient
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 299 {
		return nil, handleAPIError(resp.StatusCode, apiError{
			Body: body,
		})
	}

	return &apiResponse{
		Body: body,
	}, nil
}

var (
	ErrNotFound  = errors.New("not found")
	ErrForbidden = errors.New("forbidden")
)

func handleAPIError(code int, apiErr apiError) error {
	if code == 400 || code == 422 {
		var e Error
		if err := json.Unmarshal(apiErr.Body, &e); err != nil {
			return err
		}

		return e
	}

	return apiErr
}

type apiResponse struct {
	Body []byte
}

func (a apiResponse) Parse(v interface{}) error {
	return json.Unmarshal(a.Body, &v)
}

type apiError struct {
	Body []byte
}

func (e apiError) Error() string {
	return "API error"
}

type Error struct {
	Errors []string `json:"errors"`
}

func (e Error) Error() string {
	return "API error"
}
