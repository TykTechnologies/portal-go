package portal

import "fmt"

type ClientError struct{}

type ServerError struct{}

type APIError struct {
	*APIResponse
	Errors []string `json:"errors,omitempty"`
	Status string   `json:"status,omitempty"`
}

func (e APIError) Error() string {
	err := "unknown error"
	if len(e.Errors) != 0 {
		err = e.Errors[0]
	}

	return fmt.Sprintf(
		"%v %v: %v %v",
		e.Response.Request.Method,
		e.Response.Request.URL,
		e.Response.StatusCode,
		err,
	)
}

type UnknownError struct {
	*APIResponse
}

func (u UnknownError) Error() string {
	return "unknown error"
}
