package portal

import (
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
)

type Config struct {
	Token          string
	Debug          bool
	Insecure       bool
	BaseURL        string
	ConnectTimeout time.Duration
	ReadTimeout    time.Duration
	HTTPClient     HTTPClient
}

type Option func(*Client)

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

func WithNoVeriySSL(value bool) Option {
	return func(o *Client) {
		o.noVerifySSL = value
	}
}

func WithConnectTimeout(value time.Duration) Option {
	return func(o *Client) {
		o.connectTimeout = value
	}
}

type Client struct {
	httpClient     HTTPClient
	connectTimeout time.Duration
	readTimeout    time.Duration
	token          string
	debug          bool
	noVerifySSL    bool
	baseURL        string
	config         *Config
	providers      ProvidersService
	plans          PlansService
	users          UsersService
	orgs           OrgsService
	products       ProductsService
	catalogues     CataloguesService
	accessRequests AccessRequestsService
}

func (c Client) AccessRequests() AccessRequestsService {
	return c.accessRequests
}

func (c *Client) SetAccessRequests(ar AccessRequestsService) {
	c.accessRequests = ar
}

func (c Client) Catalogues() CataloguesService {
	return c.catalogues
}

func (c *Client) SetCatalogues(catalogues CataloguesService) {
	c.catalogues = catalogues
}

func (c Client) Products() ProductsService {
	return c.products
}

func (c *Client) SetProducts(products ProductsService) {
	c.products = products
}

func (c Client) Orgs() OrgsService {
	return c.orgs
}

func (c *Client) SetOrgs(orgs OrgsService) {
	c.orgs = orgs
}

func (c Client) Users() UsersService {
	return c.users
}

func (c *Client) SetUsers(users UsersService) {
	c.users = users
}

func (c Client) Providers() ProvidersService {
	return c.providers
}

func (c *Client) SetProviders(providers ProvidersService) {
	c.providers = providers
}

func (c Client) Plans() PlansService {
	return c.plans
}

func (c *Client) SetPlans(plans PlansService) {
	c.plans = plans
}

func (c *Client) apply(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
}

func New(opts ...Option) (*Client, error) {
	return newClient(opts...)
}

func newClient(opts ...Option) (*Client, error) {
	client := &Client{}

	client.apply(opts...)

	if client.baseURL == "" {
		client.baseURL = defaultBaseURL
	}

	if client.connectTimeout == 0 {
		client.connectTimeout = defaultConnectTimeout
	}

	client.providers = &providersService{client: client}
	client.plans = &plansService{client: client}
	client.users = &usersService{client: client}
	client.orgs = &orgsService{client: client}
	client.products = &productsService{client: client}
	client.catalogues = &cataloguesService{client: client}
	client.accessRequests = &accessRequestsService{client: client}

	return client, nil
}

func (c Client) newRequest(method string, path string, body io.Reader, params url.Values, opts ...Option) (*http.Request, error) {
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

func (c Client) newGetRequest(path string, params url.Values, opts ...Option) (*http.Request, error) {
	return c.newRequest(http.MethodGet, path, nil, params)
}

func (c Client) newPostRequest(path string, body io.Reader, params url.Values, opts ...Option) (*http.Request, error) {
	return c.newRequest(http.MethodPost, path, body, params)
}

func (c Client) newPutRequest(path string, body io.Reader, params url.Values, opts ...Option) (*http.Request, error) {
	return c.newRequest(http.MethodPut, path, body, params)
}

func (c Client) newDeleteRequest(path string, body io.Reader, params url.Values, opts ...Option) (*http.Request, error) {
	return c.newRequest(http.MethodDelete, path, body, params)
}

func (c Client) doGet(path string, params url.Values, opts ...Option) (*internalResponse, error) {
	req, err := c.newGetRequest(path, params, opts...)
	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) doPost(path string, body io.Reader, params url.Values, opts ...Option) (*internalResponse, error) {
	req, err := c.newPostRequest(path, body, params, opts...)

	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) doDelete(path string, body io.Reader, params url.Values, opts ...Option) (*internalResponse, error) {
	req, err := c.newDeleteRequest(path, body, params, opts...)

	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) doPut(path string, body io.Reader, params url.Values, opts ...Option) (*internalResponse, error) {
	req, err := c.newPutRequest(path, body, params, opts...)

	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req, opts...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) performRequest(req *http.Request, opts ...Option) (*internalResponse, error) {
	var httpClient HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: c.config.Insecure,
			},
			DialContext: (&net.Dialer{
				Timeout: c.config.ConnectTimeout * time.Millisecond,
			}).DialContext,
		},
		Timeout: c.config.ReadTimeout * time.Millisecond,
	}

	if c.config.HTTPClient != nil {
		httpClient = c.config.HTTPClient
	}

	httpResp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	if err := checkError(httpResp.StatusCode, body); err != nil {
		return nil, err
	}

	resp := internalResponse{
		Response: httpResp,
		body:     body,
	}

	return &resp, nil
}

var (
	ErrNotFound  = errors.New("not found")
	ErrForbidden = errors.New("forbidden")
)

func checkError(code int, body []byte) error {
	if code >= 200 && code < 300 {
		return nil
	}

	var e Error
	if err := json.Unmarshal(body, &e); err != nil {
		return err
	}

	return e
}

type internalResponse struct {
	*http.Response
	body []byte
}

func (a internalResponse) Unmarshal(v interface{}) error {
	return json.Unmarshal(a.body, &v)
}

type internalError struct {
	internalResponse
}

func (e internalError) Error() string {
	return fmt.Sprintf(
		"%v %v %v",
		e.internalResponse.Response.Request.Method,
		e.internalResponse.Response.Request.URL,
		e.internalResponse.Response.StatusCode,
	)
}

type Error struct {
	Kind    string
	Err     interface{}
	Message string
}

func (e Error) Error() string {
	return "API error"
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

func String(v string) *string {
	return &v
}

func StringValue(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func Int64(v int64) *int64 {
	return &v
}

func Int64Value(s *int64) int64 {
	if s == nil {
		return 0
	}

	return *s
}
