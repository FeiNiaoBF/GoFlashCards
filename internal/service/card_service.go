package service

import "database/sql"

type CardService interface {
	All() (*[]CradResp, error)
}

type sCard struct {
	DB *sql.DB
}
