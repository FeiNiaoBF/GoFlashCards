package view

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderHelper(c echo.Context, compnent templ.Component) error {
	return compnent.Render(c.Request().Context(), c.Response())
}
