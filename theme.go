package portal

type Theme struct {
	Author  string `json:"Author,omitempty"`
	ID      string `json:"ID,omitempty"`
	Name    string `json:"Name,omitempty"`
	Path    string `json:"Path,omitempty"`
	Status  string `json:"Status,omitempty"`
	Version string `json:"Version,omitempty"`
}

type Status struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type Err struct {
	Status string   `json:"status,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

type ContentBlock struct {
	Content string `json:"Content,omitempty"`
	Name    string `json:"Name,omitempty"`
	ID      int    `json:"ID,omitempty"`
	PageID  int    `json:"PageID,omitempty"`
}

type ContentBlockInput struct {
	Content string `json:"Content,omitempty"`
	Name    string `json:"Name,omitempty"`
}

type Configs struct {
	Blog               Blog      `json:"Blog,omitempty"`
	Database           Database  `json:"Database,omitempty"`
	Forms              Forms     `json:"Forms,omitempty"`
	HostPort           int       `json:"HostPort,omitempty"`
	JwtSigningKey      string    `json:"JwtSigningKey,omitempty"`
	LicenseKey         string    `json:"LicenseKey,omitempty"`
	LogFormat          string    `json:"LogFormat,omitempty"`
	LogLevel           string    `json:"LogLevel,omitempty"`
	PortalAPISecret    string    `json:"PortalAPISecret,omitempty"`
	ProductDocRenderer string    `json:"ProductDocRenderer,omitempty"`
	RefreshInterval    int       `json:"RefreshInterval,omitempty"`
	S3                 S3        `json:"S3,omitempty"`
	Site               Site      `json:"Site,omitempty"`
	Storage            string    `json:"Storage,omitempty"`
	StoreSessionName   string    `json:"StoreSessionName,omitempty"`
	TLSConfig          TLSConfig `json:"TLSConfig,omitempty"`
	Theming            Theming   `json:"Theming,omitempty"`
}
type Blog struct {
	AllowFormSubmission bool `json:"AllowFormSubmission,omitempty"`
	Enable              bool `json:"Enable,omitempty"`
}
type Database struct {
	ConnectionString string `json:"ConnectionString,omitempty"`
	Dialect          string `json:"Dialect,omitempty"`
	EnableLogs       bool   `json:"EnableLogs,omitempty"`
	MaxRetries       int    `json:"MaxRetries,omitempty"`
	RetryDelay       int    `json:"RetryDelay,omitempty"`
}
type Forms struct {
	Enable bool `json:"Enable,omitempty"`
}
type S3 struct {
	ACL         string `json:"ACL,omitempty"`
	AccessKey   string `json:"AccessKey,omitempty"`
	Bucket      string `json:"Bucket,omitempty"`
	Endpoint    string `json:"Endpoint,omitempty"`
	PresignURLs bool   `json:"PresignURLs,omitempty"`
	Region      string `json:"Region,omitempty"`
	SecretKey   string `json:"SecretKey,omitempty"`
}
type Site struct {
	Enable bool `json:"Enable,omitempty"`
}
type Certificates struct {
	CertFile string `json:"CertFile,omitempty"`
	KeyFile  string `json:"KeyFile,omitempty"`
	Name     string `json:"Name,omitempty"`
}
type TLSConfig struct {
	Certificates       []Certificates `json:"Certificates,omitempty"`
	Enable             bool           `json:"Enable,omitempty"`
	InsecureSkipVerify bool           `json:"InsecureSkipVerify,omitempty"`
	MinVersion         int            `json:"MinVersion,omitempty"`
}
type Theming struct {
	Path  string `json:"Path,omitempty"`
	Theme string `json:"Theme,omitempty"`
}

type AudienceInput struct {
	OrganisationID int `json:"OrganisationID,omitempty"`
	TeamID         int `json:"TeamID,omitempty"`
}

type Audience struct {
	ID             int    `json:"ID,omitempty"`
	OrganisationID int    `json:"OrganisationID,omitempty"`
	TeamID         int    `json:"TeamID,omitempty"`
	CreatedAt      string `json:"CreatedAt,omitempty"`
	UpdatedAt      string `json:"UpdatedAt,omitempty"`
}

type CatalogueDetails struct {
	ID               int            `json:"ID,omitempty"`
	Name             string         `json:"Name,omitempty"`
	CreatedAt        string         `json:"CreatedAt,omitempty"`
	UpdatedAt        string         `json:"UpdatedAt,omitempty"`
	OrgCatalogues    []OrgCatalogue `json:"OrgCatalogues,omitempty"`
	Plans            []string       `json:"Plans,omitempty"`
	Products         []string       `json:"Products,omitempty"`
	VisibilityStatus string         `json:"VisibilityStatus,omitempty"`
}
type OrgCatalogue struct {
	Catalogue    string `json:"Catalogue,omitempty"`
	ID           int    `json:"ID,omitempty"`
	Name         string `json:"Name,omitempty"`
	Organisation string `json:"Organisation,omitempty"`
}

type CatalogueSummary struct {
	ID               int    `json:"ID,omitempty"`
	Name             string `json:"Name,omitempty"`
	NameWithSlug     string `json:"NameWithSlug,omitempty"`
	VisibilityStatus string `json:"VisibilityStatus,omitempty"`
}
