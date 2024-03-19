package portal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

type Option func(*Client)

func NewClient(baseURL, secret string, opts ...Option) (*Client, error) {
	b, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	newClient := &Client{
		baseURL: b,
		secret:  secret,
	}

	for _, opt := range opts {
		if opt == nil {
			continue
		}

		opt(newClient)
	}

	newClient.setup()

	return newClient, err
}

type Client struct {
	client    *http.Client
	common    service
	baseURL   *url.URL
	userAgent string
	secret    string

	ARs        *ARsService
	Apps       *AppsService
	Catalogues *CataloguesService
	EAs        *EAsService
	Orgs       *OrgsService
	Pages      *PagesService
	Plans      *PlansService
	Products   *ProductsService
	Providers  *ProvidersService
	Themes     *ThemesService
	Users      *UsersService
}

func (c *Client) setup() error {
	if c.client == nil {
		c.client = &http.Client{}
	}

	if c.baseURL == nil {
		baseURL, err := url.Parse("http://localhost:3001/portal-api/")
		if err != nil {
			return err
		}

		c.baseURL = baseURL
	}

	if c.userAgent == "" {
		c.userAgent = "portal"
	}

	c.common.client = c

	c.ARs = (*ARsService)(&c.common)
	c.Apps = (*AppsService)(&c.common)
	c.Catalogues = (*CataloguesService)(&c.common)
	c.EAs = (*EAsService)(&c.common)
	c.Orgs = (*OrgsService)(&c.common)
	c.Pages = (*PagesService)(&c.common)
	c.Plans = (*PlansService)(&c.common)
	c.Products = (*ProductsService)(&c.common)
	c.Providers = (*ProvidersService)(&c.common)
	c.Themes = (*ThemesService)(&c.common)
	c.Users = (*UsersService)(&c.common)

	return nil
}

func jsonEncode(w io.Writer, body interface{}) error {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	return enc.Encode(body)
}

func (c *Client) NewRequestWithOptions(ctx context.Context, method, urlPath string, body, options interface{}, opts ...RequestOption) (*http.Request, error) {
	if v, ok := body.(Validator); ok {
		err := v.Validate()
		if err != nil {
			return nil, err
		}
	}

	req, err := c.NewRequest(ctx, method, urlPath, body, opts...)
	if err != nil {
		log.Println("Unable to create request")
		return nil, err
	}

	if err := requestWithOptions(req, options); err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) NewRequest(ctx context.Context, method, urlPath string, body interface{}, opts ...RequestOption) (*http.Request, error) {
	if !strings.HasSuffix(c.baseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.baseURL)
	}

	u := c.baseURL.JoinPath(urlPath)

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}

		err := jsonEncode(buf, body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")

	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	req.Header.Set("Authorization", c.secret)

	for _, opt := range opts {
		opt(req)
	}

	return req, nil
}

type RequestOption func(req *http.Request)

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.performRequest(ctx, req)
	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		jsonErr := json.NewDecoder(resp.Body).Decode(v)
		if jsonErr == io.EOF {
			jsonErr = nil
		}

		if jsonErr != nil {
			err = jsonErr
		}
	}

	return resp, err
}

var ErrInvalidContext = errors.New("no context")

func (c *Client) performRequest(ctx context.Context, req *http.Request) (*Response, error) {
	if ctx == nil {
		return nil, ErrInvalidContext
	}

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	err = checkResponse(resp)
	if err != nil {
		defer resp.Body.Close()
	}

	return newResponse(resp), err
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errResp := &ResponseError{Response: r}

	data, err := io.ReadAll(r.Body)
	log.Println(string(data))

	if err == nil && data != nil {
		err = json.Unmarshal(data, errResp)
		if err != nil {
			errResp = &ResponseError{Response: r}
		}
	}

	r.Body = io.NopCloser(bytes.NewBuffer(data))

	switch {
	default:
		return errResp
	}
}

type ResponseError struct {
	Response *http.Response `json:"-"`
	Message  string         `json:"message"`
	Errors   []string       `json:"errors"`
}

func (r *ResponseError) Error() string {
	if r.Response != nil && r.Response.Request != nil {
		return fmt.Sprintf("%v %v: %d %v %+v", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message, r.Errors)
	}

	if r.Response != nil {
		return fmt.Sprintf("%d %v %+v", r.Response.StatusCode, r.Message, r.Errors)
	}

	return fmt.Sprintf("%v %+v", r.Message, r.Errors)
}

type service struct {
	client *Client
}

type Response struct {
	*http.Response
}

func requestWithOptions(req *http.Request, opts interface{}) error {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return nil
	}

	qs, err := query.Values(opts)
	if err != nil {
		return err
	}

	req.URL.RawQuery = qs.Encode()
	return nil
}

func Bool(v bool) *bool { return &v }

func Int(v int) *int { return &v }

func Int64(v int64) *int64 { return &v }

func String(v string) *string { return &v }

func newResponse(r *http.Response) *Response {
	return &Response{Response: r}
}
