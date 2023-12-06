package portal

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
)

const (
	pathThemesUpload = "/portal-api/themes/upload"
	pathThemes       = "/portal-api/themes"
	pathTheme        = "/portal-api/themes/%v"
)

//go:generate mockery --name Themes --filename themes.go
type Themes interface {
	UploadTheme(ctx context.Context, input io.Reader, opts ...Option) (*UploadThemeOutput, error)
}

type themes struct {
	client *Client
}

func (t themes) UploadTheme(ctx context.Context, input io.Reader, opts ...Option) (*UploadThemeOutput, error) {
	form, contentType, err := createThemeForm(input)
	if err != nil {
		return nil, err
	}

	_, err = t.client.doPost(
		ctx,
		pathThemesUpload,
		form,
		nil,
		WithHeaders(
			map[string]string{
				"Content-Type": contentType,
			},
		),
	)

	if err != nil {
		return nil, err
	}

	return &UploadThemeOutput{}, nil
}

func createThemeForm(r io.Reader) (io.Reader, string, error) {
	buf := &bytes.Buffer{}

	formWriter := multipart.NewWriter(buf)
	defer formWriter.Close()

	fileWriter, err := formWriter.CreateFormFile("file", "theme.zip")
	if err != nil {
		return nil, "", err
	}

	if _, err = io.Copy(fileWriter, r); err != nil {
		return nil, "", err
	}

	return buf, formWriter.FormDataContentType(), nil
}

type UploadThemeOutput struct{}

type Theme struct {
	Author  string `json:"Author,omitempty"`
	ID      string `json:"ID,omitempty"`
	Name    string `json:"Name,omitempty"`
	Path    string `json:"Path,omitempty"`
	Status  string `json:"Status,omitempty"`
	Version string `json:"Version,omitempty"`
}

type ThemeOutput struct {
	Data *Theme
}

type ListThemesOutput struct {
	Data []Theme
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
