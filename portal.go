package edp

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
)

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

type Client struct {
	httpClient      HTTPClient
	connectTimeout  time.Duration
	readTimeout     time.Duration
	token           string
	debug           bool
	insecure        bool
	baseURL         string
	providers       ProvidersService
	plans           PlansService
	users           UsersService
	orgs            OrgsService
	products        ProductsService
	catalogues      CataloguesService
	accessRequests  AccessRequestsService
	maxRetries      int
	minRetryBackoff time.Duration
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

func (c Client) validate() error {
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

	client.apply(opts...)

	if err := client.validate(); err != nil {
		return nil, err
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
	newClient := c.copy(opts...)

	newPath, err := url.JoinPath(newClient.baseURL, path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, newPath, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(headerAuthorization, newClient.token)
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

func (c Client) doGet(ctx context.Context, path string, params url.Values, opts ...Option) (*internalResponse, error) {
	req, err := c.newGetRequest(path, params, opts...)
	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) doPost(ctx context.Context, path string, body io.Reader, params url.Values, opts ...Option) (*internalResponse, error) {
	req, err := c.newPostRequest(path, body, params, opts...)

	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) doDelete(ctx context.Context, path string, body io.Reader, params url.Values, opts ...Option) (*internalResponse, error) {
	req, err := c.newDeleteRequest(path, body, params, opts...)

	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) doPut(ctx context.Context, path string, body io.Reader, params url.Values, opts ...Option) (*internalResponse, error) {
	req, err := c.newPutRequest(path, body, params, opts...)

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

	newClient.apply(opts...)

	return newClient
}

func (c Client) performRequest(ctx context.Context, req *http.Request, opts ...Option) (*internalResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	newClient := c.copy(opts...)

	var httpClient HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
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
		respC    = make(chan internalResponse, 0)
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

		if err := checkError(httpResp.StatusCode, body); err != nil {
			errC <- err
			return
		}

		respC <- internalResponse{
			Response: httpResp,
			body:     body,
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

func shouldRetry(err error) bool {
	return false
}
