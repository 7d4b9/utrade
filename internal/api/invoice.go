package api

import (
	"errors"
	"net/http"

	uhttp "github.com/7d4b9/utrade/http"
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
			return c.JSON(http.StatusBadRequest, uhttp.CJSONError(c, err, "input", "bind create invoice"))
		}
		output, err := data.CreateInvoice(ctx, &input)
		if err != nil {
			if errors.Is(err, ErrAlreadyExists) {
				return c.JSON(http.StatusConflict, uhttp.JSONError{
					Message: "Conflict",
				})
			}
			return uhttp.CJSONError(c, err, "set_invoice_data", "repository create invoice")
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
			return uhttp.CJSONError(c, err, "get_invoice_data", "repository get invoice")
		}
		return c.JSON(http.StatusOK, output)
	}
}
