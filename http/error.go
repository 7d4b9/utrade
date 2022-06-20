package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type JSONError struct {
	Message string
	Code    string
}

// CJSONError is returned from the API as JSON.
func CJSONError(c echo.Context, err error, message, log string) error {
	code := uuid.New().String()
	c.Logger().Errorf("%v: %v, code: %v", log, err, code)
	return c.JSON(http.StatusInternalServerError, JSONError{
		Code:    code,
		Message: message,
	})
}
