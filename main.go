package main

import (
	"github.com/FeiNiaoBF/GoFlashCards/internal/handler"
	"github.com/FeiNiaoBF/GoFlashCards/internal/store"
	"github.com/labstack/echo/v4"
)


func main() {
	store.Run()

	e := echo.New()
	e.GET("/", handler.IndexHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
