package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ktzee/poe-inv/view"
)

func IndexHandler(c echo.Context) error {
	return Render(c, http.StatusOK, view.Index())
}
