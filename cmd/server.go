package cmd

import (
	"github.com/FeiNiaoBF/GoFlashCards/db/sqlc"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/jackc/pgx/v5"
)

type Server struct {
	// SQL
	Store sqlc.Store
}

func NewServer(conn *pgx.Conn) *Server {
	server := &Server{
		Store: sqlc.NewStore(conn),
	}
	return server
}
