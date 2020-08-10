package db

import (
	"book-survey/internal/services"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InsertItems(database *sql.DB, decResBoby *services.DecResBodyType) error {
	return nil
}
