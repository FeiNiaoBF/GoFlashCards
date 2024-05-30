package server

import (
	"context"
	"net/http"

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
