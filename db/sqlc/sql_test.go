package sqlc

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

const (
	dbSource = "postgresql://root:secret@localhost:5432/anki?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
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

	testQueries = New(conn)
	os.Exit(m.Run())
}
