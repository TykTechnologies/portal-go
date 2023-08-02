// Copyright 2023 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrgService_Get(t *testing.T) {
	srv := NewServer(t)
	defer srv.Close()

	token := "TOKEN"

	srv.mux.HandleFunc("/portal-api/organisations/1", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpParse(t, "portal-api/organisations/1.txt")

		assertMethod(t, "GET", r)
		assertHeader(t, r, "Authorization", token)
		assertHeader(t, r, "User-Agent", "Test")

		w.WriteHeader(httpResponse.StatusCode)
		_, err := io.Copy(w, httpResponse.Body)
		assert.NoError(t, err)
	})

	client, err := New(
		WithBaseURL(srv.srv.URL),
		WithToken(token),
		WithUserAgent("Test"),
	)
	assert.NoError(t, err)

	resp, err := client.Orgs().GetOrg(context.Background(), 1)
	assert.NoError(t, err)

	want := &Org{
		ID:   1,
		Name: "Default Org",
	}

	assertOrg(t, want, resp.Data)
}

func assertOrg(t *testing.T, want, got *Org) {
	assert.Equal(t, want.ID, got.ID, fmt.Sprintf("wanted id %v but got %v", want.ID, got.ID))
	assert.Equal(t, want.Name, got.Name, fmt.Sprintf("wanted id %v but got %v", want.Name, got.Name))
}
