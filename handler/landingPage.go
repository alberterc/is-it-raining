package handler

import (
	"github.com/alberterc/is-it-raining/view/input"
	"github.com/labstack/echo/v4"
)

type LandingPageHandler struct{}

func (h LandingPageHandler) HandleLandingPageShow(c echo.Context) error {
	return render(c, input.LandingPage())
}
