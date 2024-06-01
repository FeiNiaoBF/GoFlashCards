package server

import (
	"fmt"
	"net/http"

	"github.com/FeiNiaoBF/GoFlashCards/db/sqlc"
	"github.com/FeiNiaoBF/GoFlashCards/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
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
	server.setMiddle()
	// server.setTemplate()
	server.setStaticFile()

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

	router.Use(middleware.Recover())

	router.GET("/", server.home)
	// card group
	router.Group("/card")
	router.GET("/card", server.getAllCards)
	router.POST("/card/add", server.createCards)
	router.PUT("/card/know/:id", server.upKnowCard)
	// router.DELETE("/card/delete/:id", server.deleteCard)

	// tag group
	router.Group("/tag")
	router.GET("/tag", server.getAllTags)
	router.POST("/tag/add", server.createTags)
	// router.GET("/tag/edit", )
	// router.PUT("/tag/update", server.updateTag)
	// router.DELETE("/tag/delete/:id", server.deleteTag)

	router.Group("/show")
	router.GET("/show", server.ShowAll)
	// router.GET("/show/all", server.ShowAll)
	router.GET("/show/:name", server.ShowKnow) // 0 is no know or 1 is know
	// router.GET("/show/unknown", server.ShowUnKnown)

	router.Group("/look")
	router.GET("/look", server.lookCard)
	router.GET("/look/tag/:tid/:cardid", server.lookCardByTag)
	// router.POST("/look", server.lookCard)
	router.GET("/look/know", server.lookKnowCard)

	server.router = router
}

// func (server *Server) setTemplate() {
// 	// server.router.Renderer = newTemplate()
// }

func (server *Server) setStaticFile() {
	server.router.Static("/public", "public")
}

func (server *Server) setMiddle() {
	// middleware
	logger, _ := zap.NewProduction()
	server.router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))
}
