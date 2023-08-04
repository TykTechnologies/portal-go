// Copyright 2023 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApps_Get(t *testing.T) {
	srv := NewServer(t)
	defer srv.Close()

	token := "TOKEN"

	srv.mux.HandleFunc("/portal-api/apps", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpParse(t, "portal-api/apps.txt")

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
		{},{},{},{},{},
	}

	assertApps(t, want, resp.Data)
}

func assertApps(t *testing.T, want, got []App) {
	assert.Equal(t, len(want), len(got), "wanted len %v but got len", len(want), len(got))
}
