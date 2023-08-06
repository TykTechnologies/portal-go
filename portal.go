// Copyright 2023 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	headerAuthorization   = "Authorization"
	headerAccept          = "Accept"
	headerContentType     = "Content-Type"
	defaultBaseURL        = "http://localhost:3001"
	defaultConnectTimeout = 60000
	defaultReadTimeout    = 60000
	defaultUserAgent      = "ua"
)

type Option func(*Client)

func WithSkipValidation() Option {
	return func(c *Client) {
		c.skipValidation = true
	}
}

func WithDialTimeout(d time.Duration) Option {
	return func(o *Client) {
		o.connectTimeout = d
	}
}

func WithReadTimeout(d time.Duration) Option {
	return func(o *Client) {
		o.readTimeout = d
	}
}

func WithHTTPClient(c HTTPClient) Option {
	return func(o *Client) {
		o.httpClient = c
	}
}

func WithBaseURL(url string) Option {
	return func(o *Client) {
		o.baseURL = url
	}
}

func WithToken(value string) Option {
	return func(o *Client) {
		o.token = value
	}
}

func WithInsecure(value bool) Option {
	return func(o *Client) {
		o.insecure = value
	}
}

func WithConnectTimeout(value time.Duration) Option {
	return func(o *Client) {
		o.connectTimeout = value
	}
}

func WithUserAgent(ua string) Option {
	return func(o *Client) {
		o.userAgent = ua
	}
}

func WithDebug(debug bool) Option {
	return func(o *Client) {
		o.debug = debug
	}
}

type Client struct {
	httpClient      HTTPClient
	connectTimeout  time.Duration
	readTimeout     time.Duration
	userAgent       string
	token           string
	debug           bool
	insecure        bool
	baseURL         string
	maxRetries      int
	minRetryBackoff time.Duration
	skipValidation  bool

	pages     Pages
	providers Providers
	plans     Plans
	users     Users
	orgs      Orgs
	products  Products
	catalogs  Catalogs
	ars       ARs
	apps      Apps
}

func (c Client) Apps() Apps {
	return c.apps
}

func (c *Client) SetApps(app Apps) {
	c.apps = app
}

func (c Client) ARs() ARs {
	return c.ars
}

func (c *Client) SetARs(ar ARs) {
	c.ars = ar
}

func (c Client) Catalogs() Catalogs {
	return c.catalogs
}

func (c *Client) SetCatalogs(catalogs Catalogs) {
	c.catalogs = catalogs
}

func (c Client) Products() Products {
	return c.products
}

func (c *Client) SetProducts(products Products) {
	c.products = products
}

func (c Client) Orgs() Orgs {
	return c.orgs
}

func (c *Client) SetOrgs(orgs Orgs) {
	c.orgs = orgs
}

func (c Client) Users() Users {
	return c.users
}

func (c *Client) SetUsers(users Users) {
	c.users = users
}

func (c Client) Providers() Providers {
	return c.providers
}

func (c *Client) SetProviders(providers Providers) {
	c.providers = providers
}

func (c Client) Plans() Plans {
	return c.plans
}

func (c *Client) SetPlans(plans Plans) {
	c.plans = plans
}

func (c Client) Pages() Pages {
	return c.pages
}

func (c *Client) SetPages(pages Pages) {
	c.pages = pages
}

func (c *Client) Apply(opts ...Option) {
	for _, opt := range opts {
		if opt == nil {
			continue
		}

		opt(c)
	}
}

func New(opts ...Option) (*Client, error) {
	return newClient(opts...)
}

func (c Client) validate() error {
	if c.token == "" {
		return fmt.Errorf("token is required")
	}

	return nil
}

func newClient(opts ...Option) (*Client, error) {
	client := &Client{
		baseURL:        defaultBaseURL,
		connectTimeout: defaultConnectTimeout,
	}

	if client.maxRetries == 0 {
		client.maxRetries = 3
	}

	if client.minRetryBackoff == 0 {
		client.minRetryBackoff = 100 * time.Millisecond
	}

	client.Apply(opts...)

	if err := client.validate(); err != nil {
		return nil, err
	}

	client.providers = &providers{client: client}
	client.plans = &plans{client: client}
	client.users = &users{client: client}
	client.orgs = &orgs{client: client}
	client.products = &products{client: client}
	client.catalogs = &catalogs{client: client}
	client.ars = &ars{client: client}
	client.pages = &pages{client: client}
	client.apps = &apps{client: client}

	return client, nil
}

