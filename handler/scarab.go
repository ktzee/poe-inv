package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ktzee/poe-inv/model"
	"github.com/ktzee/poe-inv/view"
)

type ScarabHandler struct {
	ScarabData *model.PriceData
}

func (h ScarabHandler) ScarabHandlerShow(c echo.Context) error {
	return Render(
		c,
		http.StatusOK,
		view.Scarab(*h.ScarabData),
	)
}

func (h ScarabHandler) ScarabHandlerSearch(c echo.Context) error {
	*&h.ScarabData.SearchItems(c.FormValue("search"))
	return Render(
		c,
		http.StatusOK,
		view.ScarabSearchResults(*h.ScarabData),
	)
}
