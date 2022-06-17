package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Invoice struct {
	ID string
}

func createInvoice(repository Data) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var input Invoice
		if err := c.Bind(&input); err != nil {
			c.Logger().Errorf("bind input create invoice: %w", err)
			return c.JSON(http.StatusBadRequest, JSONError{Message: "input"})
		}
		output, err := repository.CreateInvoice(ctx, &input)
		if err != nil {
			if errors.Is(ErrAlreadyExists, err) {
				return c.JSON(http.StatusConflict, JSONError{Message: "conflict"})
			}
		}
		return c.JSON(http.StatusOK, output)
	}
}

func getInvoice(repository Data) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param("id")
		output, err := repository.GetInvoice(ctx, id)
		if err != nil {
			if errors.Is(ErrAlreadyExists, err) {
				return c.JSON(http.StatusConflict, JSONError{Message: "conflict"})
			}
		}
		return c.JSON(http.StatusOK, output)
	}
}
