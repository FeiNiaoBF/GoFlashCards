package sqlc

import (
	"github.com/jackc/pgx/v5"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	conn *pgx.Conn
	*Queries
}

// NewStore creates a new store
func NewStore(conn *pgx.Conn) Store {
	return &SQLStore{
		conn:    conn,
		Queries: New(conn),
	}
}
