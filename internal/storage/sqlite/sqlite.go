package sqlite

import (
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

// func New(storagePath string) (*Storage, error) {
// 	const op = "storage.sqlite.New"

// 	if db, err := sql.Open("sqlite3", storagePath); err != nil {
// 		return nil, fmt.Errorf("%s: %w", op, err)
// 	}
// }
