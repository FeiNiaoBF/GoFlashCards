package main

import (
	"context"
	"github.com/FeiNiaoBF/GoFlashCards/cmd"
	"github.com/FeiNiaoBF/GoFlashCards/cmd/echo"
	"github.com/jackc/pgx/v5"
	"log"
)

const (
	dbSource = "postgresql://root:secret@localhost:5432/anki?sslmode=disable"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dbSource)

	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Fatal("cannot close database: ", err)
		}
	}(conn, ctx)

	server := cmd.NewServer(conn)
	echo.Set(server, ctx)
	echo.Run()
}
