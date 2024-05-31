package server

import (
	"github.com/FeiNiaoBF/GoFlashCards/db/sqlc"
	view "github.com/FeiNiaoBF/GoFlashCards/public/template"
	"github.com/labstack/echo/v4"
)

func (server *Server) home(c echo.Context) error {
	args := sqlc.ListCardsParams{
		Limit:  10,
		Offset: 10,
	}

	cards, err := server.getListCardHelper(c, args)
	if err != nil {
		return server.errorRequest(c, err)
	}

	return view.RenderHelper(c, view.Hello(cards))
}
