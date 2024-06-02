package server

import (
	"log"
	"strconv"

	view "github.com/FeiNiaoBF/GoFlashCards/public/template"
	"github.com/labstack/echo/v4"
)

func (server *Server) editCard(c echo.Context) error {
	// ctx := c.Request().Context()
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		return server.errorRequest(c, err)
	}

	log.Println(id)
	// UPDATE CARD
	// server.updateCard(c)
	card, err := server.getCardByIdHelper(c, id)
	if err != nil {
		return server.errorRequest(c, err)
	}

	return view.RenderHelper(c, view.EditHandler(card))
}
