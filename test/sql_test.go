package test

import (
	"context"
	"github.com/FeiNiaoBF/GoFlashCards/db/sqlc"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5"
	"testing"
)

const (
	dbSource = "postgresql://root:secret@localhost:5432/anki?sslmode=disable"
)

var testQueries *sqlc.Queries

func TestSQL(t *testing.T) {
	//var testQueries *sqlc.Queries
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

	testQueries = sqlc.New(conn)

	// create a card
	insertedCards, err := testQueries.CreateCards(ctx, sqlc.CreateCardsParams{
		Front: "test",
		Back:  "test",
	})
	if err != nil {
		log.Fatal("cannot create a card: ", err)
	}
	log.Println(insertedCards)

	// get a card by id

	fetchedCards, err := testQueries.GetCard(ctx, insertedCards.ID)
	if err != nil {
		log.Fatal("cannot get a card: ", err)
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedCards, fetchedCards))

}
