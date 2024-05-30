package server

import (
	"fmt"
	"net/http"

	"github.com/FeiNiaoBF/GoFlashCards/cmd/handler"
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

// setRouter
func (server *Server) setRouter() {
	router := echo.New()
	// middleware
	router.Use(middleware.Logger())

	router.GET("/", handler.Home)
	// card group
	
	router.POST("/card", server.createCards)
	router.GET("/card/:id", server.getCards)
	router.PUT("/card/:id", server.updateCard)
	router.DELETE("/card/:id", server.deleteCard)
	// router.POST("/tokens/renew_access", server.renewAccessToken)
	// router.GET("/tags", server.getTags)
	server.router = router
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
