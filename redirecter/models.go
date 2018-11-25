package redirecter

import "fmt"

type ShortenedURL struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
	URL string `json:"url"`
}

type Error struct {
	Err          string `json:"error"`
	ErrorMessage string `json:"error_message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Err, e.ErrorMessage)
}

func NewError(error, message string) error {
	return &Error{Err: error, ErrorMessage: message}
}
