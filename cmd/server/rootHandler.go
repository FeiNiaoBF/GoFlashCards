package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (server *Server) home(c echo.Context) error {
	ctx := c.Request().Context()

	card, err := server.store.GetCard(ctx, int64(1))
	if err != nil {
		return server.errorRequest(c, err)
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"card": card,
	})
}
