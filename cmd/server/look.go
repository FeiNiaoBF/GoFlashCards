package server

import (
	"log"
	"strconv"

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

// TODO: Update the lookKnowCard method
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

func (server *Server) lookCardByTag(c echo.Context) error {
	id := c.Param("id")
	log.Print(id)
	tagId, err := strconv.Atoi(id)
	if err != nil {
		return server.errorRequest(c, err)
	}
	cards, err := server.getCardByTagHelper(c, tagId)
	if err != nil {
		return server.errorRequest(c, err)
	}

	outTags, err := server.getAllTagsHelper(c)
	if err != nil {
		return server.errorRequest(c, err)
	}

	return view.RenderHelper(c, view.LookAllCardHandler(cards, outTags))
}
