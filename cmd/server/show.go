package server

import (
	"log"

	"github.com/FeiNiaoBF/GoFlashCards/cmd/model"
	view "github.com/FeiNiaoBF/GoFlashCards/public/template"
	"github.com/labstack/echo/v4"
)

var filt *model.Filtype

func (server *Server) ShowAll(c echo.Context) error {
	cards, err := server.getAllCardsHelper(c)
	if err != nil {
		return server.errorRequest(c, err)
	}
	outTags, err := server.getAllTagsHelper(c)
	if err != nil {
		return server.errorRequest(c, err)
	}
	return view.RenderHelper(c, view.ShowHandler(cards, outTags, filt))
}

func (server *Server) ShowKnow(c echo.Context) error {
	name := c.Param("name")
	log.Print(name)

	if name == "unknow" {
		filt = model.NewFilt(name, model.Unknown)
		cards, err := server.getKnowCardsHelper(c, false)
		if err != nil {
			return server.errorRequest(c, err)
		}
		outTags, err := server.getAllTagsHelper(c)
		if err != nil {
			return server.errorRequest(c, err)
		}
		return view.RenderHelper(c, view.ShowHandler(cards, outTags, filt))
	} else if name == "know" {
		filt = model.NewFilt(name, model.Know)
		cards, err := server.getKnowCardsHelper(c, true)
		if err != nil {
			return server.errorRequest(c, err)
		}
		outTags, err := server.getAllTagsHelper(c)
		if err != nil {
			return server.errorRequest(c, err)
		}
		return view.RenderHelper(c, view.ShowHandler(cards, outTags, filt))
	} else {
		filt = model.NewFilt(name, model.All)
		return server.ShowAll(c)
	}
}
