package echo

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Get - "/"
func indexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Get - /cards
func cardsHandler(c echo.Context) error {
	cards, err := server.sql.Store.GetAllCard(server.ctx)
	if err != nil {
		server.e.Logger.Error("Not cards")
		return c.String(http.StatusInternalServerError, "Error fetching cards")
	}
	return c.Render(http.StatusOK, "cards.html", map[string]interface{}{
		"cards":       cards,
		"filter_name": "all",
	})
}
