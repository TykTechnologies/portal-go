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

func TestUser_Get(t *testing.T) {
	srv := NewServer(t)
	defer srv.Close()

	token := "TOKEN"

	srv.mux.HandleFunc("/portal-api/users/1", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpParse(t, "portal-api/users/1.txt")

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

	resp, err := client.Users().GetUser(context.Background(), 1)
	assert.NoError(t, err)

	want := &User{
		Active:            true,
		APITokenCreatedAt: "2023-06-08",
		CreatedAt:         "2023-06-08 09:16",
		Email:             "raava@tyk.io",
		First:             "John",
		Last:              "Doe",
		ID:                1,
		JWTToken:          "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJQcm92aWRlciI6Im5vbmUiLCJVc2VySUQiOiIkMmEkMTAkdk9MLkpXY0UxUDZuNWJwZDdFenZuTzBkWDJvamovU3ZIM3pZeTZHdFJobkguc3RBdy9pQTYifQ.WfWCIPHZ7isT-GlrghsFmvPNeXwoM-A7tRsInc6Qc8U",
		Org:               "",
		OrgID:             0,
		Provider:          "password",
		ResetPassword:     false,
		Role:              "super-admin",
		Teams:             nil,
		UpdatedAt:         "2023-06-08 09:16",
	}

	assert.Equal(t, want, resp.Data)
}