func (c Client) NewRequest(
	ctx context.Context,
	method string,
	path string,
	body io.Reader,
	params url.Values, opts ...Option,
) (*http.Request, error) {
	newClient := c.copy(opts...)

	newPath, err := url.JoinPath(newClient.baseURL, path)
	if err != nil {
		return nil, err
	}

	var newBody io.Reader = http.NoBody
	if body != nil {
		newBody = body
	}

	req, err := http.NewRequestWithContext(ctx, method, newPath, newBody)
	if err != nil {
		return nil, err
	}

	req.Header.Add(headerAuthorization, newClient.token)
	req.Header.Add(headerAccept, "application/json")

	if body != nil {
		req.Header.Add(headerContentType, "application/json")
	}

	req.Header.Set("User-Agent", defaultUserAgent)
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	req.URL.RawQuery = params.Encode()

	return req, nil
}

func (c Client) newGetRequest(ctx context.Context, path string, params url.Values, opts ...Option) (*http.Request, error) {
	return c.NewRequest(ctx, http.MethodGet, path, nil, params)
}

func (c Client) newPostRequest(ctx context.Context, path string, body io.Reader, params url.Values, opts ...Option) (*http.Request, error) {
	return c.NewRequest(ctx, http.MethodPost, path, body, params)
}

func (c Client) newPutRequest(ctx context.Context, path string, body io.Reader, params url.Values, opts ...Option) (*http.Request, error) {
	return c.NewRequest(ctx, http.MethodPut, path, body, params)
}

func (c Client) newDeleteRequest(
	ctx context.Context,
	path string,
	body io.Reader,
	params url.Values,
	opts ...Option,
) (*http.Request, error) {
	return c.NewRequest(ctx, http.MethodDelete, path, body, params)
}

func (c Client) doGet(ctx context.Context, path string, params url.Values, opts ...Option) (*APIResponse, error) {
	req, err := c.newGetRequest(ctx, path, params, opts...)
	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) doPost(ctx context.Context, path string, body io.Reader, params url.Values, opts ...Option) (*APIResponse, error) {
	req, err := c.newPostRequest(ctx, path, body, params, opts...)
	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) doDelete(ctx context.Context, path string, body io.Reader, params url.Values, opts ...Option) (*APIResponse, error) {
	req, err := c.newDeleteRequest(ctx, path, body, params, opts...)
	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) doPut(ctx context.Context, path string, body io.Reader, params url.Values, opts ...Option) (*APIResponse, error) {
	req, err := c.newPutRequest(ctx, path, body, params, opts...)
	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(ctx, req, opts...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) copy(opts ...Option) Client {
	newClient := c

	newClient.Apply(opts...)

	return newClient
}

func (c Client) performRequest(ctx context.Context, req *http.Request, opts ...Option) (*APIResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	newClient := c.copy(opts...)

	var httpClient HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				//nolint:gosec
				InsecureSkipVerify: newClient.insecure,
			},
			DialContext: (&net.Dialer{
				Timeout: newClient.connectTimeout * time.Millisecond,
			}).DialContext,
		},
		Timeout: newClient.readTimeout * time.Millisecond,
	}

	if newClient.httpClient != nil {
		httpClient = newClient.httpClient
	}

	var (
		attempt  int
		httpResp *http.Response
		err      error
		respC    = make(chan APIResponse)
		errC     = make(chan error)
	)

	backoff := c.minRetryBackoff

	go func() {
		for attempt = 0; attempt < c.maxRetries; attempt++ {
			httpResp, err = httpClient.Do(req)
			if err != nil {
				retry := shouldRetry(err)
				if retry && attempt < c.maxRetries-1 {
					time.Sleep(backoff)
					backoff *= 2
					continue
				}

				errC <- err
				return
			}

			break
		}

		defer httpResp.Body.Close()

		body, err := io.ReadAll(httpResp.Body)
		if err != nil {
			errC <- err
			return
		}

		r := &APIResponse{
			Body:     body,
			Response: httpResp,
		}

		if err := checkError(r); err != nil {
			errC <- err
			return
		}

		respC <- APIResponse{
			Response: httpResp,
			Body:     body,
		}
	}()

	select {
	case <-ctx.Done():
		return nil, context.DeadlineExceeded
	case err := <-errC:
		return nil, err
	case resp := <-respC:
		return &resp, nil
	}
}

var (
	ErrNotFound  = errors.New("not found")
	ErrForbidden = errors.New("forbidden")
)

func checkError(resp *APIResponse) error {
	switch resp.Response.StatusCode {
	case 200, 201:
		return nil
	default:
		e := APIError{
			APIResponse: resp,
		}

		if err := json.Unmarshal(resp.Body, &e); err != nil {
			return err
		}
		return e
	}
}

type APIResponse struct {
	Response *http.Response
	Body     []byte
}

func (a APIResponse) Unmarshal(v interface{}) error {
	return json.Unmarshal(a.Body, &v)
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

func shouldRetry(err error) bool {
	return false
}

type Status struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
