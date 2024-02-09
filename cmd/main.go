package main

import (
	"github.com/labstack/echo/v4"

	"github.com/ktzee/poe-inv/handler"
	"github.com/ktzee/poe-inv/model"
)

func main() {
	data := model.NewPriceData(
		"https://poe.ninja/api/data/itemoverview?league=Affliction&type=",
		"Scarab",
	)
	app := echo.New()
	app.Static("/static", "./static")
	app.Static("/node_modules/", "./node_modules")
	scarabHandler := handler.ScarabHandler{ScarabData: data}
	app.GET("/", scarabHandler.ScarabHandlerShow)
	app.POST("/search-scarab", scarabHandler.ScarabHandlerSearch)
	app.Logger.Fatal(app.Start(":3000"))
}
