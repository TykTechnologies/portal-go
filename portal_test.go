package edp

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type server struct {
	mux *http.ServeMux
	srv *httptest.Server
}

func NewServer(t *testing.T) *server {
	mux := http.NewServeMux()
	srv := httptest.NewServer(mux)

	return &server{
		mux: mux,
		srv: srv,
	}
}

func (s server) Close() {
	s.srv.Close()
}

func httpParse(t *testing.T, filename string) *http.Response {
	return httpParseWithRequest(t, filename, nil)
}

func httpParseWithRequest(t *testing.T, filename string, r *http.Request) *http.Response {
	fixture := strings.NewReader(readFixture(t, filename))

	resp, err := http.ReadResponse(bufio.NewReader(fixture), r)
	assert.NoError(t, err)
	return resp
}

func readFixture(t *testing.T, filename string) string {
	data, err := os.ReadFile(filepath.Join("testdata", filename))
	assert.NoError(t, err)

	return string(data[:])
}

func assertMethod(t *testing.T, method string, r *http.Request) {
	assert.Equal(t, method, r.Method)
}

func assertHeader(t *testing.T, r *http.Request, header, value string) {
	assert.Equal(t, value, r.Header.Get(header))
}

func TestNew(t *testing.T) {
	tt := map[string]struct {
		opt  []Option
		want Client
		err  bool
	}{
		"no option": {
			err: true,
		},
		"with base url": {
			want: Client{baseURL: "http://example.com", token: "random token"},
			opt:  []Option{WithBaseURL("http://example.com"), WithToken("random token")},
		},
		"with token": {
			want: Client{token: "random token", baseURL: defaultBaseURL},
			opt:  []Option{WithToken("random token")},
		},
	}

	t.Parallel()

	for k, v := range tt {
		t.Run(k, func(t *testing.T) {
			client, err := New(v.opt...)
			if v.err {
				assert.Error(t, err)
				assert.Nil(t, client)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, client)
				assert.Equal(t, v.want.baseURL, client.baseURL)
				assert.Equal(t, v.want.token, client.token)
			}
		})
	}
}

func TestOptions(t *testing.T) {
	tt := map[string]struct {
		opt  Option
		want Client
	}{
		"with base url": {
			want: Client{baseURL: "http://example.com"},
			opt:  WithBaseURL("http://example.com"),
		},
		"with token": {
			want: Client{token: "random token"},
			opt:  WithToken("random token"),
		},
		"with insecure": {
			want: Client{insecure: true},
			opt:  WithInsecure(true),
		},
		"with secure": {
			want: Client{insecure: false},
			opt:  WithInsecure(false),
		},
		"with read timeout": {
			want: Client{insecure: false, readTimeout: 10 * time.Second},
			opt:  WithReadTimeout(10 * time.Second),
		},
		"with connect timeout": {
			want: Client{insecure: false, connectTimeout: 10 * time.Second},
			opt:  WithConnectTimeout(10 * time.Second),
		},
		"with user agent": {
			want: Client{insecure: false, userAgent: "Test"},
			opt:  WithUserAgent("Test"),
		},
	}

	t.Parallel()

	for k, v := range tt {
		t.Run(k, func(t *testing.T) {
			client := &Client{}
			client.Apply(v.opt)

			assert.Equal(t, v.want.baseURL, client.baseURL)
			assert.Equal(t, v.want.token, client.token)
			assert.Equal(t, v.want.insecure, client.insecure)
			assert.Equal(t, v.want.readTimeout, client.readTimeout)
		})
	}
}
