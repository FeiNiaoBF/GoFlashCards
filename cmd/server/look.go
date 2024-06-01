package server

import (
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
	tagId := c.Param("tid")
	cardid := c.Param("cardid")
	// log.Print(tagId)
	// log.Print(cardid)
	var err error
	// set tag id
	tid := 0
	if tagId == "" {
		tid = 0
	} else {
		tid, err = strconv.Atoi(tagId)
		if err != nil {
			return server.errorRequest(c, err)
		}
	}
	// set card id
	cid := 0
	if cardid == "" {
		cid = 0
	} else {
		cid, _ = strconv.Atoi(cardid)
	}

	cards, err := server.getCardByTagHelper(c, tid)
	if err != nil {
		return server.errorRequest(c, err)
	}

	outTags, err := server.getAllTagsHelper(c)
	if err != nil {
		return server.errorRequest(c, err)
	}

	return view.RenderHelper(c, view.LookAllCardHandler(cards, outTags, tid, cid))
}
