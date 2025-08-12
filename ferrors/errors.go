package ferrors

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	Status           int    `json:"-"`
	Code             string `json:"code"`
	ErrorDescription string `json:"-"`
}

func (e HttpError) Error() string {
	return fmt.Sprintf("code: %s, error_desc: %s", e.Code, e.ErrorDescription)
}

func NewHttpError(status int, code, errorDesc string) HttpError {
	return HttpError{status, code, errorDesc}
}

func BadRequestError(code, errorFormat string, vals ...any) HttpError {
	return NewHttpError(
		http.StatusBadRequest,
		code,
		fmt.Sprintf(errorFormat, vals...),
	)
}

// Internal server error
func InternalServerError(errorFmt string, vals ...any) error {
	return NewHttpError(
		http.StatusInternalServerError,
		"internal_server_error",
		fmt.Sprintf(errorFmt, vals...),
	)
}
