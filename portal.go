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

type ProductsService interface {
	CreateProduct(input CreateProductInput) (*CreateProductOutput, error)
	GetProduct(id uint64) (*GetProductOutput, error)
	ListProducts(options *ListProductsOptions) (*ListProductsOutput, error)
	UpdateProduct(id uint64, input UpdateProductInput) (*UpdateProductOutput, error)
}

type CataloguesService interface {
	CreateCatalogue(input CreateCatalogueInput) (*CreateCatalogueOutput, error)
	GetCatalogue(id uint64) (*GetCatalogueOutput, error)
	UpdateCatalogue(id uint64, params UpdateCatalogueInput) (*UpdateCatalogueOutput, error)
	DeleteCatalogue(id uint64) (*DeleteCatalogueOutput, error)
	ListCatalogues(opts *ListCataloguesOptions) (*ListCataloguesOutput, error)
}

type ProvidersService interface {
	CreateProvider(input CreateProviderInput) (*CreateProviderOutput, error)
	GetProvider(id uint64) (*GetProviderOutput, error)
	ListProviders(options *ListProvidersOptions) (*ListProvidersOutput, error)
	UpdateProvider(id uint64, input UpdateProviderInput) (*UpdateProviderOutput, error)
}

type Client struct {
	config     *Config
	users      *usersService
	catalogues CataloguesService
	providers  ProvidersService
	teams      *teamsService
	products   ProductsService
	orgs       *orgsService
}

func (c Client) Users() *usersService {
	return c.users
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

func (c Client) Providers() ProvidersService {
	return c.providers
}

func (c *Client) SetProviders(providers ProvidersService) {
	c.providers = providers
}

func (c Client) Teams() *teamsService {
	return c.teams
}

func (c Client) Orgs() *orgsService {
	return c.orgs
}

func (c *Client) SetProducts(products ProductsService) {
	c.products = products
}

func (c *Client) SetTeams(teams *teamsService) {
	c.teams = teams
}

func (c *Client) SetOrgs(orgs *orgsService) {
	c.orgs = orgs
}

func newPortal(config *Config) (*Client, error) {
	client := &Client{
		config: config,
	}

	client.users = &usersService{client: client}
	client.catalogues = &cataloguesService{client: client}
	client.products = &productsService{client: client}
	client.providers = &providersService{client: client}
	client.teams = &teamsService{client: client}
	client.orgs = &orgsService{client: client}

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
