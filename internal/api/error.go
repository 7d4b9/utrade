package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type jsonError struct {
	Message string
	Code    string
}

// cJSONError is returned from the API as JSON.
func cJSONError(c echo.Context, err error, message, log string) error {
	code := uuid.New().String()
	c.Logger().Errorf("%v: %v, code: %v", log, err, code)
	return c.JSON(http.StatusInternalServerError, jsonError{
		Code:    code,
		Message: message,
	})
}
