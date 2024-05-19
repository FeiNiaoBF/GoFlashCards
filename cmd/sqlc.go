package main

import (
	"context"
	"github.com/FeiNiaoBF/GoFlashCards/db/sqlc"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=root dbname=anki sslmode=verify-full")
	if err != nil {
		return err
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
		}
	}(conn, ctx)

	queries := sqlc.New(conn)

	// list all authors
	authors, err := queries.ListCards(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedCards, err := queries.CreateCards(ctx, sqlc.CreateCardsParams{
		Front: "hello",
		Back:  "world",
	})
	if err != nil {
		return err
	}
	log.Println(insertedCards)

	// get the author we just inserted
	fetchedCards, err := queries.GetCard(ctx, insertedCards.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedCards, fetchedCards))
	return nil
}
