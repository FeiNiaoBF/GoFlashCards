package server

import (
	view "github.com/FeiNiaoBF/GoFlashCards/public/template"
	"github.com/labstack/echo/v4"
)

func (server *Server) lookCard(c echo.Context) error {
	cards, err := server.getAllCardsHelper(c)
	if err != nil {
		return server.errorRequest(c, err)
	}
	outTags, err := server.getAllTagsHelper(c)
	if err != nil {
		return server.errorRequest(c, err)
	}
	return view.RenderHelper(c, view.LookHandler(cards, outTags))
}

func (server *Server) lookKnowCard(c echo.Context) error {
	cards, err := server.getAllCardsHelper(c)
	if err != nil {
		return server.errorRequest(c, err)
	}
	outTags, err := server.getAllTagsHelper(c)
	if err != nil {
		return server.errorRequest(c, err)
	}
	return view.RenderHelper(c, view.LookHandler(cards, outTags))
}
