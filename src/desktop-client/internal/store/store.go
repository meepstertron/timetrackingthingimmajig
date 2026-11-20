package store

import (
	"database/sql"
	_ "embed"
	"fmt"
)

type Store struct {
	db *sql.DB
}

//go:embed schema.sql
var schema string

func Open(path string) (*Store, error) {
	db, err := sql.Open("sqlite", path+
		"?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)&_pragma=foreign_keys(1)")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	if _, err := db.Exec(schema); err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}
	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}
