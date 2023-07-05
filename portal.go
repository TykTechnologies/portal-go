package portal

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
)

type Config struct {
	Token      string
	Debug      bool
	Insecure   bool
	BaseURL    string
	HTTPClient *http.Client
}

type Client struct {
	config     *Config
	Users      *Users
	Catalogues *Catalogues
	Providers  *Providers
	Teams      *Teams
	Products   *Products
	Orgs       *Orgs
}

func newPortal(config *Config) (*Client, error) {
	client := &Client{
		config: config,
	}

	client.Users = &Users{client: client}
	client.Catalogues = &Catalogues{client: client}
	client.Products = &Products{client: client}
	client.Providers = &Providers{client: client}
	client.Teams = &Teams{client: client}
	client.Orgs = &Orgs{client: client}

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

func (c Client) performRequest(req *http.Request) (*http.Response, error) {
	httpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: c.config.Insecure,
			},
		},
	}

	if c.config.HTTPClient != nil {
		httpClient = *c.config.HTTPClient
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
