package server

import (
	"net/http"
	"strconv"

	"github.com/FeiNiaoBF/GoFlashCards/cmd/model"
	"github.com/labstack/echo/v4"
)

// func (server *Server) createTags(c echo.Context) error {
// 	var Tag model.TagInput
// 	ctx := c.Request().Context()
// 	if err := c.Bind(&Tag); err != nil {
// 		return server.errorRequest(c, err)
// 	}

// 	// log.Println(Tag)
// 	name := util.RandText(5)
// 	// log.Println(arg)

// 	newTag, err := server.store.CreateTags(ctx, name) // Assign the result to a new variable
// 	// log.Println(newTag)
// 	if err != nil {
// 		// Handle the error
// 		return server.errorRequest(c, err)
// 	}

// 	outTag := model.TagOutput{
// 		ID:   int(newTag.ID),
// 		Name: newTag.Name,
// 	}

// 	// Use the newTag variable as needed

// 	return c.JSON(http.StatusOK, outTag)
// }

// func (server *Server) updateTag(c echo.Context) error {
// 	var Tag model.TagInput
// 	ctx := c.Request().Context()
// 	if err := c.Bind(&Tag); err != nil {
// 		return server.errorRequest(c, err)
// 	}

// 	// log.Println(Tag)
// 	arg := sqlc.UpdateTagsParams{
// 		ID:   Tag.ID,
// 		Name: Tag.Name,
// 	}
// 	// log.Println(arg)

// 	newTag, err := server.store.UpdateTags(ctx, arg) // Assign the result to a new variable
// 	// log.Println(newTag)
// 	if err != nil {
// 		return server.errorRequest(c, err)
// 	}

// 	outTag := model.TagOutput{
// 		ID:   int(newTag.ID),
// 		Name: newTag.Name,
// 	}

// 	// Use the newTag variable as needed

// 	return c.JSON(http.StatusOK, outTag)
// }

func (server *Server) deleteTag(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	TagID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return server.errorRequest(c, err)
	}

	err = server.store.DeleteTags(ctx, TagID)
	if err != nil {
		return server.errorRequest(c, err)
	}

	return c.JSON(http.StatusOK, "Deleted successfully")
}

// helper function to get all tags
func (server *Server) getAllTags(c echo.Context) ([]model.TagOutput, error) {
	ctx := c.Request().Context()

	tags, err := server.store.GetAllTags(ctx)
	if err != nil {
		return []model.TagOutput{}, server.errorRequest(c, err)
	}

	outTag := make([]model.TagOutput, len(tags))
	for i, tag := range tags {
		outTag[i] = model.TagOutput{
			ID:   strconv.Itoa(int(tag.ID)),
			Name: tag.Name,
		}
	}

	return outTag, nil
}
