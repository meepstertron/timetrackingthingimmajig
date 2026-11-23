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

type Session struct {
	id      uuid.UUID
	started time.Time
	ended   time.Time
	project uuid.UUID
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

func (p *Project) Name() string { return p.name }

func (p *Project) Directory() string { return p.directory }

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

func (s *Store) GetProjects() ([]*Project, error) {
	if s == nil || s.db == nil {
		return nil, errors.New("database not open")
	}

	rows, err := s.db.Query(`SELECT id, name, directory, "git-online", created, modified FROM projects`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*Project
	for rows.Next() {
		project := &Project{}
		var id string

		err := rows.Scan(
			&id,
			&project.name,
			&project.directory,
			&project.git_online,
			&project.created_at,
			&project.modified_at,
		)
		if err != nil {
			return nil, err
		}

		project.id, err = uuid.Parse(id)
		if err != nil {
			return nil, fmt.Errorf("invalid project id %q: %w", id, err)
		}

		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func (s *Store) GetProjectById(id uuid.UUID) (*Project, error) {
	if s == nil || s.db == nil {
		return nil, errors.New("database not open")
	}

	project := &Project{}

	err := s.db.QueryRow(`SELECT id, name, directory, "git-online", created, modified
		FROM projects WHERE id = ?`, id.String()).Scan(
		&project.id,
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

	return project, nil
}

func (s *Store) GetSessionForProject(project uuid.UUID) (*Session, error) {
	if s == nil || s.db == nil {
		return nil, errors.New("database not open")
	}

	var (
		id, projectID string
		started       time.Time
		ended         sql.NullTime
	)

	err := s.db.QueryRow(`
		SELECT id, started, ended, project
		FROM sessions
		WHERE project = ?
		ORDER BY started DESC
		LIMIT 1`, project.String()).Scan(&id, &started, &ended, &projectID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	sessionID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid session id %q: %w", id, err)
	}
	storedProjectID, err := uuid.Parse(projectID)
	if err != nil {
		return nil, fmt.Errorf("invalid session project id %q: %w", projectID, err)
	}

	session := &Session{
		id:      sessionID,
		started: started,
		project: storedProjectID,
	}
	if ended.Valid {
		session.ended = ended.Time
	}

	return session, nil
}

func (s *Store) StartSession(project_id uuid.UUID) (*Session, error) {
	project, err := s.GetProjectById(project_id)
	if err != nil {
		return nil, err
	}
	if project == nil {
		return nil, errors.New("project not found")
	}

	existingSession, err := s.GetSessionForProject(project.id)
	if err != nil {
		return nil, err
	}

	if existingSession != nil {
		return nil, errors.New("session already exists")
	}

	session := &Session{uuid.NewV4(), time.Now(), time.Time{}, project_id}

	_, err = s.db.Exec("INSERT INTO sessions (id, started, ended, project) VALUES (?, ?, ?, ?)", session.id, session.started, session.ended, session.project)
	if err != nil {
		return nil, err
	}
	return session, nil
}
