package echo

import (
	"context"
	"github.com/FeiNiaoBF/GoFlashCards/cmd"
	"github.com/labstack/echo/v4"
)

type Server struct {
	e   *echo.Echo
	sql *cmd.Server
	ctx context.Context
}

var server *Server

func init() {
	// 初始化 server
	server = newServer()
	// 注册路由
	RegisterRouter()
}

func newServer() *Server {
	return &Server{
		e: echo.New(),
	}
}

func Run() {
	server.e.Logger.Fatal(server.e.Start(":1314"))
}

func Set(sqlServer *cmd.Server, c context.Context) {
	server.sql = sqlServer
	server.ctx = c
}
