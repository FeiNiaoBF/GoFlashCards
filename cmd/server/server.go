package server

import (
	"fmt"
	"net/http"

	"github.com/FeiNiaoBF/GoFlashCards/db/sqlc"
	"github.com/FeiNiaoBF/GoFlashCards/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config util.Config
	store  *sqlc.Queries
	router *echo.Echo
}

func NewServer(config util.Config, store *sqlc.Queries) *Server {
	server := &Server{
		config: config,
		store:  store,
	}
	server.setRouter()

	return server
}

// Run starts the server
func (server *Server) Run() {
	serverAddr := fmt.Sprintf("%s:%s", server.config.ServerAddress, server.config.ServerPort)
	server.router.Start(serverAddr)
}

// errorRequest is a helper function to handle errors
func (server *Server) errorRequest(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, echo.Map{
		"error": err.Error(),
	})
}

// setRouter
func (server *Server) setRouter() {
	router := echo.New()
	// middleware
	router.Use(middleware.Logger())

	router.GET("/", server.home)
	// card group
	router.Group("/card")
	router.GET("/card/:id", server.getCards)
	router.POST("/card/add", server.createCards)
	router.PUT("/card/update/:id", server.updateCard)
	router.DELETE("/card/delete/:id", server.deleteCard)
	// tag group
	router.Group("/tag")
	router.GET("/tag/:id", server.getTags)
	router.POST("/tag/add", server.createTags)
	router.PUT("/tag/update/:id", server.updateTag)
	router.DELETE("/tag/delete/:id", server.deleteTag)

	server.router = router
	server.setTemplate()
	server.setStaticFile()
}

func (server *Server) setTemplate() {
	server.router.Renderer = newTemplate()
}

func (server *Server) setStaticFile() {
}
