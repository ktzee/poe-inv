package main

import (
	"github.com/labstack/echo/v4"

	"github.com/ktzee/poe-inv/handler"
)

func main() {
	app := echo.New()
	app.Static("/static", "../static")
	app.GET("/", handler.IndexHandler)
	app.Logger.Fatal(app.Start(":3000"))
	// data := fetchPriceData(
	// 	"https://poe.ninja/api/data/itemoverview?league=Affliction&type=",
	// 	"Scarab",
	// )
}
