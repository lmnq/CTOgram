package store

import (
	"database/sql"

	// sqlite3 ..
	_ "github.com/mattn/go-sqlite3"
)

// Store ..
type Store struct {
	DB *sql.DB
}

// NewStore ..
func NewStore() (*Store, error) {
	db, err := sql.Open("sqlite3", "./cities.db")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err = createTables(db); err != nil {
		return nil, err
	}

	return &Store{DB: db}, nil
}

func createTables(db *sql.DB) error {
	_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS "cities" (
				"ID" INTEGER NOT NULL UNIQUE,
				"name" TEXT NOT NULL,
				"code" TEXT NOT NULL,
				"country_code" TEXT NOT NULL,
				PRIMARY KEY("ID" AUTOINCREMENT)
			);
	`)

	return err
}
