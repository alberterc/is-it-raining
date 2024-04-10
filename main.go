package main

import (
	"github.com/alberterc/is-it-raining/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Static("/static", "static")

	inputHandler := handler.LandingPageHandler{}
	app.GET("/", inputHandler.HandleLandingPageShow)

	// remove localhost for production
	app.Start("localhost:3000")
}
