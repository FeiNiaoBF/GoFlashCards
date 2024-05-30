package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/FeiNiaoBF/GoFlashCards/cmd/model"
	"github.com/FeiNiaoBF/GoFlashCards/db/sqlc"
	"github.com/labstack/echo/v4"
)

func (server *Server) createCards(c echo.Context) error {
	var card model.CardInput
	ctx := context.Background()
	if err := c.Bind(&card); err != nil {
		return server.errorRequest(c, err)
	}

	// log.Println(card)
	arg := sqlc.CreateCardsParams{
		Front:  card.Front,
		Back:   card.Back,
		TagsID: card.TagsID,
	}
	// log.Println(arg)

	newCard, err := server.store.CreateCards(ctx, arg) // Assign the result to a new variable
	// log.Println(newCard)
	if err != nil {
		// Handle the error
		return server.errorRequest(c, err)
	}

	outCard := model.CardOutput{
		Front:  newCard.Front,
		Back:   newCard.Back,
		TagsID: newCard.TagsID,
		Konw:   newCard.Know,
	}

	// Use the newCard variable as needed

	return c.JSON(http.StatusOK, outCard)
}

// /card:id method: GET
func (server *Server) getCards(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")

	cardID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return server.errorRequest(c, err)
	}
	card, err := server.store.GetCard(ctx, cardID)
	if err != nil {
		return server.errorRequest(c, err)
	}

	outCard := model.CardOutput{
		Front:  card.Front,
		Back:   card.Back,
		TagsID: card.TagsID,
		Konw:   card.Know,
	}

	return c.JSON(http.StatusOK, outCard)
}

func (server *Server) updateCard(c echo.Context) error {
	var card model.CardInput
	ctx := context.Background()
	if err := c.Bind(&card); err != nil {
		return server.errorRequest(c, err)
	}

	// log.Println(card)
	arg := sqlc.UpdateCardsParams{
		Front:  card.Front,
		Back:   card.Back,
		TagsID: card.TagsID,
	}
	// log.Println(arg)

	newCard, err := server.store.UpdateCards(ctx, arg) // Assign the result to a new variable
	// log.Println(newCard)
	if err != nil {
		return server.errorRequest(c, err)
	}

	outCard := model.CardOutput{
		Front:  newCard.Front,
		Back:   newCard.Back,
		TagsID: newCard.TagsID,
		Konw:   newCard.Know,
	}

	// Use the newCard variable as needed

	return c.JSON(http.StatusOK, outCard)
}

func (server *Server) deleteCard(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")

	cardID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return server.errorRequest(c, err)
	}

	err = server.store.DeleteCards(ctx, cardID)
	if err != nil {
		return server.errorRequest(c, err)
	}

	return c.JSON(http.StatusOK, "Deleted successfully")
}
