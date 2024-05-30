package server

import (
	"fmt"
	"net/http"

	"github.com/FeiNiaoBF/GoFlashCards/cmd/handler"
	"github.com/FeiNiaoBF/GoFlashCards/db/sqlc"
	"github.com/FeiNiaoBF/GoFlashCards/util"
	"github.com/labstack/echo/v4"
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

type user struct {
	ID string `query:"id" json:"id" form:"id"`
}

func (server *Server) setRouter() {
	router := echo.New()

	router.GET("/", handler.Home)
	router.POST("/card", server.createCards)
	// router.POST("/users/login", server.loginUser)
	// router.POST("/tokens/renew_access", server.renewAccessToken)
	router.POST("/users", func(c echo.Context) error {
		// in the handler for /users?id=<userID>
		var user user
		err := c.Bind(&user)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		return c.JSON(http.StatusOK, user)
	})
	server.router = router
}

func (server *Server) Run() {
	serverAddr := fmt.Sprintf("%s:%s", server.config.ServerAddress, server.config.ServerPort)
	server.router.Start(serverAddr)
}

func (server *Server) errorRequest(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, echo.Map{
		"error": err.Error(),
	})
}
