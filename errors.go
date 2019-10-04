package sportmonks

import "fmt"

// ErrBadStatusCode is returned when the API returns a non 200 error code
type ErrBadStatusCode struct {
	Err *struct {
		Message      string `json:"message"`
		Code         int	`json:"code"`
	} `json:"error"`
}

func (e *ErrBadStatusCode) Error() string {
	return fmt.Sprintf("Request failed with message: %s, code: %d", e.Err.Message, e.Err.Code)
}
