package store

import (
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"uuid"
)

type Store struct {
	db *sql.DB
}

type Project struct {
	id          uuid.UUID
	name        string
	directory   string
	git_online  string
	created_at  time.Time
	modified_at time.Time
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
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}

func GetDefaultPath() string {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".fasttrack")
	os.MkdirAll(dir, os.ModePerm)

	return filepath.Join(dir, "fasttrack.db")
}

func (s *Store) CreateProject(name string, dir string) (*Project, error) {
	if s == nil || s.db == nil {
		return nil, errors.New("database not open")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	project := &Project{
		uuid.NewV4(), name, dir, "unset", time.Now(), time.Now(),
	}

	_, err = tx.Exec(`INSERT INTO projects (id, name, directory, "git-online", created, modified)
		VALUES (?, ?, ?, ?, ?, ?)`, project.id.String(), project.name, project.directory, project.git_online, project.created_at, project.modified_at)
	if err != nil {
		return nil, err
	}

	return project, tx.Commit()
}

func (s *Store) GetProjectByDir(dir string) (*Project, error) {
	if s == nil || s.db == nil {
		return nil, errors.New("database not open")
	}

	project := &Project{}
	var id string

	err := s.db.QueryRow(`SELECT id, name, directory, "git-online", created, modified
		FROM projects WHERE directory = ?`, dir).Scan(
		&id,
		&project.name,
		&project.directory,
		&project.git_online,
		&project.created_at,
		&project.modified_at,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	project.id, err = uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid project id %q: %w", id, err)
	}

	return project, nil
}
