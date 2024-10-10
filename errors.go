package sportmonks

import "fmt"

// ErrBadStatusCode is returned when the API returns a non 200 error code.
type ErrBadStatusCode struct {
	Message string `json:"message"`
}

func (e *ErrBadStatusCode) Error() string {
	return fmt.Sprintf("Request failed with the message: %s", e.Message)
}

// ErrRateLimit is returned when the API returns a 429 error code.
type ErrRateLimit struct {
	Message   string `json:"message"`
	Link      string `json:"link"`
	ResetCode string `json:"reset_code"`
}

func (e *ErrRateLimit) Error() string {
	return fmt.Sprintf("Request failed with the message: '%s', link: '%s', reset code: '%s'", e.Message, e.Link, e.ResetCode)
}
