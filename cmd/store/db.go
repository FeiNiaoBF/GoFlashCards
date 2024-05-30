package store

import (
	"context"
	"fmt"
	"log"

	"github.com/FeiNiaoBF/GoFlashCards/db/sqlc"
	"github.com/FeiNiaoBF/GoFlashCards/util"
	"github.com/jackc/pgx/v5"
)

var (
	db   *sqlc.Queries
	conn *pgx.Conn
	ctx  = context.Background()
	err  error
)

func InitDB(config util.Config) {
	// postgresql://root:secret@localhost:5432/anki?sslmode=disable
	dbSource := fmt.Sprintf("%s://%s:%s@localhost:%s/%s?sslmode=disable",
		config.DBDriver, config.DBRoot, config.DBPassword, config.DBPort, config.DBName)
	conn, err = pgx.Connect(ctx, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}
	db = sqlc.New(conn)
	log.Println("Successfully connected to database")
}

func GetDB() *sqlc.Queries {
	return db
}

func CloseDB() error {
	err = conn.Close(ctx)
	if err != nil {
		log.Fatal("cannot close database: ", err)
	}

	log.Println("Successfully closed database connection")
	return nil
}
