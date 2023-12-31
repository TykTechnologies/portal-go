// Copyright 2023 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApps_Get(t *testing.T) {
	srv := NewServer(t)
	defer srv.Close()

	token := "TOKEN"

	srv.mux.HandleFunc("/portal-api/apps", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpParse(t, "portal-api/apps.txt")
		defer httpResponse.Body.Close()

		assertMethod(t, "GET", r)
		assertHeader(t, r, "Authorization", token)

		w.WriteHeader(httpResponse.StatusCode)
		_, err := io.Copy(w, httpResponse.Body)
		assert.NoError(t, err)
	})

	client, err := New(
		WithBaseURL(srv.srv.URL),
		WithToken(token),
	)
	assert.NoError(t, err)

	resp, err := client.Apps().ListApps(context.Background())
	assert.NoError(t, err)

	want := []App{
		{
			CreatedAt:    "2023-08-03 20:37:41",
			Description:  "Foo",
			ID:           5,
			Name:         "Big App 2",
			RedirectURLs: "",
			UserID:       1,
		},
		{
			CreatedAt:    "2023-08-03 20:37:07",
			Description:  "",
			ID:           4,
			Name:         "Big App",
			RedirectURLs: "",
			UserID:       1,
		},
		{
			CreatedAt:    "2023-08-03 19:45:32",
			Description:  "",
			ID:           3,
			Name:         "xyz",
			RedirectURLs: "",
			UserID:       1,
		},
		{
			CreatedAt:    "2023-08-03 00:22:48",
			Description:  "",
			ID:           2,
			Name:         "app2",
			RedirectURLs: "",
			UserID:       1,
		},
		{
			CreatedAt:    "2023-07-13 09:46:52",
			Description:  "",
			ID:           1,
			Name:         "bgd",
			RedirectURLs: "",
			UserID:       1,
		},
	}

	assertApps(t, want, resp.Data)
}

func TestApps_GetAR(t *testing.T) {
	srv := NewServer(t)
	defer srv.Close()

	token := "TOKEN"

	srv.mux.HandleFunc("/portal-api/apps/1/access-requests/30", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpParse(t, "portal-api/get_ar_success.txt")
		defer httpResponse.Body.Close()

		assertMethod(t, "GET", r)
		assertHeader(t, r, "Authorization", token)

		w.WriteHeader(httpResponse.StatusCode)
		_, err := io.Copy(w, httpResponse.Body)
		assert.NoError(t, err)
	})

	client, err := New(
		WithBaseURL(srv.srv.URL),
		WithToken(token),
	)
	assert.NoError(t, err)

	resp, err := client.Apps().GetAR(context.Background(), 1, 30)
	assert.NoError(t, err)

	want := &ARDetails{
		Client:    "v34",
		Catalogue: "Public Catalogue",
		Plan:      "free_plan",
		Products: []string{
			"puvlic_product",
		},
	}

	assertAR(t, want, resp.Data)
}

func assertApps(t *testing.T, want, got []App) {
	require.Equal(t, len(want), len(got), "wanted len %v but got len", len(want), len(got))

	for k, v := range want {
		assert.Equal(t, v, got[k])
	}
}

func assertAR(t *testing.T, want, got *ARDetails) {
	assert.Equal(t, want.AuthType, got.AuthType, "wanted auth type %v but got %v", want.AuthType, got.AuthType)
	assert.Equal(t, want.Catalogue, got.Catalogue, "wanted catalogue %v but got %v", want.Catalogue, got.Catalogue)
	assert.Equal(t, want.Client, got.Client, "wanted client %v but got %v", want.Client, got.Client)
	assert.Equal(t, want.Plan, got.Plan, "wanted plan %v but got %v", want.Plan, got.Plan)
}
