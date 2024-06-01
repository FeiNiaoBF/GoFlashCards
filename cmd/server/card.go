package server

import (
	"net/http"
	"strconv"

	"github.com/FeiNiaoBF/GoFlashCards/cmd/model"
	"github.com/FeiNiaoBF/GoFlashCards/db/sqlc"
	view "github.com/FeiNiaoBF/GoFlashCards/public/template"
	"github.com/labstack/echo/v4"
)

func (server *Server) createCards(c echo.Context) error {
	var card model.CardInput
	ctx := c.Request().Context()
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

	_, err := server.store.CreateCards(ctx, arg) // Assign the result to a new variable
	// log.Println(newCard)
	if err != nil {
		// Handle the error
		return server.errorRequest(c, err)
	}
	// TODO: Use log
	c.Logger()
	outCard, err := server.getAllCardsHelper(c)
	if err != nil {
		return server.errorRequest(c, err)
	}
	outTags, err := server.getAllTagsHelper(c)
	if err != nil {
		return server.errorRequest(c, err)
	}
	return view.RenderHelper(c, view.AddCardHandler(outCard, outTags))
}

// /card:id method: GET
func (server *Server) getAllCards(c echo.Context) error {
	outCard, err := server.getAllCardsHelper(c)
	if err != nil {
		return server.errorRequest(c, err)
	}
	outTags, err := server.getAllTagsHelper(c)
	if err != nil {
		return server.errorRequest(c, err)
	}

	return view.RenderHelper(c, view.CardHandler(outCard, outTags))
}

// getAllCardsHelper is a helper function
func (server *Server) getAllCardsHelper(c echo.Context) ([]model.CardOutput, error) {
	ctx := c.Request().Context()

	cards, err := server.store.GetAllCard(ctx)
	if err != nil {
		return nil, err
	}

	outCard := make([]model.CardOutput, len(cards))
	for i, card := range cards {
		outCard[i] = model.CardOutput{
			Front:  card.Front,
			Back:   card.Back,
			TagsID: card.TagsID,
			Know:   card.Know,
		}
	}

	return outCard, nil
}

// getListCard is helper function
func (server *Server) getListCardHelper(c echo.Context, page sqlc.ListCardsParams) ([]model.CardOutput, error) {
	ctx := c.Request().Context()

	cards, err := server.store.ListCards(ctx, page)
	if err != nil {
		return nil, err
	}

	outCard := make([]model.CardOutput, len(cards))
	for i, card := range cards {
		outCard[i] = model.CardOutput{
			Front:  card.Front,
			Back:   card.Back,
			TagsID: card.TagsID,
			Know:   card.Know,
		}
	}

	return outCard, nil
}

// getAllCardsHelper is a helper function to get all cards is know/unknow
func (server *Server) getKnowCardsHelper(c echo.Context, know bool) ([]model.CardOutput, error) {
	ctx := c.Request().Context()

	cards, err := server.store.GetCardByKnow(ctx, know)
	if err != nil {
		return nil, err
	}

	outCard := make([]model.CardOutput, len(cards))
	for i, card := range cards {
		outCard[i] = model.CardOutput{
			Front:  card.Front,
			Back:   card.Back,
			TagsID: card.TagsID,
			Know:   card.Know,
		}
	}

	return outCard, nil
}

// getCardByTag is a helper function to get card by tag
func (server *Server) getCardByTagHelper(c echo.Context, tagid int) ([]model.CardOutput, error) {
	ctx := c.Request().Context()

	cards, err := server.store.GetCardByTag(ctx, int32(tagid))
	if err != nil {
		return nil, err
	}

	outCard := make([]model.CardOutput, len(cards))
	for i, card := range cards {
		outCard[i] = model.CardOutput{
			Front:  card.Front,
			Back:   card.Back,
			TagsID: card.TagsID,
			Know:   card.Know,
		}
	}

	return outCard, nil
}

func (server *Server) updateCard(c echo.Context) error {
	var card model.CardInput
	ctx := c.Request().Context()
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
		Know:   newCard.Know,
	}

	// Use the newCard variable as needed

	return c.JSON(http.StatusOK, outCard)
}

func (server *Server) deleteCard(c echo.Context) error {
	ctx := c.Request().Context()
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
