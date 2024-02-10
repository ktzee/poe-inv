package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ktzee/poe-inv/model"
)

type PriceData struct {
	ScarabPriceData  *model.ItemPriceData
	DeliorbPriceData *model.ItemPriceData
}

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
}

func main() {
	pricedata := new(PriceData)
	pricedata.ScarabPriceData = model.NewPriceData(
		"https://poe.ninja/api/data/itemoverview?league=Affliction&type=",
		"Scarab")
	pricedata.DeliorbPriceData = model.NewPriceData(
		"https://poe.ninja/api/data/itemoverview?league=Affliction&type=",
		"DeliriumOrb")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()

	e.Static("/static", "./static")
	e.Static("/node_modules/", "./node_modules")
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", pricedata)
	})
	e.POST("/add-scarab", func(c echo.Context) error {
		selectedScarabName := c.FormValue("name")
		selectedScarabValueC := c.FormValue("chaosvalue")
		fmt.Println(selectedScarabName)
		fmt.Println(selectedScarabValueC)
		return c.Render(200, "selectedScarabs", selectedScarabValueC)
	})
	e.Logger.Fatal(e.Start(":3000"))
}
