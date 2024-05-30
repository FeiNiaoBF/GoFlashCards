package main

import (
	"log"

	"github.com/FeiNiaoBF/GoFlashCards/cmd/server"
	"github.com/FeiNiaoBF/GoFlashCards/cmd/store"
	"github.com/FeiNiaoBF/GoFlashCards/util"
)

func main() {
	// config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	// RUN SQL
	store.InitDB(config)
	defer store.CloseDB()
	// RUN echo
	server := server.NewServer(config, store.GetDB())
	server.Run()
}
