package db

import (
	"book-survey/internal/io"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Initialize(dbInfo *io.DbInfoType) (*sql.DB, error) {
	dataSrcName := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", dbInfo.User, dbInfo.Pw, dbInfo.Database)
	database, err := sql.Open("mysql", dataSrcName)
	if err != nil {
		return nil, fmt.Errorf("While opening database: %w", err)
	}

	return database, nil
}
