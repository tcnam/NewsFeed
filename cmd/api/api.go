package api

import (
	"database/sql"
)

type APIServer struct{
	address string
	db *sql.DB
}

func new