package store

import (
	"context"
	"log"

	"github.com/FeiNiaoBF/GoFlashCards/db/sqlc"
	"github.com/jackc/pgx/v5"
)

const (
	dbSource = "postgresql://root:secret@localhost:5432/anki?sslmode=disable"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	sqlc.Querier
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	conn *pgx.Conn
	*sqlc.Queries
}

// NewStore creates a new store
func NewStore(conn *pgx.Conn) Store {
	return &SQLStore{
		conn:    conn,
		Queries: sqlc.New(conn),
	}
}

func Run() {
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

	NewStore(conn)
}
