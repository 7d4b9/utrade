package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Invoice struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func createInvoice(data Data) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var input Invoice
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, cJSONError(c, err, "input", "bind create invoice"))
		}
		output, err := data.CreateInvoice(ctx, &input)
		if err != nil {
			if errors.Is(err, ErrAlreadyExists) {
				return c.JSON(http.StatusConflict, jsonError{
					Message: "Conflict",
				})
			}
			return cJSONError(c, err, "set_invoice_data", "repository create invoice")
		}
		return c.JSON(http.StatusOK, output)
	}
}

func getInvoice(data Data) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param(invoiceParam)
		output, err := data.GetInvoice(ctx, id)
		if err != nil {
			if errors.Is(ErrNotFound, err) {
				return c.JSON(http.StatusNotFound, "Not Found")
			}
			return cJSONError(c, err, "get_invoice_data", "repository get invoice")
		}
		return c.JSON(http.StatusOK, output)
	}
}
