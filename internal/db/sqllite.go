package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB() *sql.DB {
	database, err := sql.Open("sqlite", "./expenses.db")
	if err != nil {
		log.Fatal(err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS expenses (
		id TEXT PRIMARY KEY,
		amount REAL NOT NULL,
		category TEXT NOT NULL,
		description TEXT,
		expense_date TEXT NOT NULL,
		created_at TEXT NOT NULL
	);
	`

	_, err = database.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	return database
}
